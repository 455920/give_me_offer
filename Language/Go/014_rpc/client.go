package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Dial error:", err)
	}

	args := &Args{A: 10, B: 5}
	reply := new(Reply)

	err = client.Call("MathService.Add", args, reply)
	if err != nil {
		log.Fatal("Add error:", err)
	}

	fmt.Println("Add result:", reply.Result)

	err = client.Call("MathService.Divide", args, reply)
	if err != nil {
		log.Fatal("Divide error:", err)
	}

	fmt.Println("Divide result:", reply.Result)
}
