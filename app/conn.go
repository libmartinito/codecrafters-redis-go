package main

import (
	"fmt"
	"net"
	"os"
)

func handleConnection(c net.Conn, s *Store) {
	b := make([]byte, 1024)

	for {
		_, err := c.Read(b)
		if err != nil {
			if err.Error() == "EOF" {
				return
			}

			fmt.Println("Error reading from connection: ", err.Error())
			os.Exit(1)
		}

		cmd, args := parseResp(b)
		response := generateResponse(cmd, args, s)
		c.Write([]byte(response))
	}
}
