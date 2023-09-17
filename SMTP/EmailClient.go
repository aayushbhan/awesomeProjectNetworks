package main

import (
	helper "awesomeProject/HelperLibrary"
	"bufio"
	"encoding/base64"
	"fmt"
	"net"
)

func main() {
	server := "smtp.gmail.com"
	port := 587
	sender := "test@gmail.com"
	recipient := "test@gmail.com"

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", server, port))
	helper.HandleError(err)

	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			helper.HandleError(err)
		}
	}(conn)

	reader := bufio.NewReader(conn)

	ehlo := "EHLO host\r\n"
	sendSmtpCommand(conn, reader, ehlo)

	authPlain := "\x00" + sender + "\x00" + "password"
	authCommand := "AUTH PLAIN " + base64.StdEncoding.EncodeToString([]byte(authPlain)) + "\r\n"
	sendSmtpCommand(conn, reader, authCommand)

	mailFrom := "MAIL FROM:" + sender + "\r\n"
	sendSmtpCommand(conn, reader, mailFrom)

	rcptTo := "RCPT TO:" + recipient + "\r\n"
	sendSmtpCommand(conn, reader, rcptTo)

	data := "DATA\r\n"
	sendSmtpCommand(conn, reader, data)

	emailMessage := "Subject: Test Subject\r\nTo:" + recipient + "\r\nFrom:" + sender + "\r\n\r\nThis is the body of the email.\r\n.\r\n"
	sendSmtpCommand(conn, reader, emailMessage)

	quit := "QUIT\r\n"
	sendSmtpCommand(conn, reader, quit)
}

func sendSmtpCommand(conn net.Conn, reader *bufio.Reader, command string) {
	_, err := conn.Write([]byte(command))
	helper.HandleError(err)
	response, err := reader.ReadString('\n')
	helper.HandleError(err)
	fmt.Println(fmt.Sprintf("Server response after %s:", command), response)
}
