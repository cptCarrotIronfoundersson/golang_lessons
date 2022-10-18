// build integration

package integration_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/cptCarrotIronfoundersson/hw12_13_14_15_calendar/cmd"
	"github.com/cptCarrotIronfoundersson/hw12_13_14_15_calendar/configs/config"
	"github.com/cptCarrotIronfoundersson/hw12_13_14_15_calendar/internal/service/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func init() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

var (
	UserUUID1  = uuid.MustParse("d35e2669-6d1f-408d-a19f-38dab00a8772")
	UserUUID2  = uuid.MustParse("d35e2669-6d1f-408d-a19f-38dab00a8771")
	testEvents = []entity.Event{
		{
			Title:            "lorem",
			Datetime:         time.Now(),
			StartDatetime:    time.Date(2022, 9, 8, 11, 45, 10, 1000, time.UTC),
			EndDatetime:      time.Date(2022, 9, 9, 11, 45, 10, 1000, time.UTC),
			Description:      "lorem ipsum sit amet",
			UserID:           UserUUID1,
			RemindTimeBefore: 100 * time.Second,
		},
		{
			Title:            "от винтааа",
			Datetime:         time.Now(),
			StartDatetime:    time.Date(2022, 5, 8, 11, 45, 10, 1000, time.UTC),
			EndDatetime:      time.Date(2022, 6, 9, 11, 45, 10, 1000, time.UTC),
			Description:      "Потому что в самолете жизнь зависит от.... от... ",
			UserID:           UserUUID2,
			RemindTimeBefore: 500 * time.Second,
		},
		{
			Title:            "lorem",
			Datetime:         time.Now(),
			StartDatetime:    time.Now().Add(-time.Hour),
			EndDatetime:      time.Now().Add(+time.Hour),
			Description:      "lorem ipsum sit amet",
			UserID:           UserUUID1,
			RemindTimeBefore: 100 * time.Second,
		},
		{
			Title:            "от винтааа",
			Datetime:         time.Now(),
			StartDatetime:    time.Now().Add(-time.Hour),
			EndDatetime:      time.Now().Add(+time.Hour * 2),
			Description:      "Потому что в самолете жизнь зависит от.... от... ",
			UserID:           UserUUID2,
			RemindTimeBefore: 500 * time.Second,
		},
		{
			Title:            "lorem",
			Datetime:         time.Now(),
			StartDatetime:    time.Now().Add(time.Hour * 24 * 7),
			EndDatetime:      time.Now().Add(+time.Hour * 24 * 8),
			Description:      "lorem ipsum sit amet",
			UserID:           UserUUID1,
			RemindTimeBefore: 100 * time.Second,
		},
		{
			Title:            "от винтааа",
			Datetime:         time.Now(),
			StartDatetime:    time.Now().Add(time.Hour * 24 * 7 * 27),
			EndDatetime:      time.Now().Add(+time.Hour * 24 * 7 * 28),
			Description:      "Потому что в самолете жизнь зависит от.... от... ",
			UserID:           UserUUID2,
			RemindTimeBefore: 500 * time.Second,
		},
	}
)

type BaseCalendarSuite struct {
	suite.Suite
	ctx        context.Context
	host       string
	httpClient *http.Client
	conf       *config.Config
}

func (h *BaseCalendarSuite) SetupTest() {
	HOST := "127.0.0.1"
	PORT := "8888"
	if h.conf.HTTPServer.Host != "" {
		HOST = h.conf.HTTPServer.Host
	}

	if h.conf.HTTPServer.Port != "" {
		PORT = h.conf.HTTPServer.Port
	}
	h.host = fmt.Sprintf("%s:%s", HOST, PORT)
	h.httpClient = http.DefaultClient
	h.ctx = context.Background()
}

type CalendarSuite struct {
	BaseCalendarSuite
}

func (h *CalendarSuite) SetupSuite() {
	HOST := os.Getenv("http_server.Host")
	if HOST == "" {
		HOST = "127.0.0.1"
	}
	PORT := os.Getenv("http_server.Host")
	if PORT == "" {
		PORT = "2021"
	}

	h.host = fmt.Sprintf("%s:%s", HOST, PORT)
	h.httpClient = http.DefaultClient
	h.ctx = context.Background()
}

func (h *CalendarSuite) createEvents(events []entity.Event) ([]entity.Event, error) {
	CreateEventURL := url.URL{
		Scheme:      "HTTP",
		Opaque:      "",
		User:        nil,
		Host:        h.host,
		Path:        "/event/create",
		RawPath:     "",
		ForceQuery:  false,
		RawQuery:    "",
		Fragment:    "",
		RawFragment: "",
	}
	createdEvents := make([]entity.Event, len(events))
	for i, event := range events {
		eventJSON, _ := json.Marshal(event)
		req, err := http.NewRequestWithContext(h.ctx, http.MethodPost, CreateEventURL.String(), bytes.NewReader(eventJSON))
		req.Header.Add("Content-Type", "application/json")
		if err != nil {
			return createdEvents, err
		}
		resp, err := h.httpClient.Do(req)
		if err != nil {
			return createdEvents, err
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return createdEvents, err
		}
		err = resp.Body.Close()
		if err != nil {
			return createdEvents, err
		}
		if err != nil {
			return createdEvents, err
		}
		var eventResp entity.Event
		err = json.Unmarshal(body, &eventResp)
		if err != nil {
			return createdEvents, err
		}

		if resp.StatusCode != http.StatusOK {
			return createdEvents, fmt.Errorf("BAD StatusCode: %b", resp.StatusCode)
		}
		createdEvents[i] = eventResp
	}
	return createdEvents, nil
}

func (h *CalendarSuite) deleteEvents(events []entity.Event) error {
	DeleteEventURL := url.URL{
		Scheme:      "HTTP",
		Opaque:      "",
		User:        nil,
		Host:        h.host,
		Path:        "/event/delete",
		RawPath:     "",
		ForceQuery:  false,
		RawQuery:    "",
		Fragment:    "",
		RawFragment: "",
	}
	for _, event := range events {
		eventJSON, _ := json.Marshal(event)
		req, err := http.NewRequestWithContext(h.ctx, http.MethodPost, DeleteEventURL.String(), bytes.NewReader(eventJSON))
		req.Header.Add("Content-Type", "application/json")
		if err != nil {
			return err
		}
		resp, err := h.httpClient.Do(req)
		if err != nil {
			return err
		}
		if err != nil {
			return err
		}
		resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("BAD StatusCode: %b", resp.StatusCode)
		}
	}
	return nil
}

func (h *CalendarSuite) TestCreateEvent() {
	events, err := h.createEvents(testEvents)
	require.NoError(h.T(), err)
	err = h.deleteEvents(events)
	require.NoError(h.T(), err)
}

func (h *CalendarSuite) TestGetEventsByDay() {
	GetEventsURL := url.URL{
		Scheme:      "HTTP",
		Opaque:      "",
		User:        nil,
		Host:        h.host,
		Path:        "/event/get_events_by_day",
		RawPath:     "",
		ForceQuery:  false,
		RawQuery:    "",
		Fragment:    "",
		RawFragment: "",
	}
	createdEvents, err := h.createEvents(testEvents)
	require.NoError(h.T(), err)
	req, err := http.NewRequestWithContext(h.ctx, http.MethodGet, GetEventsURL.String(), nil)
	require.NoError(h.T(), err)
	req.Header.Add("Content-Type", "application/json")
	resp, err := h.httpClient.Do(req)
	require.NoError(h.T(), err)
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	var events []entity.Event
	require.NoError(h.T(), err)
	err = json.Unmarshal(body, &events)
	require.NoError(h.T(), err)
	err = h.deleteEvents(createdEvents)
	require.NoError(h.T(), err)
}

func (h *CalendarSuite) TestGetEventsByWeek() {
	GetEventsURL := url.URL{
		Scheme:      "HTTP",
		Opaque:      "",
		User:        nil,
		Host:        h.host,
		Path:        "/event/get_events_by_week",
		RawPath:     "",
		ForceQuery:  false,
		RawQuery:    "",
		Fragment:    "",
		RawFragment: "",
	}
	createdEvents, err := h.createEvents(testEvents)
	require.NoError(h.T(), err)
	req, err := http.NewRequestWithContext(h.ctx, http.MethodGet, GetEventsURL.String(), nil)
	require.NoError(h.T(), err)
	req.Header.Add("Content-Type", "application/json")
	resp, err := h.httpClient.Do(req)
	require.NoError(h.T(), err)
	body, err := io.ReadAll(resp.Body)
	require.NoError(h.T(), err)
	resp.Body.Close()
	var events []entity.Event
	require.NoError(h.T(), err)
	err = json.Unmarshal(body, &events)
	require.NoError(h.T(), err)
	err = h.deleteEvents(createdEvents)
	require.NoError(h.T(), err)
}

func (h *CalendarSuite) TestGetEventsByMonth() {
	GetEventsURL := url.URL{
		Scheme:      "HTTP",
		Opaque:      "",
		User:        nil,
		Host:        h.host,
		Path:        "/event/get_events_by_month",
		RawPath:     "",
		ForceQuery:  false,
		RawQuery:    "",
		Fragment:    "",
		RawFragment: "",
	}
	createdEvents, err := h.createEvents(testEvents)
	require.NoError(h.T(), err)
	req, err := http.NewRequestWithContext(h.ctx, http.MethodGet, GetEventsURL.String(), nil)
	require.NoError(h.T(), err)
	req.Header.Add("Content-Type", "application/json")
	resp, err := h.httpClient.Do(req)
	require.NoError(h.T(), err)
	body, err := io.ReadAll(resp.Body)
	require.NoError(h.T(), err)
	resp.Body.Close()
	var events []entity.Event
	require.NoError(h.T(), err)
	err = json.Unmarshal(body, &events)
	require.NoError(h.T(), err)
	err = h.deleteEvents(createdEvents)
	require.NoError(h.T(), err)
}

func TestHttpSuite(t *testing.T) {
	conf := cmd.Config.NewConfig()
	createSuite := new(CalendarSuite)
	createSuite.conf = conf
	// conf.HTTPServer.Host = ""
	// conf.HTTPServer.Port = ""
	createSuite.ctx = context.Background()
	suite.Run(t, createSuite)
}
