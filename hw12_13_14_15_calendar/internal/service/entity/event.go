package entity

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	UUID             uuid.UUID     `json:"uuid" db:"uuid"`
	Title            string        `json:"title" db:"title"`
	Datetime         time.Time     `json:"datetime" db:"datetime"`
	StartDatetime    time.Time     `json:"startDatetime" db:"start_datetime"`
	EndDatetime      time.Time     `json:"endDatetime" db:"end_datetime"`
	Description      string        `json:"description" db:"description"`
	UserID           uuid.UUID     `json:"userId" db:"userid"`
	RemindTimeBefore time.Duration `json:"remindTimeBefore" db:"remind_time_before"`
}
