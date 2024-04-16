#### Usage

```bash
$ go run main.go
```

```bash
$ curl localhost:8080/ -4 -H 'Host: www.example.com' --data 'a=1&b=2' 
404 page not found
```

```json
{
    "timestamp": 1713269512993,
    "events": [
        {
            "timestamp": 1713269512993,
            "type": "accept",
            "payload": "",
            "error": ""
        },
        {
            "timestamp": 1713269512994,
            "type": "read",
            "payload": "POST / HTTP/1.1\r\nHost: www.example.com\r\nUser-Agent: curl/8.2.1\r\nAccept: */*\r\nContent-Length: 7\r\nContent-Type: application/x-www-form-urlencoded\r\n\r\na=1\u0026b=2",
            "error": ""
        },
        {
            "timestamp": 1713269512994,
            "type": "write",
            "payload": "HTTP/1.1 404 Not Found\r\nContent-Type: text/plain\r\nDate: Tue, 16 Apr 2024 12:11:52 GMT\r\nContent-Length: 18\r\n\r\n404 page not found",
            "error": ""
        },
        {
            "timestamp": 1713269512994,
            "type": "read",
            "payload": "",
            "error": "read tcp 127.0.0.1:8080-\u003e127.0.0.1:50842: i/o timeout"
        },
        {
            "timestamp": 1713269512994,
            "type": "read",
            "payload": "",
            "error": "EOF"
        },
        {
            "timestamp": 1713269512995,
            "type": "close",
            "payload": "",
            "error": ""
        }
    ],
    "bytes_written": "HTTP/1.1 404 Not Found\r\nContent-Type: text/plain\r\nDate: Tue, 16 Apr 2024 12:11:52 GMT\r\nContent-Length: 18\r\n\r\n404 page not found",
    "bytes_read": "POST / HTTP/1.1\r\nHost: www.example.com\r\nUser-Agent: curl/8.2.1\r\nAccept: */*\r\nContent-Length: 7\r\nContent-Type: application/x-www-form-urlencoded\r\n\r\na=1\u0026b=2",
    "local_addr": "127.0.0.1:8080",
    "remote_addr": "127.0.0.1:50842"
}
```