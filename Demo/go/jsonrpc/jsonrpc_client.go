package main

import (
	"fmt"
	"net"
	"net/rpc/jsonrpc"
)

type Args struct {
	A, B int
}

func main() {
	conn, err := net.Dial("tcp", ":12346")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := jsonrpc.NewClient(conn)
	defer client.Close()
	var res float64
	err = client.Call("DemoService.Div", Args{3, 4}, &res)
	if err != nil {
		panic(err)
	}
	fmt.Println("res is ", res)

}
