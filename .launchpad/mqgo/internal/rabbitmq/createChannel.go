package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/zakhaev26/launchpad/internal"
)

func CreateChannel(conn *amqp.Connection) *amqp.Channel {
	// Channel opens a unique, concurrent server channel to process the bulk of AMQP messages. Any error from methods on this receiver will render the receiver invalid and a new Channel should be opened
	ch, err := conn.Channel()
	internal.HandleError(err, "Failed to open a Channel on RabbitMQ")

	return ch
}
