package main

import (
	"fmt"
	"strings"
)

func parseResp(b []byte) (cmd string, args []string) {
	ignorePrefixes := []string{"*", "$"}
	respElements := strings.Split(string(b), "\r\n")

	for _, v := range respElements {
		if v == "" || contains(ignorePrefixes, v[0:1]) {
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

func generateBulkString(s string) string {
	return "$" + fmt.Sprint(len(s)) + "\r\n" + s + "\r\n"
}

func generateResponse(cmd string, args []string) string {
	if cmd == "ping" {
		return "+PONG\r\n"
	} else if cmd == "echo" {
		return generateBulkString(args[0])
	}

	return ""
}
