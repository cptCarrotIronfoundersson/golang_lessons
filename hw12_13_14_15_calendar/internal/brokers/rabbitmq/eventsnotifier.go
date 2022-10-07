package rabbitmq

import (
	"context"
	"encoding/json"
	"log"

	"github.com/cptCarrotIronfoundersson/hw12_13_14_15_calendar/cmd"
	"github.com/cptCarrotIronfoundersson/hw12_13_14_15_calendar/internal/service/entity"
	amqp "github.com/rabbitmq/amqp091-go"
)

type EventsNotifier struct {
	ch      *amqp.Channel
	chQueue amqp.Queue
}

func NewEventsNotifier() EventsNotifier {
	conn, err := amqp.Dial(cmd.Config.Queue.DSN)
	if err != nil {
		log.Fatal(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	q, err := ch.QueueDeclare(
		"OTUSScheduler",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	notifier := EventsNotifier{ch, q}
	return notifier
}

func (e EventsNotifier) NotifyEvent(ctx context.Context, event entity.Event) error {
	eventJSON, _ := json.Marshal(event)
	err := e.ch.PublishWithContext(
		context.Background(),
		"",
		e.chQueue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "json/plain",
			Body:        eventJSON,
		},
	)
	if err != nil {
		return err
	}
	return nil
}
