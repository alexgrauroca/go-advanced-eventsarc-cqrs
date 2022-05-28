package eventsrepository

type Message interface {
	Type() string
}
