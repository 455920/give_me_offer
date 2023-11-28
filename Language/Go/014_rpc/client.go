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

	rqst := &Request{A: 10, B: 5}
	rsp := new(Respond)

	err = client.Call("MathService.Add", rqst, rsp)
	if err != nil {
		log.Fatal("Add error:", err)
	}

	fmt.Println("Add result:", rsp.Result)

	err = client.Call("MathService.Divide", rqst, rsp)
	if err != nil {
		log.Fatal("Divide error:", err)
	}

	fmt.Println("Divide result:", rsp.Result)
}
