package events

type Message interface {
	Type() string
}
