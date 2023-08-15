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
	network, address := "tcp", "localhost:8000"

	ln, err := net.Listen(network, address)
	helper.HandleError(err)

	defer func(conn net.Listener) {
		err := conn.Close()
		if err != nil {
			helper.HandleError(err)
		}
	}(ln)

	conn, err := ln.Accept()
	helper.HandleError(err)

	scanner := bufio.NewScanner(os.Stdin)
	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		helper.HandleError(err)

		clientOutput := string(buffer[:n])

		fmt.Println("Received From Client: " + string(clientOutput))

		if strings.ToLower(clientOutput) == "exit" {
			break
		}

		scanner.Scan()
		clientInput := scanner.Bytes()

		_, err = conn.Write(clientInput)
		helper.HandleError(err)

		if strings.ToLower(string(clientInput)) == "exit" {
			break
		}
	}
}
