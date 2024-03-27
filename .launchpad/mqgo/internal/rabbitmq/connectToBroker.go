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
	conn, err := amqp.Dial("amqps://pqzzjkrf:gx_v4To044J7vJnnMsRBVl1m09tJJhMU@puffin.rmq2.cloudamqp.com/pqzzjkrf")
	internal.HandleError(err, "Failed to connect to RabbitMQ")
	log.Info("Connected To RabbitMQ")
	return conn
}
