package rabbitmq

import (
	fileLog "log"
	"os"
	"sync"
	"time"

	"github.com/charmbracelet/log"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/zakhaev26/launchpad/internal"
)

func ConsumeMessage(ch *amqp.Channel, q amqp.Queue, latencyChan chan time.Time) {
	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	internal.HandleError(err, "Failed to Register a consumer ")

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for d := range msgs {
			log.Info("RECEIVED", "message", string(d.Body))
			var publishTime time.Time = <-latencyChan
			log.Info("LATENCY", time.Since(publishTime), "wow!")

			logFile, err := os.OpenFile("../../internal/logs/log_gen.erl", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatal(err)
			}
			defer logFile.Close()
			fileLog.SetOutput(logFile)
			fileLog.Printf("%s", time.Since(publishTime))
		}
	}()
	wg.Wait()
}
