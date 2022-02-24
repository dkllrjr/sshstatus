package main

import (
    "fmt"
    "net/http"
	"os"
	"strconv"
	"time"
	"sync"
)

func updateServer(config map[string]Server, check_timeout int, tcp_timeout int) {

	for {

		status.mu.Lock()

		for name, server := range config {
			name := name
			server := server
			if IsUp(server.HOST, server.PORT, tcp_timeout) {
				status.updown[name] = "up"
			} else {
				status.updown[name] = "down"
			}
		}

		status.mu.Unlock()
		time.Sleep(time.Duration(check_timeout) * time.Second)

	}
}

type safeStatus struct {
	mu sync.RWMutex
	updown map[string]string
}

var status safeStatus

func main() {

	status.updown = make(map[string]string)

	config_file := os.Args[1]
	config := GetConfig(config_file)

	port := os.Args[2]

	check_timeout, err := strconv.Atoi(os.Args[3])
	LogError(err)

	tcp_timeout, err := strconv.Atoi(os.Args[4])
	LogError(err)

	go updateServer(config, check_timeout, tcp_timeout)

	for name, _ := range config {
		name := name

		http.HandleFunc("/" + name, func(writer http.ResponseWriter, request *http.Request) {
	
			status.mu.RLock()
			defer status.mu.RUnlock()
			fmt.Fprintf(writer, status.updown[name] + "\n")
		})
	}

	http.ListenAndServe(":" + port, nil)

}
