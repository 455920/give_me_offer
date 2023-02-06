package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

func main() {
	client, err := jsonrpc.Dial("tcp", "localhost:12345")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Call("Calculator.Sub", "1|2", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
