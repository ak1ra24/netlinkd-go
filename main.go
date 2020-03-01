package main

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/vishvananda/netlink"
)

func main() {

	netlink_done := make(chan struct{})
	defer close(netlink_done)

	addr_ch := make(chan netlink.AddrUpdate, 10)
	link_ch := make(chan netlink.LinkUpdate, 10)
	route_ch := make(chan netlink.RouteUpdate, 10)

	if err := netlink.AddrSubscribe(addr_ch, netlink_done); err != nil {
		return
	}
	if err := netlink.LinkSubscribe(link_ch, netlink_done); err != nil {
		return
	}
	if err := netlink.RouteSubscribe(route_ch, netlink_done); err != nil {
		return
	}

	exitCh := make(chan struct{})

	go func() {
		defer func() { close(exitCh) }()

		for {
			select {
			case msg := <-addr_ch:
				fmt.Println("############################################")
				fmt.Println("[addr channel]")
				pp.Println(&msg)
			case msg := <-link_ch:
				fmt.Println("############################################")
				fmt.Println("[link channel]")
				pp.Println(&msg)
			case msg := <-route_ch:
				fmt.Println("############################################")
				fmt.Println("[route channel]")
				pp.Println(&msg)
			}
		}

	}()

	<-exitCh
}
