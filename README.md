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
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, World!"))
    })
    http.Serve(l, nil)
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
    "timestamp": 1713268599117,
    "local_addr": "127.0.0.1:8080",
    "remote_addr": "127.0.0.1:46942",
    "events": [
        {
            "timestamp": 1713268599117,
            "type": "accept",
            "payload": "",
            "error": ""
        },
        {
            "timestamp": 1713268599117,
            "type": "read",
            "payload": "POST / HTTP/1.1\r\nHost: www.example.com\r\nUser-Agent: curl/8.2.1\r\nAccept: */*\r\nContent-Length: 7\r\nContent-Type: application/x-www-form-urlencoded\r\n\r\na=1\u0026b=2",
            "error": ""
        },
        {
            "timestamp": 1713268599117,
            "type": "write",
            "payload": "HTTP/1.1 200 OK\r\nDate: Tue, 16 Apr 2024 11:56:39 GMT\r\nContent-Length: 13\r\nContent-Type: text/plain; charset=utf-8\r\n\r\nHello, World!",
            "error": ""
        },
        {
            "timestamp": 1713268599117,
            "type": "read",
            "payload": "",
            "error": "EOF"
        },
        {
            "timestamp": 1713268599117,
            "type": "read",
            "payload": "",
            "error": "EOF"
        }
    ],
    "bytes_written": "HTTP/1.1 200 OK\r\nDate: Tue, 16 Apr 2024 11:56:39 GMT\r\nContent-Length: 13\r\nContent-Type: text/plain; charset=utf-8\r\n\r\nHello, World!",
    "bytes_read": "POST / HTTP/1.1\r\nHost: www.example.com\r\nUser-Agent: curl/8.2.1\r\nAccept: */*\r\nContent-Length: 7\r\nContent-Type: application/x-www-form-urlencoded\r\n\r\na=1\u0026b=2"
}
```

## Examples

1. [Capture events for HTTP client](./examples/net-http-client/)
2. [Capture events for HTTP server](./examples/net-http-server/)
3. [Capture events for Gin server](./examples/gin-server/)
4. [Capture events for SMTP server](./examples/go-smtp-server/)
