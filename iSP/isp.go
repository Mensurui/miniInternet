package isp

import (
	"github.com/Mensurui/miniInternet/router"
)

type ISP struct {
	Name        string
	Routers     []*router.Router
	IsUpperTier bool
}

func ConnectISPs(isp1 *ISP, isp2 *ISP) {
	router1 := isp1.Routers[0]
	router2 := isp2.Routers[0]
	router1.Outgoing = router2.Incoming
	router2.Outgoing = router1.Incoming
}
