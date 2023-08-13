package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	var network, address string

	network, address = "udp", "localhost:2001"

	udpAdr, err := net.ResolveUDPAddr(network, address)

	conn, _ := net.DialUDP(network, nil, udpAdr)

	handleError(err)

	defer func(conn *net.UDPConn) {
		_ = conn.Close()
	}(conn)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Enter some text")

		scanner.Scan()

		inputText := scanner.Text()

		_, err = conn.Write([]byte(inputText))

		handleError(err)

		buffer := make([]byte, 1024)

		_, err = conn.Read(buffer)

		handleError(err)

		fmt.Println("Received response from server: " + string(buffer))
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
