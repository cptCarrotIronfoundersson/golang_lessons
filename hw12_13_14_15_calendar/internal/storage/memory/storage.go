package memorystorage

import (
	"context"
	"sync"
	"time"

	"github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/service/entity"
	"github.com/google/uuid"
)

type Storage struct {
	mu     sync.RWMutex
	events map[uuid.UUID]entity.Event
}

func New() *Storage {
	return &Storage{mu: sync.RWMutex{}}
}

func (s *Storage) Create(ctx context.Context, event entity.Event) error {
	s.mu.Lock()
	s.events[uuid.New()] = event
	defer s.mu.Unlock()
	return nil
}

func (s *Storage) Update(ctx context.Context, event entity.Event, uuid uuid.UUID) error {
	s.mu.Lock()
	s.events[uuid] = event
	defer s.mu.Unlock()
	return nil
}

func (s *Storage) Delete(ctx context.Context, uuid uuid.UUID) error {
	s.mu.Lock()
	delete(s.events, uuid)
	defer s.mu.Unlock()
	return nil
}

func (s *Storage) EventsListDateRange(ctx context.Context, start time.Time, end time.Time) ([]entity.Event, error) {
	EventList := make([]entity.Event, 0)
	s.mu.Lock()
	for _, value := range s.events {
		if value.StartDatetime.After(start) || value.EndDatetime.Before(end) {
			EventList = append(EventList, value)
		}
	}
	defer s.mu.Unlock()
	return EventList, nil
}
