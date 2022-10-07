package app

import (
	"context"
	"time"

	"github.com/cptCarrotIronfoundersson/hw12_13_14_15_calendar/internal/logger"
	"github.com/cptCarrotIronfoundersson/hw12_13_14_15_calendar/internal/service/entity"
	"github.com/cptCarrotIronfoundersson/hw12_13_14_15_calendar/internal/storage"
	"github.com/google/uuid"
)

type Application interface {
	CreateEvent(ctx context.Context, event entity.Event) (entity.Event, error)
	UpdateEvent(ctx context.Context, event entity.Event) error
	DeleteEvent(ctx context.Context, eventUUID uuid.UUID) error
	GetEventsBetween(ctx context.Context, startDate time.Time, endDate time.Time) ([]entity.Event, error)
	GetAllEvents(ctx context.Context) ([]entity.Event, error)
}

type App struct {
	Storage storage.Storage
	Logger  *logger.Logger
}

func New(logger *logger.Logger, storage storage.Storage) *App {
	return &App{Storage: storage, Logger: logger}
}

func (a *App) CreateEvent(ctx context.Context, event entity.Event) (entity.Event, error) {
	event, err := a.Storage.Create(ctx, event)
	if err != nil {
		return entity.Event{}, err
	}
	return event, nil
}

func (a *App) UpdateEvent(ctx context.Context, event entity.Event) error {
	err := a.Storage.Update(ctx, event, event.UUID)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) DeleteEvent(ctx context.Context, eventUUID uuid.UUID) error {
	err := a.Storage.Delete(ctx, eventUUID)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) GetEventsBetween(ctx context.Context, startDate time.Time, endDate time.Time) ([]entity.Event, error) {
	events, err := a.Storage.EventsListDateRange(ctx, startDate, endDate)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (a *App) GetAllEvents(ctx context.Context) ([]entity.Event, error) {
	events, err := a.Storage.AllEvents(ctx)
	if err != nil {
		return nil, err
	}
	return events, nil
}
