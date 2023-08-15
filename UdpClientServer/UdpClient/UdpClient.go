package main

import (
	helper "awesomeProject/HelperLibrary"
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	var network, address string

	network, address = "udp", "localhost:2001"

	udpAdr, err := net.ResolveUDPAddr(network, address)

	conn, _ := net.DialUDP(network, nil, udpAdr)

	helper.HandleError(err)

	defer func(conn *net.UDPConn) {
		_ = conn.Close()
	}(conn)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Enter some text")

		scanner.Scan()

		inputText := scanner.Text()

		_, err = conn.Write([]byte(inputText))

		helper.HandleError(err)

		buffer := make([]byte, 1024)

		_, err = conn.Read(buffer)

		helper.HandleError(err)

		fmt.Println("Received response from server: " + string(buffer))
	}
}
