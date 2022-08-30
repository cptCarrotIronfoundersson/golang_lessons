package entity

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	UUID             uuid.UUID     `json:"UUID"`
	Title            string        `json:"Title"`
	Datetime         time.Time     `json:"Datetime"`
	StartDatetime    time.Time     `json:"StartDatetime"`
	EndDatetime      time.Time     `json:"EndDatetime"`
	Description      string        `json:"Description"`
	UserID           uuid.UUID     `json:"UserID"`
	RemindTimeBefore time.Duration `json:"RemindTimeBefore"`
}
