package packettransmission

import (
	"github.com/Mensurui/miniInternet"
)

type Packet struct {
	SourceIP      string
	DestinationIP string
	ProtocolType  string
	Payload       string
	TTL           int
}

type EndSystem struct {
	IP   string
	Data string
}

func (es *EndSystem) SendData(data, destIP string) []Packet {
	chunks := miniinternet.SsplitIntoChunks(data, 100)

	var packets []Packet
	for _, chunks := range chunks {
		packets = append(packets, Packet{
			SourceIP:      es.IP,
			DestinationIP: destIP,
			Payload:       chunks,
			TTL:           64,
		})
	}

	return packets
}
