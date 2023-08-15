package main

import (
	helper "awesomeProject/HelperLibrary"
	"bytes"
	"fmt"
	"net"
	"strconv"
)

func main() {
	network, address := "udp", "localhost:2001"

	udpAddr, err := net.ResolveUDPAddr(network, address)
	helper.HandleError(err)

	ln, err := net.ListenUDP(network, udpAddr)
	helper.HandleError(err)

	defer func(ln *net.UDPConn) {
		err := ln.Close()
		if err != nil {
			helper.HandleError(err)
		}
	}(ln)

	buffer := make([]byte, 1024)

	for {
		n, clientUdpAddr, err := ln.ReadFromUDP(buffer)
		helper.HandleError(err)

		text := bytes.ToUpper(buffer)

		fmt.Println("The received text is: " + string(text[:]) + " of length: " + strconv.Itoa(n))

		_, err = ln.WriteToUDP(text, clientUdpAddr)

		fmt.Println("Sent response back to client")

		break
	}
}
