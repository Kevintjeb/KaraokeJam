package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

// Emitter for publishing AMQP events
type Emitter struct {
	connection *amqp.Connection
}

func (e *Emitter) setup() error {
	channel, err := e.connection.Channel()
	if err != nil {
		panic(err)
	}

	defer channel.Close()
	return declareExchange(channel)
}

// Push (Publish) a specified message to the AMQP exchange
func (e *Emitter) Push(event string, topic string) error {
	channel, err := e.connection.Channel()
	if err != nil {
		return err
	}

	defer channel.Close()
	log.Printf("Sending message: %s -> %s", event, getExchangeName())

	err = channel.Publish(
		getExchangeName(),
		topic,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(event),
		},
	)
	return nil
}

// NewEventEmitter returns a new event.Emitter object
// ensuring that the object is initialised, without error
func NewEmitter(conn *amqp.Connection) (Emitter, error) {
	emitter := Emitter{
		connection: conn,
	}

	err := emitter.setup()
	if err != nil {
		return Emitter{}, err
	}

	return emitter, nil
}
