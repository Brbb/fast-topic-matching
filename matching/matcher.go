package matching

import "fmt"

const (
	delimiter = "."
	wildcard  = "*"
	empty     = ""
)

// Subscriber is a value associated with a subscription.
type Subscriber interface {
}

type NotificationType string

const (
	Http      NotificationType = "Http"
	WebSocket                  = "WS"
)

type Fascriber struct {
	Notification NotificationType
	ClientId     string
	// probably extended configuration like endpoints
}

func (f Fascriber) ConfirmSubscription() {
	fmt.Printf("%s will be notified through %s\n", f.ClientId, f.Notification)
}

// Subscription represents a topic subscription.
type Subscription struct {
	id         uint32
	topic      string
	subscriber Subscriber
}

// Matcher contains topic subscriptions and performs matches on them.
type Matcher interface {
	// Subscribe adds the Subscriber to the topic and returns a Subscription.
	Subscribe(topic string, sub Subscriber) (*Subscription, error)

	// Unsubscribe removes the Subscription.
	Unsubscribe(sub *Subscription)

	// Lookup returns the Subscribers for the given topic.
	Lookup(topic string) []Subscriber
}
