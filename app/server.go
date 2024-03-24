package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	port := flag.String("port", "6379", "port to listen on")
	replicaof := flag.String("replicaof", "", "replicate to another redis server")

	flag.Parse()

	var l net.Listener
	var err error

	s := NewStore()

	if *port != "" {
		l, err = net.Listen("tcp", "0.0.0.0:"+*port)
		if err != nil {
			fmt.Println("Failed to bind to port " + *port)
			os.Exit(1)
		}
	} else {
		l, err = net.Listen("tcp", "0.0.0.0:6379")
		if err != nil {
			fmt.Println("Failed to bind to port 6379")
			os.Exit(1)
		}
	}

	s.InitInfo(*replicaof)

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go handleConnection(c, s)
	}
}
