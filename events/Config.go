package events

type Config struct {
	NatsAddress string `envconfig:"NATS_ADDRESS"`
}
