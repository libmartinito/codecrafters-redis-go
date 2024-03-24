package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func generateEchoResponse(s string) string {
	return generateBulkString(s)
}

func generateGetResponse(k string, s *Store) string {
	now := time.Now().UnixNano() / 1e6

	if s.data[k].expiry == -1 || s.data[k].expiry >= now {
		return generateBulkString(s.Get(k))
	} else {
		return generateNullBulkString()
	}
}

func generateInfoResponse(s *Store) string {
	return generateBulkString("role:" + s.info["replication"]["role"])
}

func generatePingResponse() string {
	return generateSimpleString("PONG")
}

func generateResponse(cmd string, args []string, s *Store) string {
	if cmd == "ping" {
		return generatePingResponse()
	} else if cmd == "echo" {
		return generateEchoResponse(args[0])
	} else if cmd == "set" {
		return generateSetResponse(args, s)
	} else if cmd == "get" {
		return generateGetResponse(args[0], s)
	} else if cmd == "info" {
		return generateInfoResponse(s)
	}

	return ""
}

func generateSetResponse(args []string, s *Store) string {
	if contains(args, "px") {
		expiryDurationMilli, err := strconv.Atoi(args[3])
		if err != nil {
			fmt.Println("Error converting expiry to int: ", err.Error())
			os.Exit(1)
		}

		expiry := time.Now().UnixNano()/1e6 + int64(expiryDurationMilli)
		s.SetWithExpiry(args[0], args[1], expiry)
	} else {
		s.Set(args[0], args[1])
	}

	return generateSimpleString("OK")
}
