package factory

import (
	"design-patterns-go/constants"
	"fmt"
)

// basically all creational patterns are a way of writing if else in diff manner

type Notification interface {
	Send(to, message string) error
}

type EmailNotification struct{}
type SmsNotification struct{}

func (s SmsNotification) Send(to, message string) error {
	fmt.Printf("Sending SMS to %s, message: %s\n", to, message)
	return nil
}

func (e EmailNotification) Send(to, message string) error {
	fmt.Printf("Sending EMAIL to %s, message: %s\n", to, message)
	return nil
}

// Factory function
func Create(channel constants.Channel) (Notification, error) {
	switch channel {
	case constants.Email:
		return EmailNotification{}, nil
	case constants.Sms:
		return SmsNotification{}, nil
	default:
		return nil, fmt.Errorf("unknown channel: %s", channel)
	}
}
