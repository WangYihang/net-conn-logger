#### Usage

```bash
go run main.go
```

```json
{
    "timestamp": 1713267869383,
    "events": [
        {
            "timestamp": 1713267869383,
            "type": "accept",
            "payload": "",
            "error": ""
        },
        {
            "timestamp": 1713267869383,
            "type": "read",
            "payload": "GET / HTTP/1.1\r\nHost: localhost:8080\r\nUser-Agent: curl/8.2.1\r\nAccept: */*\r\n\r\n",
            "error": ""
        },
        {
            "timestamp": 1713267869384,
            "type": "write",
            "payload": "HTTP/1.1 200 OK\r\nContent-Type: text/plain; charset=utf-8\r\nDate: Tue, 16 Apr 2024 11:44:29 GMT\r\nContent-Length: 13\r\n\r\nHello, World!",
            "error": ""
        },
        {
            "timestamp": 1713267869384,
            "type": "read",
            "payload": "",
            "error": "EOF"
        },
        {
            "timestamp": 1713267869384,
            "type": "read",
            "payload": "",
            "error": "EOF"
        }
    ],
    "bytes_written": "HTTP/1.1 200 OK\r\nContent-Type: text/plain; charset=utf-8\r\nDate: Tue, 16 Apr 2024 11:44:29 GMT\r\nContent-Length: 13\r\n\r\nHello, World!",
    "bytes_read": "GET / HTTP/1.1\r\nHost: localhost:8080\r\nUser-Agent: curl/8.2.1\r\nAccept: */*\r\n\r\n",
    "local_addr": "127.0.0.1:8080",
    "remote_addr": "127.0.0.1:59860"
}
```