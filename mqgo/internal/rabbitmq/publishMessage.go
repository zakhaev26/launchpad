package rabbitmq

import (
	"context"
	"time"

	"github.com/charmbracelet/log"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/zakhaev26/launchpad/internal"
)

func PublishMessage(ch *amqp.Channel, q amqp.Queue, latencyChan chan time.Time) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "Punit Saswat saare gay hai"

	err := ch.PublishWithContext(ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	internal.HandleError(err, "Failed to Publish a message")
	log.Info("PUBLISHED", "message", body)
	latencyChan <- time.Now()
}
