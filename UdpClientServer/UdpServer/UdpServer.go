package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"strconv"
)

func main() {

	var network, address string

	network, address = "udp", "localhost:2001"

	udpAddr, _ := net.ResolveUDPAddr(network, address)

	ln, err := net.ListenUDP(network, udpAddr)

	handleError(err)

	defer func(ln *net.UDPConn) {
		_ = ln.Close()
	}(ln)

	for {

		buffer := make([]byte, 1024)

		n, _, err := ln.ReadFromUDP(buffer)

		handleError(err)

		text := bytes.ToUpper(buffer)

		fmt.Println("The received text is: " + string(text[:]) + " of length: " + strconv.Itoa(n))
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
