package entity

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	UUID             uuid.UUID     `json:"uuid"`
	Title            string        `json:"title"`
	Datetime         time.Time     `json:"datetime"`
	StartDatetime    time.Time     `json:"startDatetime"`
	EndDatetime      time.Time     `json:"endDatetime"`
	Description      string        `json:"description"`
	UserID           uuid.UUID     `json:"userId"`
	RemindTimeBefore time.Duration `json:"remindTimeBefore"`
}
