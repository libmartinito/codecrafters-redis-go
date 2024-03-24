package main

import (
	"fmt"
	"net"
	"os"
)

func handleConnection(c net.Conn) {
	b := make([]byte, 1024)

	for {
		_, err := c.Read(b)
		if err != nil {
			fmt.Println("Error reading from connection: ", err.Error())
			os.Exit(1)
		}

		cmd, _ := parseResp(b)
		response := generateResponse(cmd, nil)
		c.Write([]byte(response))
	}
}
