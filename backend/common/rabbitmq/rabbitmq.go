package rabbitmq

import (
	"fmt"
	"os"
	"time"

	"github.com/streadway/amqp"
)

func Init() *amqp.Connection {

	name := os.Getenv("RABBITMQ_USERNAME")
	password := os.Getenv("RABBITMQ_PASSWORD")

	connectUrl := fmt.Sprintf("amqp://%s:%s@rabbitmq:5672", name, password)
	conn, e := connect(connectUrl)

	for e != nil {
		time.Sleep(3 * time.Second)
		conn, e = connect(connectUrl)
	}

	return conn
}

func connect(connectUrl string) (*amqp.Connection, error) {
	return amqp.Dial(connectUrl)
}
