package main

import (
	"log"
	"net"
	"net/rpc"
)

func main() {
	mathService := new(MathService)
	rpc.Register(mathService)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Listen error:", err)
	}

	log.Println("RPC server is listening on port 1234")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go rpc.ServeConn(conn)
	}
}
