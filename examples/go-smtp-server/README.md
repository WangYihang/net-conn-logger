#### Usage

```bash
go run main.go
```

```bash
$ swaks --to test@example.com --from sender@example.com --server localhost:1025 --auth PLAIN --auth-user username --auth-password password        
=== Trying localhost:1025...
=== Connected to localhost.
<-  220 localhost ESMTP Service Ready
 -> EHLO ubuntu
<-  250-Hello ubuntu
<-  250-PIPELINING
<-  250-8BITMIME
<-  250-ENHANCEDSTATUSCODES
<-  250-CHUNKING
<-  250-AUTH PLAIN
<-  250-SIZE 1048576
<-  250 LIMITS RCPTMAX=50
 -> AUTH PLAIN AHVzZXJuYW1lAHBhc3N3b3Jk
<-  235 2.0.0 Authentication succeeded
 -> MAIL FROM:<sender@example.com>
<-  250 2.0.0 Roger, accepting mail from <sender@example.com>
 -> RCPT TO:<test@example.com>
<-  250 2.0.0 I'll make sure <test@example.com> gets this
 -> DATA
<-  354 Go ahead. End your data with <CR><LF>.<CR><LF>
 -> Date: Tue, 16 Apr 2024 19:50:41 +0800
 -> To: test@example.com
 -> From: sender@example.com
 -> Subject: test Tue, 16 Apr 2024 19:50:41 +0800
 -> Message-Id: <20240416195041.1690294@ubuntu>
 -> X-Mailer: swaks v20201014.0 jetmore.org/john/code/swaks/
 -> 
 -> This is a test mailing
 -> 
 -> 
 -> .
<-  250 2.0.0 OK: queued
 -> QUIT
<-  221 2.0.0 Bye
=== Connection closed with remote host.
```

```json
{
    "timestamp": 1713268241389,
    "events": [
        {
            "timestamp": 1713268241389,
            "type": "accept",
            "payload": "",
            "error": ""
        },
        {
            "timestamp": 1713268241389,
            "type": "write",
            "payload": "220 localhost ESMTP Service Ready\r\n",
            "error": ""
        },
        {
            "timestamp": 1713268241390,
            "type": "read",
            "payload": "EHLO ubuntu\r\n",
            "error": ""
        },
        {
            "timestamp": 1713268241390,
            "type": "write",
            "payload": "250-Hello ubuntu\r\n",
            "error": ""
        },
        {
            "timestamp": 1713268241390,
            "type": "write",
            "payload": "250-PIPELINING\r\n",
            "error": ""
        },
        {
            "timestamp": 1713268241390,
            "type": "write",
            "payload": "250-8BITMIME\r\n",
            "error": ""
        },
        {
            "timestamp": 1713268241390,
            "type": "write",
            "payload": "250-ENHANCEDSTATUSCODES\r\n",
            "error": ""
        },
        {
            "timestamp": 1713268241390,
            "type": "write",
            "payload": "250-CHUNKING\r\n",
            "error": ""
        },
        {
            "timestamp": 1713268241390,
            "type": "write",
            "payload": "250-AUTH PLAIN\r\n",
            "error": ""
        },
        {
            "timestamp": 1713268241390,
            "type": "write",
            "payload": "250-SIZE 1048576\r\n",
            "error": ""
        },
        {
            "timestamp": 1713268241390,
            "type": "write",
            "payload": "250 LIMITS RCPTMAX=50\r\n",
            "error": ""
        },
        {
            "timestamp": 1713268241390,
            "type": "read",
            "payload": "AUTH PLAIN AHVzZXJuYW1lAHBhc3N3b3Jk\r\n",
            "error": ""
        },
        {
            "timestamp": 1713268241390,
            "type": "write",
            "payload": "235 2.0.0 Authentication succeeded\r\n",
            "error": ""
        },
        {
            "timestamp": 1713268241390,
            "type": "read",
            "payload": "MAIL FROM:\u003csender@example.com\u003e\r\n",
            "error": ""
        },
        {
            "timestamp": 1713268241390,
            "type": "write",
            "payload": "250 2.0.0 Roger, accepting mail from \u003csender@example.com\u003e\r\n",
            "error": ""
        },
        {
            "timestamp": 1713268241390,
            "type": "read",
            "payload": "RCPT TO:\u003ctest@example.com\u003e\r\n",
            "error": ""
        },
        {
            "timestamp": 1713268241390,
            "type": "write",
            "payload": "250 2.0.0 I'll make sure \u003ctest@example.com\u003e gets this\r\n",
            "error": ""
        },
        {
            "timestamp": 1713268241390,
            "type": "read",
            "payload": "DATA\r\n",
            "error": ""
        },
        {
            "timestamp": 1713268241390,
            "type": "write",
            "payload": "354 Go ahead. End your data with \u003cCR\u003e\u003cLF\u003e.\u003cCR\u003e\u003cLF\u003e\r\n",
            "error": ""
        },
        {
            "timestamp": 1713268241391,
            "type": "read",
            "payload": "Date: Tue, 16 Apr 2024 19:50:41 +0800\r\nTo: test@example.com\r\nFrom: sender@example.com\r\nSubject: test Tue, 16 Apr 2024 19:50:41 +0800\r\nMessage-Id: \u003c20240416195041.1690294@ubuntu\u003e\r\nX-Mailer: swaks v20201014.0 jetmore.org/john/code/swaks/\r\n\r\nThis is a test mailing\r\n\r\n\r\n.\r\n",
            "error": ""
        },
        {
            "timestamp": 1713268241391,
            "type": "write",
            "payload": "250 2.0.0 OK: queued\r\n",
            "error": ""
        },
        {
            "timestamp": 1713268241391,
            "type": "read",
            "payload": "QUIT\r\n",
            "error": ""
        },
        {
            "timestamp": 1713268241391,
            "type": "write",
            "payload": "221 2.0.0 Bye\r\n",
            "error": ""
        }
    ],
    "bytes_written": "220 localhost ESMTP Service Ready\r\n250-Hello ubuntu\r\n250-PIPELINING\r\n250-8BITMIME\r\n250-ENHANCEDSTATUSCODES\r\n250-CHUNKING\r\n250-AUTH PLAIN\r\n250-SIZE 1048576\r\n250 LIMITS RCPTMAX=50\r\n235 2.0.0 Authentication succeeded\r\n250 2.0.0 Roger, accepting mail from \u003csender@example.com\u003e\r\n250 2.0.0 I'll make sure \u003ctest@example.com\u003e gets this\r\n354 Go ahead. End your data with \u003cCR\u003e\u003cLF\u003e.\u003cCR\u003e\u003cLF\u003e\r\n250 2.0.0 OK: queued\r\n221 2.0.0 Bye\r\n",
    "bytes_read": "EHLO ubuntu\r\nAUTH PLAIN AHVzZXJuYW1lAHBhc3N3b3Jk\r\nMAIL FROM:\u003csender@example.com\u003e\r\nRCPT TO:\u003ctest@example.com\u003e\r\nDATA\r\nDate: Tue, 16 Apr 2024 19:50:41 +0800\r\nTo: test@example.com\r\nFrom: sender@example.com\r\nSubject: test Tue, 16 Apr 2024 19:50:41 +0800\r\nMessage-Id: \u003c20240416195041.1690294@ubuntu\u003e\r\nX-Mailer: swaks v20201014.0 jetmore.org/john/code/swaks/\r\n\r\nThis is a test mailing\r\n\r\n\r\n.\r\nQUIT\r\n",
    "local_addr": "127.0.0.1:1025",
    "remote_addr": "127.0.0.1:47270"
}
```