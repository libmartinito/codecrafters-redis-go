package main

import (
	"fmt"
	"net"
	"os"
	"strings"
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
		response := generateResponse(strings.ToLower(cmd), args, s)
		c.Write([]byte(response))
	}
}

func initiateMasterHandshake(host string, port string) {
	c, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		fmt.Println("Error connecting to master: ", err.Error())
		os.Exit(1)
	}

	c.Write([]byte("*1\r\n$4\r\nping\r\n"))
	c.Write([]byte("*3\r\n$8\r\nREPLCONF\r\n$14\r\nlistening-port\r\n$4\r\n6380\r\n"))
	c.Write([]byte("*3\r\n$8\r\nREPLCONF\r\n$4\r\ncapa\r\n$6\r\npsync2\r\n"))
	c.Write([]byte("*3\r\n$5\r\nPSYNC\r\n$1\r\n?\r\n$2\r\n-1\r\n"))
}
