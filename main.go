package main

import (
	"fmt"
	"time"
)

type RouteRequest struct {
	Source      string
	Destination string
	Path        []string
	TTL         int
}

type RouteReply struct {
	Source      string
	Destination string
	Path        []string
}

var routeCache = make(map[string][]string)

func sendRREQ(src, dest string) {
	rreq := RouteRequest{Source: src, Destination: dest, Path: []string{src}, TTL: 5}
	fmt.Println("Sending RREQ:", rreq)
	propagateRREQ(rreq)
}

func propagateRREQ(rreq RouteRequest) {
	// Simulate propagation
	rreq.TTL--
	if rreq.TTL > 0 {
		rreq.Path = append(rreq.Path, "IntermediateNode")
		fmt.Println("Propagating RREQ:", rreq)
		if rreq.Destination == "NodeB" {
			sendRREP(rreq)
		} else {
			propagateRREQ(rreq)
		}
	} else {
		fmt.Println("RREQ TTL expired")
	}
}

func sendRREP(rreq RouteRequest) {
	rrep := RouteReply{Source: rreq.Destination, Destination: rreq.Source, Path: append(rreq.Path, rreq.Destination)}
	fmt.Println("Sending RREP:", rrep)
	updateRouteCache(rrep)
}

func updateRouteCache(rrep RouteReply) {
	routeCache[rrep.Destination] = rrep.Path
	fmt.Println("Updated Route Cache:", routeCache)
}

func main() {
	sendRREQ("NodeA", "NodeB")
	time.Sleep(time.Second)
}
