package main

import (
	"time"

	"github.com/zakhaev26/launchpad/internal/rabbitmq"
)

func main() {
	conn := rabbitmq.ConnectToQueue()
	defer conn.Close()

	ch := rabbitmq.CreateChannel(conn)
	q := rabbitmq.DeclareQueue(ch, "loml1")

	var latencyChannel chan time.Time = make(chan time.Time)

	go func() {
		for {
			rabbitmq.PublishMessage(ch, q, latencyChannel)
		}
	}()
	rabbitmq.ConsumeMessage(ch, q, latencyChannel)
}
