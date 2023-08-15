package main

import (
	helper "awesomeProject/HelperLibrary"
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	network, address := "udp", "localhost:2001"

	udpAdr, err := net.ResolveUDPAddr(network, address)
	helper.HandleError(err)

	conn, _ := net.DialUDP(network, nil, udpAdr)
	helper.HandleError(err)

	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			helper.HandleError(err)
		}
	}(conn)

	scanner := bufio.NewScanner(os.Stdin)
	buffer := make([]byte, 1024)

	for {
		fmt.Println("Enter some text")

		scanner.Scan()
		inputText := scanner.Text()

		_, err = conn.Write([]byte(inputText))
		helper.HandleError(err)

		_, err = conn.Read(buffer)
		helper.HandleError(err)

		fmt.Println("Received response from server: " + string(buffer))

		break
	}
}
