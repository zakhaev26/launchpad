package rabbitmq

import (
	"github.com/charmbracelet/log"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/zakhaev26/launchpad/internal"
)

/*
The connection abstracts the socket connection, and takes care of protocol version negotiation and authentication and so on for us.
*/
func ConnectToQueue() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	internal.HandleError(err, "Failed to connect to RabbitMQ")
	log.Info("Connected To RabbitMQ")
	return conn
}
