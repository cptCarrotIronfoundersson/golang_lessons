package main

import (
	"context"
	"github.com/cptCarrotIronfoundersson/hw12_13_14_15_calendar/cmd"
	"github.com/cptCarrotIronfoundersson/hw12_13_14_15_calendar/configs/config"
	"github.com/cptCarrotIronfoundersson/hw12_13_14_15_calendar/internal/brokers/rabbitmq"
	"github.com/cptCarrotIronfoundersson/hw12_13_14_15_calendar/internal/logger"
	"github.com/cptCarrotIronfoundersson/hw12_13_14_15_calendar/internal/service/entity"
	sqlstorage "github.com/cptCarrotIronfoundersson/hw12_13_14_15_calendar/internal/storage/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func init() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

type Queue interface {
	NotifyEvent(ctx context.Context, events entity.Event) error
}

type TasksManager struct {
	config  *config.Config
	logger  *logger.Logger
	storage *sqlstorage.Storage
	events  chan entity.Event
	mu      *sync.Mutex
	rabbitmq.EventsNotifier
}

func (t *TasksManager) Stop() {

}

func (t *TasksManager) deleteOldEvents() {
	ctx := context.Background()
	if err := t.storage.DeleteOldEvents(ctx); err != nil {
		t.logger.Error(err)
	}
}

func (t *TasksManager) getEventsForNotify() []entity.Event {
	ctx := context.Background()
	events, err := t.storage.AllEvents(ctx)
	if err != nil {
		t.logger.Error(err)
	}
	var eventsToSend []entity.Event
	for _, event := range events {
		if event.StartDatetime.Add(-event.RemindTimeBefore).After(time.Now()) && event.EndDatetime.Before(event.EndDatetime) {
			eventsToSend = append(eventsToSend, event)
		}
	}
	return eventsToSend
}

func (t *TasksManager) EventNotifyWriter() {
	var events = make(map[string]entity.Event)
	for {
		for _, event := range t.getEventsForNotify() {
			uuid := event.UUID.String()
			if _, ok := events[uuid]; !ok {
				t.mu.Lock()
				events[event.UUID.String()] = event
				t.mu.Unlock()
				t.events <- event
			}
		}
		time.Sleep(time.Minute * 1)
	}
}

func (t *TasksManager) EventsNotifyReader() {
	for {
		event, ok := <-t.events
		if !ok {
			return
		}
		err := t.NotifyEvent(context.Background(), event)
		if err != nil {
			cmd.Logger.Error(err)
		}
	}
}

func main() {
	conf := cmd.Config.NewConfig()
	logg := logger.New(conf.Logger.Level)
	storage := sqlstorage.New()
	eventsNotifier := rabbitmq.NewEventsNotifier()
	errs := make(chan error, 2)
	tm := TasksManager{
		config:         cmd.Config.NewConfig(),
		events:         make(chan entity.Event),
		logger:         logg,
		storage:        storage,
		EventsNotifier: eventsNotifier,
	}
	go func() {
		c := make(chan os.Signal, 2)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		<-c
		errs <- nil
		defer tm.Stop()
	}()
	go tm.EventNotifyWriter()
	tm.EventsNotifyReader()

}
