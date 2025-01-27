package router

import (
	"fmt"

	"github.com/Mensurui/miniInternet/packetTransmission"
)

type Router struct {
	RoutingTable map[string]string //DestIP -> NextHop
	Incoming     chan packettransmission.Packet
	Outgoing     chan packettransmission.Packet
}

func (r *Router) Start() {
	go func() {
		for p := range r.Incoming {
			p.TTL--
			if p.TTL <= 0 {
				fmt.Println("Packet dropped(TTL expired)")
				continue
			}
			nextHop := r.RoutingTable[p.DestinationIP]
			fmt.Printf("Routing packet to %s via %s\n", p.DestinationIP, nextHop)
			r.Outgoing <- p
		}
	}()
}
