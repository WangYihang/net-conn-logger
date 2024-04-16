package model

import (
	"encoding/json"
	"net"
	"time"
)

type LogEntry struct {
	timestamp    int64
	events       []*Event
	bytesWritten JSONMarshalableByteSlice
	bytesRead    JSONMarshalableByteSlice
	localAddr    net.Addr
	remoteAddr   net.Addr
}

func NewLogEntry(localAddr, remoteAddr net.Addr) *LogEntry {
	return &LogEntry{
		timestamp:  time.Now().UnixMilli(),
		events:     []*Event{},
		localAddr:  localAddr,
		remoteAddr: remoteAddr,
	}
}

func (l *LogEntry) AddEvent(event *Event) {
	l.events = append(l.events, event)
}

func (l *LogEntry) MarshalJSON() ([]byte, error) {
	bytesWritten := []byte{}
	bytesRead := []byte{}
	for _, event := range l.events {
		switch event.Type() {
		case WRITE:
			bytesWritten = append(bytesWritten, event.payload...)
		case READ:
			bytesRead = append(bytesRead, event.payload...)
		}
	}
	l.bytesWritten = bytesWritten
	l.bytesRead = bytesRead
	return json.Marshal(struct {
		Timestamp    int64                    `json:"timestamp"`
		Events       []*Event                 `json:"events"`
		BytesWritten JSONMarshalableByteSlice `json:"bytes_written"`
		BytesRead    JSONMarshalableByteSlice `json:"bytes_read"`
		LocalAddr    string                   `json:"local_addr"`
		RemoteAddr   string                   `json:"remote_addr"`
	}{
		Timestamp:    l.timestamp,
		Events:       l.events,
		BytesWritten: l.bytesWritten,
		BytesRead:    l.bytesRead,
		LocalAddr:    l.localAddr.String(),
		RemoteAddr:   l.remoteAddr.String(),
	})
}
