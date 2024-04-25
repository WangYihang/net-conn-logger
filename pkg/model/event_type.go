package model

import "fmt"

type EventType int64

const (
	Undefined EventType = iota
	Accept
	Read
	Write
	Close
)

func (e EventType) String() string {
	switch e {
	case Accept:
		return "accept"
	case Read:
		return "read"
	case Write:
		return "write"
	case Close:
		return "close"
	default:
		return "unknown"
	}
}

func (e EventType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", e.String())), nil
}
