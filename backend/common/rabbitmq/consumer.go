package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
)

// Consumer for receiving AMPQ events
type Consumer struct {
	conn      *amqp.Connection
	queueName string
}

func (consumer *Consumer) setup() error {
	channel, err := consumer.conn.Channel()
	if err != nil {
		return err
	}
	return declareExchange(channel)
}

// NewConsumer returns a new Consumer
func NewConsumer(conn *amqp.Connection) (Consumer, error) {
	consumer := Consumer{
		conn: conn,
	}
	err := consumer.setup()
	if err != nil {
		return Consumer{}, err
	}

	return consumer, nil
}

// Listen will listen for all new Queue publications
// and print them to the console.
func (consumer *Consumer) Listen(topics []string, messages chan string) error {
	ch, err := consumer.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := declareRandomQueue(ch)
	if err != nil {
		return err
	}

	for _, s := range topics {
		err = ch.QueueBind(
			q.Name,
			s,
			getExchangeName(),
			false,
			nil,
		)

		if err != nil {
			return err
		}
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		return err
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			messages <- string(d.Body)
		}
	}()

	fmt.Printf("[*] Waiting for message [Exchange, Queue][%s, %s].", getExchangeName(), q.Name)
	<-forever
	return nil
}