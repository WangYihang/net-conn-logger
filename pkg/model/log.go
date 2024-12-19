package model

import (
	"encoding/json"
	"net"
	"time"
)

type LogEntry struct {
	Timestamp       int64
	Events          []*Event
	BytesWritten    JSONMarshalableByteSlice
	NumBytesWritten uint64
	BytesRead       JSONMarshalableByteSlice
	NumBytesRead    uint64
	LocalAddr       net.Addr
	RemoteAddr      net.Addr
}

func NewLogEntry(localAddr, remoteAddr net.Addr) *LogEntry {
	return &LogEntry{
		Timestamp:  time.Now().UnixMilli(),
		Events:     []*Event{},
		LocalAddr:  localAddr,
		RemoteAddr: remoteAddr,
	}
}

func (l *LogEntry) AddEvent(event *Event) {
	l.Events = append(l.Events, event)
}

func (l *LogEntry) MarshalJSON() ([]byte, error) {
	bytesWritten := []byte{}
	bytesRead := []byte{}
	for _, event := range l.Events {
		switch event.Type() {
		case Write:
			bytesWritten = append(bytesWritten, event.Payload...)
			l.NumBytesWritten += uint64(len(event.Payload))
		case Read:
			bytesRead = append(bytesRead, event.Payload...)
			l.NumBytesRead += uint64(len(event.Payload))
		}
	}
	l.BytesWritten = bytesWritten
	l.BytesRead = bytesRead
	return json.Marshal(struct {
		Timestamp    int64                    `json:"timestamp"`
		Events       []*Event                 `json:"events"`
		BytesWritten JSONMarshalableByteSlice `json:"bytes_written"`
		BytesRead    JSONMarshalableByteSlice `json:"bytes_read"`
		LocalAddr    string                   `json:"local_addr"`
		RemoteAddr   string                   `json:"remote_addr"`
	}{
		Timestamp:    l.Timestamp,
		Events:       l.Events,
		BytesWritten: l.BytesWritten,
		BytesRead:    l.BytesRead,
		LocalAddr:    l.LocalAddr.String(),
		RemoteAddr:   l.RemoteAddr.String(),
	})
}
