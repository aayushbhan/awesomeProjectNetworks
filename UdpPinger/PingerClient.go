package main

import (
	helper "awesomeProject/HelperLibrary"
	"errors"
	"fmt"
	"net"
	"time"
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

	var buffer []byte
	inputText := []byte("Ping")

	for i := 0; i < 10; i++ {

		_, err = conn.Write(inputText)
		start := time.Now()
		helper.HandleError(err)

		err := conn.SetReadDeadline(time.Now().Add(1 * time.Second))
		helper.HandleError(err)

		_, err = conn.Read(buffer)
		helper.HandleError(err)
		elapsed := time.Since(start)

		if err != nil {
			var netErr net.Error
			if errors.As(err, &netErr) && netErr.Timeout() {
				fmt.Println("Taking too long to get response. Timed Out")
			}
			continue
		}

		fmt.Println("Received response from server: " + string(buffer))
		fmt.Println("The TAT was : " + elapsed.String())
	}
}
