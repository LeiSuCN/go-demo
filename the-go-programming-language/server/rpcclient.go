package main

import (
	"log"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {

	client, err := rpc.DialHTTP("tcp", "localhost:1234")

	if err != nil {
		log.Fatal("dialing:", err)
		return
	}

	args := &Args{
		7, 8,
	}

	var reply int

	err = client.Call("Arith.Mul", args, &reply)

	if err != nil {
		log.Fatal("arith error:", err)
	}

	log.Println(reply)
}
