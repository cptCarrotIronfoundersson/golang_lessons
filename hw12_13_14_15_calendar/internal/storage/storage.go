package storage

import (
	"context"
	"time"

	"github.com/cptCarrotIronfoundersson/hw12_13_14_15_calendar/internal/service/entity"
	"github.com/google/uuid"
)

//go:generate mockgen -destination=mocks/storage_mock.go -package=mocks . Storage

type Storage interface {
	Create(ctx context.Context, event entity.Event) error
	Delete(ctx context.Context, uuid uuid.UUID) error
	Update(ctx context.Context, event entity.Event, UUID uuid.UUID) error
	EventsListDateRange(ctx context.Context, startDate time.Time, endDate time.Time) ([]entity.Event, error)
	AllEvents(ctx context.Context) ([]entity.Event, error)
}
