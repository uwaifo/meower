package event

import (
 	"github.com/nats-io/nats.go"
)

type NatsEventStore struct {
	nc                      *nats.Conn
	meowCreatedSubscription *nats.Subscription
	meowCreatedChan         chan MeowCreatedMessage
	//meowCreateSubscriptio *nats.S
}

func NewNats(url string) (*NatsEventStore, error) {
	nc, err := nats.Connect(url)
	if err != nil {

	}
	return &NatsEventStore{nc: nc}, nil

}
