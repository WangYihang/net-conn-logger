package transport

import (
	"crypto/tls"
	"net"
	"net/http"

	"github.com/WangYihang/net-conn-logger/pkg/connection"
	"github.com/WangYihang/net-conn-logger/pkg/model"
)

type Transport struct {
	http.Transport
	logChan chan *model.LogEntry
}

func NewTransport() *Transport {
	transport := &Transport{
		logChan: make(chan *model.LogEntry),
	}
	transport.Transport = http.Transport{
		Dial: func(network, addr string) (net.Conn, error) {
			conn, err := net.Dial(network, addr)
			if err != nil {
				return nil, err
			}
			logConn := connection.NewConn(conn)
			go func() {
				for logEntry := range logConn.LogChan() {
					transport.logChan <- logEntry
				}
			}()
			return logConn, nil
		},
		DialTLS: func(network, addr string) (net.Conn, error) {
			conn, err := tls.Dial(network, addr, &tls.Config{})
			if err != nil {
				return nil, err
			}
			logConn := connection.NewConn(conn)
			go func() {
				for logEntry := range logConn.LogChan() {
					transport.logChan <- logEntry
				}
			}()
			return logConn, nil
		},
		DisableKeepAlives:  true,
		DisableCompression: true,
	}
	return transport
}

func (t *Transport) LogChan() chan *model.LogEntry {
	return t.logChan
}
