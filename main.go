package main

import (
	"fmt"
	"gopractice/public"
	"gopractice/server"
	"io"
	"net"
	"net/http"
	"time"
)

func readTcpData(ln net.Conn) {

	for {
		ln.SetReadDeadline(time.Now().Add(time.Second * 30))
		data := make([]byte, 1024)

		_, err := ln.Read(data)
		if err != nil {
			if err != io.EOF {
				ln.Close()
				break
			}
		}

		pathName := server.ParsePathName(data)
		fileReader := public.ParsePublic(pathName)

		if fileReader == nil {
			response := "<h1>Error 404 | File Not Found</h1>"
			ctx := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/html\r\nContent-Length: %d\r\nConnection: keep-alive\r\nKeep-Alive: timeout=30, max=100\r\n\r\n%s", len(response), response)
			ln.Write([]byte(ctx))
		} else {
			readData := make([]byte, 1024)
			isFirstReq := true

			var contentType *string
			for {
				n, fileErr := fileReader.Read(readData)

				if contentType == nil {
					bytesData := make([]byte, 1024)
					copy(bytesData, readData)
					cType := http.DetectContentType(bytesData)
					contentType = &cType
				}

				htmlContentType := *contentType

				if fileErr != nil {
					if fileErr == io.EOF {
						ln.Write([]byte("0\r\n\r\n"))
						fileReader.Close()
						break
					}

					response := "<h1>Error 404 | File Not Found</h1>"
					ctx := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/html\r\nContent-Length: %d\r\nConnection: keep-alive\r\nKeep-Alive: timeout=30, max=100\r\n\r\n%s", len(response), response)
					ln.Write([]byte(ctx))
					fileReader.Close()
					break
				}

				if isFirstReq {
					byteData := readData[:n]
					ctx := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: %s\r\nTransfer-Encoding: chunked\r\nConnection: keep-alive\r\nKeep-Alive: timeout=30, max=100\r\n\r\n%x\r\n%s\r\n", htmlContentType, len(byteData), string(byteData))
					ln.Write([]byte(ctx))
					isFirstReq = false
				} else {
					byteData := readData[:n]
					ctx := fmt.Sprintf("%x\r\n%s\r\n", len(byteData), string(byteData))
					ln.Write([]byte(ctx))
				}
			}
		}

	}
}

func main() {

	port := 3000

	conn, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Server is Running on Port ", port)
	for {
		ln, err := conn.Accept()
		if err != nil {
			fmt.Println("Something Wrong with the Connection")
			continue
		}

		fmt.Println("Received New User ", ln.RemoteAddr())
		go readTcpData(ln)
	}

}
