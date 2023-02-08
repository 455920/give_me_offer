package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type DemoService struct{}

type Args struct {
	A, B int
}

func (DemoService) Div(args Args, result *float64) error {
	if args.B == 0 {
		return fmt.Errorf("div by zero")
	}
	*result = float64(args.A) / float64(args.B)
	return nil
}

//telnet localhost 123457
//回车后输入 {"method":"DemoService.Div","Params":[{"A":8,"B":6}],"id":12} 然后再回车
//{"id":12,"result":1.3333333333333333,"error":null}

func main() {
	rpc.Register(DemoService{})
	listener, err := net.Listen("tcp", ":12346")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		fmt.Printf("accept a conn")
		if err != nil {
			log.Printf("%v\n", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}
