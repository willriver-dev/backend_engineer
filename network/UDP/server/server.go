package main

import (
	"fmt"
	"net"
)

func main() {
	p := make([]byte, 100)
	addr := net.UDPAddr{
		Port: 8081,
		IP:   net.ParseIP("127.0.0.1"),
	}
	server, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Some error %v\n", err)
		return
	}

	for {
		_, remoteAddr, err := server.ReadFromUDP(p)
		fmt.Printf("%v %s \n", remoteAddr, p)
		if err != nil {
			fmt.Printf("Some errors %v", err)
			continue
		}
	}
}
