package model

import "fmt"

type EventType int64

const (
	ACCEPT EventType = iota
	READ
	WRITE
	CLOSE
)

func (e EventType) String() string {
	switch e {
	case ACCEPT:
		return "accept"
	case READ:
		return "read"
	case WRITE:
		return "write"
	case CLOSE:
		return "close"
	default:
		return "unknown"
	}
}

func (e EventType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", e.String())), nil
}
