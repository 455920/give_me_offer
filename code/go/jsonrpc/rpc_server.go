package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"strconv"
	"strings"
)

type Calculator struct{}

func (p *Calculator) Add(request string, reply *string) error {
	arr := strings.Split(request, "|")
	if len(arr) != 2 {
		*reply = "data format error a|b"
		return nil
	}

	a, err := strconv.Atoi(arr[0])
	if err != nil {
		return err
	}
	b, err := strconv.Atoi(arr[1])
	if err != nil {
		return err
	}

	*reply = "result:" + strconv.Itoa(a+b)
	return nil
}

func (p *Calculator) Sub(request string, reply *string) error {
	arr := strings.Split(request, "|")
	if len(arr) != 2 {
		*reply = "data format error a|b"
		return nil
	}

	a, err := strconv.Atoi(arr[0])
	if err != nil {
		return err
	}
	b, err := strconv.Atoi(arr[1])
	if err != nil {
		return err
	}

	*reply = "result:" + strconv.Itoa(a-b)
	return nil
}

func (p *Calculator) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func main() {
	rpc.RegisterName("Calculator", new(Calculator))

	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for true {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go func(conn net.Conn) {
			jsonrpc.ServeConn(conn)
		}(conn)
	}
}
