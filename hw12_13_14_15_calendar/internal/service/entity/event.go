package entity

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	UUID             uuid.UUID
	Title            string
	Datetime         time.Time
	StartDatetime    time.Time
	EndDatetime      time.Time
	Description      string
	UserID           uuid.UUID
	RemindTimeBefore time.Duration
}
