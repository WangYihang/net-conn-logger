# net-conn-logger

[![Go Reference](https://pkg.go.dev/badge/github.com/WangYihang/net-conn-logger.svg)](https://pkg.go.dev/github.com/WangYihang/net-conn-logger)
[![Go Report Card](https://goreportcard.com/badge/github.com/WangYihang/net-conn-logger)](https://goreportcard.com/report/github.com/WangYihang/net-conn-logger)

A Go package to log network connection events by wrapping `net.Listener` and `net.Conn`.

## Installation

```bash
go get -u github.com/WangYihang/net-conn-logger
```

## Usage

Let's create a simple HTTP server that logs all connection events (`Accept`, `Read`, `Write` and `Close`).

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/WangYihang/go-net-conn-logger/pkg/listener"
)

func main() {
	// Create a new listener
	l, err := listener.NewListener(":8080")
	if err != nil {
		panic(err)
	}

	// Start the log consumer
	go func() {
		for logEntry := range l.LogChan() {
			data, _ := json.Marshal(logEntry)
			fmt.Println(string(data))
		}
	}()

	// Start server with the listener
	http.Serve(l, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	}))
}
```

```bash
$ go run main.go
```

```bash
$ curl http://localhost:8080/ -4 -H 'Host: www.example.com' --data 'a=1&b=2' 
Hello, World!
```

```json
{
    "timestamp": 1713267806921,
    "events": [
        {
            "timestamp": 1713267806921,
            "type": "accept",
            "payload": "",
            "error": ""
        },
        {
            "timestamp": 1713267806921,
            "type": "read",
            "payload": "GET / HTTP/1.1\r\nHost: localhost:8080\r\nUser-Agent: curl/8.2.1\r\nAccept: */*\r\n\r\n",
            "error": ""
        },
        {
            "timestamp": 1713267806921,
            "type": "write",
            "payload": "HTTP/1.1 200 OK\r\nDate: Tue, 16 Apr 2024 11:43:26 GMT\r\nContent-Length: 13\r\nContent-Type: text/plain; charset=utf-8\r\n\r\nHello, World!",
            "error": ""
        },
        {
            "timestamp": 1713267806921,
            "type": "read",
            "payload": "",
            "error": "read tcp 127.0.0.1:8080-\u003e127.0.0.1:48884: i/o timeout"
        },
        {
            "timestamp": 1713267806921,
            "type": "read",
            "payload": "",
            "error": "EOF"
        },
        {
            "timestamp": 1713267806921,
            "type": "close",
            "payload": "",
            "error": ""
        }
    ],
    "bytes_written": "HTTP/1.1 200 OK\r\nDate: Tue, 16 Apr 2024 11:43:26 GMT\r\nContent-Length: 13\r\nContent-Type: text/plain; charset=utf-8\r\n\r\nHello, World!",
    "bytes_read": "GET / HTTP/1.1\r\nHost: localhost:8080\r\nUser-Agent: curl/8.2.1\r\nAccept: */*\r\n\r\n",
    "local_addr": "127.0.0.1:8080",
    "remote_addr": "127.0.0.1:48884"
}
```

## Examples

1. [Capture events for HTTP client](./examples/net-http-client/)
2. [Capture events for HTTP server](./examples/net-http-server/)
3. [Capture events for Gin server](./examples/gin-server/)
4. [Capture events for SMTP server](./examples/go-smtp-server/)
