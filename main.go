package main

import (
	"fmt"
	"gopractice/public"
	"gopractice/server"
	"io"
	"net"
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
		pageFile := public.ParsePublic(pathName)
		ctx := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/html\r\nContent-Length: %d\r\nConnection: keep-alive\r\nKeep-Alive: timeout=30, max=100\r\n\r\n%s", len(pageFile), pageFile)

		ln.Write([]byte(ctx))

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
