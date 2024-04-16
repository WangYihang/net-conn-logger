package listener

import (
	"net"

	"github.com/WangYihang/go-net-conn-logger/pkg/connection"
	"github.com/WangYihang/go-net-conn-logger/pkg/model"
)

type loggingListener struct {
	net.Listener
	logChan chan *model.LogEntry
}

func NewListener(addr string) (*loggingListener, error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	logChan := make(chan *model.LogEntry)
	return &loggingListener{
		Listener: listener,
		logChan:  logChan,
	}, nil
}

func (l *loggingListener) Accept() (net.Conn, error) {
	conn, err := l.Listener.Accept()
	if err != nil {
		return nil, err
	}
	loggingConn := connection.NewConn(conn)
	go func() {
		for logEntry := range loggingConn.LogChan() {
			l.logChan <- logEntry
		}
	}()
	return loggingConn, nil
}

func (l *loggingListener) LogChan() chan *model.LogEntry {
	return l.logChan
}
