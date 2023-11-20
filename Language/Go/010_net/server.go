package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}

		go handleConn(conn) // handle one connection at a time
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		buffer := make([]byte, 1024)
		n, err := c.Read(buffer)
		if err != nil {
			fmt.Printf("receiving data error\n")
			break
		}

		data := string(buffer[:n])
		fmt.Printf("received data:%s\n", data)
	}
}
