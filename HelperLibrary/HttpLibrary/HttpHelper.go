package HttpLibrary

import (
	helper "awesomeProject/HelperLibrary"
	"os"
	"strconv"
	"strings"
	"time"
)

type HTTPStatusCode int

const (
	OK       HTTPStatusCode = 200
	NotFound HTTPStatusCode = 404
)

func GenerateResponse(statusCode HTTPStatusCode, filePath string, bufferSize int) string {
	file, err := os.Open(filePath)
	helper.HandleError(err)

	fileBuffer := make([]byte, bufferSize)
	n, err := file.Read(fileBuffer)
	helper.HandleError(err)

	response := "HTTP/1.1 " + strconv.Itoa(int(statusCode)) + "\n"
	response += "Date: " + time.Now().Format(time.RFC1123) + "\n"
	response += "Content-Length: " + strconv.Itoa(n) + "\n"
	response += "\n"
	response += string(fileBuffer)
	return response
}

func GetFileNameFromRequest(requestBuffer []byte, n int) string {
	header := strings.SplitN(string(requestBuffer[:n]), "\r\n", 2)[0]
	fileName := strings.SplitN(header, " ", 3)[1]
	return fileName
}
