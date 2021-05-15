package event

import "time"

type Message interface {
	Key() string
}

type MeowCreatedMessage struct {
	ID        string
	Body      string
	CreayedAt time.Time
}

func (m *MeowCreatedMessage) Key() string {
	return "meow.created"
}
