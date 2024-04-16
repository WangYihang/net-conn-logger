package model

import (
	"encoding/json"
	"time"
)

type Event struct {
	timestamp int64
	eventType EventType
	payload   JSONMarshalableByteSlice
	errString string
}

func NewEvent(eventType EventType, payload []byte, err error) *Event {
	errorString := ""
	if err != nil {
		errorString = err.Error()
	}
	return &Event{
		timestamp: time.Now().UnixMilli(),
		eventType: eventType,
		payload:   payload,
		errString: errorString,
	}
}

func (e *Event) Type() EventType {
	return e.eventType
}

func (e *Event) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp int64                    `json:"timestamp"`
		Type      EventType                `json:"type"`
		Payload   JSONMarshalableByteSlice `json:"payload"`
		Error     string                   `json:"error"`
	}{
		Timestamp: e.timestamp,
		Type:      e.eventType,
		Payload:   e.payload,
		Error:     e.errString,
	})
}
