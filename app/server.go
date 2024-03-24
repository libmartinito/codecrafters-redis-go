package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func parseResp(b []byte) (cmd string, args []string) {
	respElementsArr := strings.Split(string(b), "\r\n")

	for _, v := range respElementsArr {
		if v == "" {
			continue
		}

		if v[0] == '*' {
			continue
		}

		if v[0] == '$' {
			continue
		}

		if cmd == "" {
			cmd = v
		} else {
			args = append(args, v)
		}
	}

	return cmd, args
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	c, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	b := make([]byte, 1024)
	_, err = c.Read(b)
	if err != nil {
		fmt.Println("Error reading from connection: ", err.Error())
		os.Exit(1)
	}

	cmd, _ := parseResp(b)
	if cmd == "ping" {
		c.Write([]byte("+PONG\r\n"))
	}
}
