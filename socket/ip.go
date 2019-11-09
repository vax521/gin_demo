package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	if len(os.Args) != 2{
		fmt.Fprintf(os.Stderr,"useage:%s ip-addr\n",os.Args[0])
		os.Exit(0)
	}
	name := os.Args[1]
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid addess;")
	}else {
		fmt.Println("the address is ", addr.String())
	}
	os.Exit(0)
}
