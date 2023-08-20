package main

import (
	helper "awesomeProject/HelperLibrary"
	httpHelper "awesomeProject/HelperLibrary/HttpLibrary"
	"net"
	"os"
)

const (
	address    = "localhost:8765"
	network    = "tcp"
	bufferSize = 1024
)

func main() {
	ln, err := net.Listen(network, address)
	helper.HandleError(err)

	defer func(ln net.Listener) {
		err := ln.Close()
		helper.HandleError(err)
	}(ln)

	for {
		conn, err := ln.Accept()
		helper.HandleError(err)

		requestBuffer := make([]byte, bufferSize)
		n, err := conn.Read(requestBuffer)
		helper.HandleError(err)

		fileName := httpHelper.GetFileNameFromRequest(requestBuffer, n)
		filePath := helper.GetFilePath(fileName)

		_, err = os.Stat(filePath)
		var response string

		if err != nil {
			// Return 404
			response = httpHelper.GenerateResponse(httpHelper.NotFound, helper.GetFilePath("NotFound.html"), bufferSize)
		} else {
			// Return 200
			response = httpHelper.GenerateResponse(httpHelper.OK, filePath, bufferSize)
		}

		_, err = conn.Write([]byte(response))
		helper.HandleError(err)

		_ = conn.Close() // Close the connection
	}
}
