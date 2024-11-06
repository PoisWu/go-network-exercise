package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	server_address := net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8080}
	server, err := net.ListenUDP("udp", &server_address)

	if err != nil {
		log.Fatal(err)
	}

	defer server.Close()

	fmt.Println(server.LocalAddr().Network())

	for {
		var buf = make([]byte, 512)
		n, addr, err := server.ReadFrom(buf)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Printf("The message if from %v, %v\n", addr.Network(), addr.String())
		fmt.Printf("Receive %v bytes\n", n)
		fmt.Println(string(buf[0:]))
	}
}
