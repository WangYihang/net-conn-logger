package connection

import (
	"net"
	"sync"

	"github.com/WangYihang/go-net-conn-logger/pkg/model"
)

type LoggingConn struct {
	net.Conn
	eventChan chan *model.Event
	logEntry  *model.LogEntry
	logChan   chan *model.LogEntry
	once      sync.Once
}

type LoggingConnOption func(*LoggingConn)

func NewConn(conn net.Conn, opts ...LoggingConnOption) *LoggingConn {
	c := &LoggingConn{
		Conn:      conn,
		eventChan: make(chan *model.Event),
		logEntry:  model.NewLogEntry(conn.LocalAddr(), conn.RemoteAddr()),
		logChan:   make(chan *model.LogEntry),
		once:      sync.Once{},
	}
	go func() {
		for event := range c.eventChan {
			c.logEntry.AddEvent(event)
		}
	}()
	c.eventChan <- model.NewEvent(model.Accept, nil, nil)
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *LoggingConn) Read(b []byte) (int, error) {
	n, err := c.Conn.Read(b)
	if n > 0 {
		newBuf := make([]byte, n)
		copy(newBuf, b[:n])
		c.eventChan <- model.NewEvent(model.Read, newBuf, err)
	} else {
		c.eventChan <- model.NewEvent(model.Read, nil, err)
	}
	return n, err
}

func (c *LoggingConn) Write(b []byte) (int, error) {
	n, err := c.Conn.Write(b)
	if n > 0 {
		newBuf := make([]byte, n)
		copy(newBuf, b[:n])
		c.eventChan <- model.NewEvent(model.Write, newBuf, err)
	} else {
		c.eventChan <- model.NewEvent(model.Write, nil, err)
	}
	return n, err
}

func (c *LoggingConn) Close() error {
	err := c.Conn.Close()
	c.once.Do(func() {
		c.eventChan <- model.NewEvent(model.Close, nil, err)
		close(c.eventChan)
		c.logChan <- c.logEntry
		close(c.logChan)
	})
	return err
}

func (c *LoggingConn) LogChan() chan *model.LogEntry {
	return c.logChan
}
