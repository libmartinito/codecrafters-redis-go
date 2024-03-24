package main

import (
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
