package sqlstorage

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cptCarrotIronfoundersson/hw12_13_14_15_calendar/cmd"
	"github.com/cptCarrotIronfoundersson/hw12_13_14_15_calendar/internal/service/entity"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	conn *sqlx.DB
}

func New() *Storage {
	newSt := &Storage{}
	if err := newSt.Connect(context.Background()); err != nil {
		log.Fatal(err)
	}
	return newSt
}

func (s *Storage) Connect(ctx context.Context) error {
	conn, err := sqlx.Connect(`postgres`, cmd.Config.Storage.DSN)
	s.conn = conn
	return err
}

func (s *Storage) Close(ctx context.Context) error {
	err := s.conn.Close()
	return err
}

func (s *Storage) Create(ctx context.Context, event entity.Event) error {
	_, err := s.conn.ExecContext(ctx,
		CreateEvent,
		uuid.New().String(),
		event.Title,
		event.Datetime,
		event.StartDatetime,
		event.EndDatetime,
		event.Description,
		event.UserID,
		event.RemindTimeBefore,
	)
	if err != nil {
		panic(err)
	}
	return err
}

func (s *Storage) Delete(ctx context.Context, uuid uuid.UUID) error {
	_, err := s.conn.Exec(DeleteEvent, uuid)
	return err
}

func (s *Storage) Update(ctx context.Context, event entity.Event, uuid uuid.UUID) error {
	_, err := s.conn.Exec(
		UpdateEvent,
		uuid,
		event.UUID,
		event.Title,
		event.Datetime,
		event.StartDatetime,
		event.EndDatetime,
		event.Description,
		event.UserID,
		event.RemindTimeBefore)
	return err
}

func (s *Storage) EventsListDateRange(ctx context.Context, start time.Time, end time.Time) ([]entity.Event, error) {
	var eventsList []entity.Event
	err := s.conn.SelectContext(context.Background(), &eventsList, GetEventsByTimeRange, start, end)
	if err != nil {
		return nil, err
	}
	return eventsList, nil
}

func (s *Storage) AllEvents(ctx context.Context) ([]entity.Event, error) {
	var eventsList []entity.Event
	err := s.conn.SelectContext(context.Background(), &eventsList, GetAllEvents)
	fmt.Println(eventsList, "sdasd")
	if err != nil {
		return nil, err
	}
	return eventsList, nil
}
