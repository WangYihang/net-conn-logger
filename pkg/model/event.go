package model

import (
	"encoding/json"
	"time"
)

type Event struct {
	Timestamp int64
	EventType EventType
	Payload   JSONMarshalableByteSlice
	Error     string
}

func NewEvent(eventType EventType, payload []byte, err error) *Event {
	errorString := ""
	if err != nil {
		errorString = err.Error()
	}
	return &Event{
		Timestamp: time.Now().UnixMilli(),
		EventType: eventType,
		Payload:   payload,
		Error:     errorString,
	}
}

func (e *Event) Type() EventType {
	return e.EventType
}

func (e *Event) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp int64                    `json:"timestamp"`
		Type      EventType                `json:"type"`
		Payload   JSONMarshalableByteSlice `json:"payload"`
		Error     string                   `json:"error"`
	}{
		Timestamp: e.Timestamp,
		Type:      e.EventType,
		Payload:   e.Payload,
		Error:     e.Error,
	})
}
