package zigbee

import (
	"math/rand"
	"time"
)

type NetworkConfiguration struct {
	PANID         PANID
	ExtendedPANID ExtendedPANID
	NetworkKey    NetworkKey
	Channel       uint8
}

func GenerateNetworkConfiguration() (nc NetworkConfiguration, err error) {
	rand.Seed(time.Now().UnixNano())

	nc.PANID = PANID(rand.Uint32() & 0x3fff)
	nc.ExtendedPANID = ExtendedPANID(rand.Uint64())

	networkKey := make([]byte, 16)
	_, err = rand.Read(networkKey)
	if err != nil {
		return
	}
	copy(nc.NetworkKey[:], networkKey)

	nc.Channel = DefaultChannel
	return
}
