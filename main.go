package main

import (
    "fmt"
    "net/http"
	"os"
	"strconv"
	"time"
	"sync"
)

func upDown(writer http.ResponseWriter, request *http.Request) {
	
	status.mu.Lock()

	if status.status {
		fmt.Fprintf(writer, "up\n")
	} else {
		fmt.Fprintf(writer, "down\n")
	}

	status.mu.Unlock()

}

func updateServer(config map[string]Server, check_timeout int, tcp_timeout int) {

	for {

		status.mu.Lock()

		for _, server := range config {
			status.status = IsUp(server.HOST, server.PORT, tcp_timeout)
		}

		status.mu.Unlock()
		time.Sleep(time.Duration(check_timeout) * time.Second)

	}
}

type safeStatus struct {
	mu sync.Mutex
	status bool
}

var status safeStatus

func main() {

	config_file := "config.yml"
	config := GetConfig(config_file)

	port := os.Args[1]

	check_timeout, err := strconv.Atoi(os.Args[2])
	LogError(err)

	tcp_timeout, err := strconv.Atoi(os.Args[3])
	LogError(err)

	go updateServer(config, check_timeout, tcp_timeout)

	for name, _ := range config {
		http.HandleFunc("/" + name, upDown)
	}

	http.ListenAndServe(":" + port, nil)

}
