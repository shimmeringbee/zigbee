package zigbee

import "crypto/rand"

var DEFAULT_ZLL_CHANNEL uint8 = 11
var ZLL_CHANNELS = []uint8{11, 15, 20, 25}

type NetworkConfiguration struct {
	PANID         PANID
	ExtendedPANID ExtendedPANID
	NetworkKey    NetworkKey
	Channel       uint8
}

func GenerateNetworkConfiguration() (nc NetworkConfiguration, err error) {
	panId := make([]byte, 2)
	_, err = rand.Read(panId)
	if err != nil {
		return
	}
	copy(nc.PANID[:], panId)

	extendedPanId := make([]byte, 8)
	_, err = rand.Read(extendedPanId)
	if err != nil {
		return
	}
	copy(nc.ExtendedPANID[:], extendedPanId)

	networkKey := make([]byte, 16)
	_, err = rand.Read(networkKey)
	if err != nil {
		return
	}
	copy(nc.NetworkKey[:], networkKey)

	nc.Channel = DEFAULT_ZLL_CHANNEL
	return
}
