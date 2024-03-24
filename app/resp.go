package main

import (
	"fmt"
	"strings"
)

func generateBulkString(s string) string {
	return "$" + fmt.Sprint(len(s)) + "\r\n" + s + "\r\n"
}

func generateNullBulkString() string {
	return "$-1\r\n"
}

func generateSimpleString(s string) string {
	return "+" + s + "\r\n"
}

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
