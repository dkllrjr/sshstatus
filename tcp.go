package main

import (
    "fmt"
    "net"
	"time"
	"os"
	"strconv"
)

func IsOpen(host string, port string, sec int) bool {

	timeout := time.Duration(sec) * time.Second

	conn, err := net.DialTimeout("tcp", host + ":" + port, timeout)
	if err != nil {
		return false
	}

	if conn != nil {
		conn.Close()
		return true
	}

	return false
}

func main() {

	host, port := os.Args[1], os.Args[2]
	sec, _ := strconv.Atoi(os.Args[3])

	fmt.Println(IsOpen(host, port, sec))
}
