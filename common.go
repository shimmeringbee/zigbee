package zigbee

import (
	"fmt"
)

type IEEEAddress uint64

func (a IEEEAddress) String() string {
	return fmt.Sprintf("%016x", uint64(a))
}

var EmptyIEEEAddress = IEEEAddress(0)

type NetworkAddress uint16

type PANID uint16
type ExtendedPANID uint64
type NetworkKey [16]byte

type ZCLClusterID uint16

type EndpointDescription struct {
	Endpoint          byte
	ProfileID         uint16
	DeviceID          uint16
	DeviceVersion     uint8
	InClusterList     []ZCLClusterID
	OutClusterList    []ZCLClusterID
}

// https://en.wikipedia.org/wiki/Zigbee#Radio_hardware
// https://acuitysupport.zendesk.com/hc/en-us/articles/225413967-Zigbee-Networking-Basics-35-000ft-view-
// https://community.smartthings.com/t/faq-networking-and-reducing-channel-interference-between-wifi-and-zigbee/40159
// https://www.metageek.com/training/resources/zigbee-wifi-coexistence.html
// https://www.researchgate.net/figure/WIFI-and-Zigbee-Overlapping-Channels-in-The-24Ghz-ISM-Band_fig1_265226405

// Channels 15, 20 and 25 are in the space between standard 20Mhz 802.11bg channels, so will be least likely to receive
// WiFi interference. Though 11 does clash slightly with Wifi Channel 1 lower sideband, thus default to channel 15.

// Channel 26 is also outside of normal WiFi clashes, however it's not supported on all Zigbee devices and is not
// recognised by ZLL as a valid channel.

var Channels = []uint8{11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26}
var ZLLChannels = []uint8{11, 15, 20, 25}
var DefaultChannel = ZLLChannels[1]

var ChannelToFrequencies = map[uint8]uint16{
	11: 2405,
	12: 2410,
	13: 2415,
	14: 2420,
	15: 2425,
	16: 2430,
	17: 2435,
	18: 2440,
	19: 2445,
	20: 2450,
	21: 2455,
	22: 2460,
	23: 2465,
	24: 2470,
	25: 2475,
	26: 2480,
}

var TCLinkKey = NetworkKey{0x5a, 0x69, 0x67, 0x42, 0x65, 0x65, 0x41, 0x6c, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x30, 0x39}

var (
	BroadcastAll                 NetworkAddress = 0xffff
	BroadcastAlwaysOnReceivers   NetworkAddress = 0xfffd
	BroadcastRoutersCoordinators NetworkAddress = 0xfffc
	BroadcastLowPowerRouters     NetworkAddress = 0xfffb
)
