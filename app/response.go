package main

import "fmt"

func generateBulkString(s string) string {
	return "$" + fmt.Sprint(len(s)) + "\r\n" + s + "\r\n"
}

func generateSimpleString(s string) string {
	return "+" + s + "\r\n"
}

func generateResponse(cmd string, args []string, s *Store) string {
	if cmd == "ping" {
		return "+PONG\r\n"
	} else if cmd == "echo" {
		return generateBulkString(args[0])
	} else if cmd == "set" {
		s.Set(args[0], args[1])
		return generateSimpleString("OK")
	} else if cmd == "get" {
		s.Get(args[0])
		return generateBulkString(s.Get(args[0]))
	}

	return ""
}
