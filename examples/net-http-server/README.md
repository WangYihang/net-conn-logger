#### Usage

```bash
go run main.go
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