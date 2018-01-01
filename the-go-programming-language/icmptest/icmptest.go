package main

import (
	"net"
	"os"
	"bytes"
	"fmt"
)

func main() {
	service := os.Args[1]

	conn,err := net.Dial("ip4:icmp", service)

	
}
