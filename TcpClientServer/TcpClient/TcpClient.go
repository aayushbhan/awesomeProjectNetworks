package main

import (
	helper "awesomeProject/HelperLibrary"
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	var network, address string

	network, address = "tcp", "localhost:8000"

	conn, err := net.Dial(network, address)

	helper.HandleError(err)

	scanner := bufio.NewScanner(os.Stdin)
	buffer := make([]byte, 1024)

	for {

		scanner.Scan()

		serverInput := scanner.Bytes()

		_, err = conn.Write(serverInput)

		helper.HandleError(err)

		if strings.ToLower(string(serverInput)) == "exit" {
			break
		}

		n, err := conn.Read(buffer)

		helper.HandleError(err)

		serverOutput := string(buffer[:n])

		fmt.Println("Received From Server: " + serverOutput)

		if strings.ToLower(serverOutput) == "exit" {
			break
		}
	}

}
