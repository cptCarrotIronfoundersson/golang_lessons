package sqlstorage

import (
	"context"
	"time"

	"github.com/cptCarrotIronfoundersson/hw12_13_14_15_calendar/cmd"
	"github.com/cptCarrotIronfoundersson/hw12_13_14_15_calendar/internal/service/entity"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Storage struct { // TODO
	conn *sqlx.DB
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) Connect(ctx context.Context) error {
	conn, err := sqlx.Connect(`pgx`, cmd.Config.Storage.DSN)
	conn.SetMaxIdleConns(1)
	conn.SetMaxOpenConns(1)
	s.conn = conn
	return err
}

func (s *Storage) Close(ctx context.Context) error {
	err := s.conn.Close()
	return err
}

func (s *Storage) Create(ctx context.Context, event entity.Event) error {
	s.conn.MustExec(
		CreateEvent,
		event.UUID,
		event.Title,
		event.Datetime,
		event.StartDatetime,
		event.EndDatetime,
		event.Description,
		event.UserID,
		event.RemindTimeBefore)
	return nil
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
	eventsList := make([]entity.Event, 0)
	err := s.conn.SelectContext(context.Background(), eventsList, GetEventsByTimeRange, start, end)
	if err != nil {
		return nil, err
	}
	return eventsList, nil
}

func (s *Storage) AllEvents(ctx context.Context) ([]entity.Event, error) {
	eventsList := make([]entity.Event, 0)
	err := s.conn.SelectContext(context.Background(), eventsList, GetAllEvents)
	if err != nil {
		return nil, err
	}
	return eventsList, nil
}
