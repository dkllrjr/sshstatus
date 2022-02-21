package main

import (
    "net"
	"time"
)

func IsUp(host string, port string, sec int) bool {

	timeout := time.Duration(sec) * time.Second

	conn, err := net.DialTimeout("tcp", host + ":" + port, timeout)
	LogNotError(err)

	if conn != nil {
		conn.Close()
		return true
	}

	return false
}

