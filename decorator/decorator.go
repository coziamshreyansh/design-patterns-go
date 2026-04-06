package decorator

import "fmt"

// base class
type Notifier interface {
	Send(message string)
}

type MessageNotifier struct{}

func (m *MessageNotifier) Send(message string) {
	println("Sending msg: ", message)
}

// base decorator class
type NotifierDecorator struct {
	Wrapped Notifier
}

type EmailDecorator struct {
	NotifierDecorator
}

func (e *EmailDecorator) Send(message string) {
	e.Wrapped.Send(message)
	fmt.Println("Sending email also: ", message)
}

type SmsDecorator struct {
	NotifierDecorator
}

func (e *SmsDecorator) Send(message string) {
	e.Wrapped.Send(message)
	fmt.Println("Sending sms also: ", message)
}
