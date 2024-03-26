package rabbitmq

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/zakhaev26/launchpad/internal"
)

func DeclareQueue(ch *amqp.Channel, queueName string) amqp.Queue {
	q, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	internal.HandleError(err, fmt.Sprintf("Failed to Declare queue %s on %v channel", queueName, ch))

	return q
}

//Since publishings are asynchronous, any undeliverable message will get returned by the server. Add a listener with Channel.NotifyReturn to handle any undeliverable message when calling publish with either the mandatory or immediate parameters as true.
