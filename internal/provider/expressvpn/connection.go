package expressvpn

import (
	"github.com/qdm12/gluetun/internal/configuration/settings"
	"github.com/qdm12/gluetun/internal/models"
	"github.com/qdm12/gluetun/internal/provider/utils"
)

func (p *Provider) GetConnection(selection settings.ServerSelection) (
	connection models.Connection, err error) {
	port := getPort(selection)
	protocol := utils.GetProtocol(selection)

	servers, err := utils.FilterServers(p.servers, selection)
	if err != nil {
		return connection, err
	}

	var connections []models.Connection
	for _, server := range servers {
		for _, IP := range server.IPs {
			connection := models.Connection{
				Type:     selection.VPN,
				IP:       IP,
				Port:     port,
				Protocol: protocol,
				Hostname: server.Hostname,
			}
			connections = append(connections, connection)
		}
	}

	return utils.PickConnection(connections, selection, p.randSource)
}

func getPort(selection settings.ServerSelection) (port uint16) {
	const (
		defaultOpenVPNTCP = 0
		defaultOpenVPNUDP = 1195
		defaultWireguard  = 0
	)
	return utils.GetPort(selection, defaultOpenVPNTCP,
		defaultOpenVPNUDP, defaultWireguard)
}
