package zigbee

type IEEEAddress [8]byte
type NetworkAddress uint16

type PANID [2]byte
type ExtendedPANID [8]byte
type NetworkKey [16]byte

// https://en.wikipedia.org/wiki/Zigbee#Radio_hardware
// https://acuitysupport.zendesk.com/hc/en-us/articles/225413967-Zigbee-Networking-Basics-35-000ft-view-
// https://community.smartthings.com/t/faq-networking-and-reducing-channel-interference-between-wifi-and-zigbee/40159
// https://www.metageek.com/training/resources/zigbee-wifi-coexistence.html
// https://www.researchgate.net/figure/WIFI-and-Zigbee-Overlapping-Channels-in-The-24Ghz-ISM-Band_fig1_265226405

// Channels 15, 20 and 26 are in the space between standard 20Mhz 802.11bg channels, so will be least likely to receive
// WiFi interference. Though 11 does clash slightly with Wifi Channel 1 lower sideband, thus default to channel 15.

var Channels = []uint8{11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26}
var ZLLChannels = []uint8{11, 15, 20, 25}
var DefaultZLLChannel = ZLLChannels[1]

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
