package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/WangYihang/go-net-conn-logger/pkg/transport"
)

func main() {
	// Create a wait group
	wg := &sync.WaitGroup{}
	defer wg.Wait()
	wg.Add(1)

	// Create a new transport
	transport := transport.NewTransport()
	go func() {
		for logEntry := range transport.LogChan() {
			data, _ := json.Marshal(logEntry)
			fmt.Println(string(data))
			wg.Done()
		}
	}()

	// Create a new client with the transport
	client := &http.Client{Transport: transport}

	// Make a request
	resp, err := client.Get("https://www.baidu.com/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Read the response body
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
}
