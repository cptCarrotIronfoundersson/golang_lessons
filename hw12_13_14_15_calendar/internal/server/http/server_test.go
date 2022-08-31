package internalhttp

import (
	"encoding/json"
	"github.com/cptCarrotIronfoundersson/hw12_13_14_15_calendar/internal/app"
	"github.com/cptCarrotIronfoundersson/hw12_13_14_15_calendar/internal/logger"
	"github.com/cptCarrotIronfoundersson/hw12_13_14_15_calendar/internal/service/entity"
	memorystorage "github.com/cptCarrotIronfoundersson/hw12_13_14_15_calendar/internal/storage/memory"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestCreateEvents(t *testing.T) {
	storage := memorystorage.New()
	logg := logger.New("DEBUG")
	calendar := app.New(logg, storage)
	var strTestUUids = []string{"d35e2669-6d1f-408d-a19f-38dab00a8772", "d35e2669-6d1f-408d-a19f-38dab00a8771"}
	var testUUIDS []uuid.UUID
	for _, val := range strTestUUids {
		testUuid, _ := uuid.Parse(val)
		testUUIDS = append(testUUIDS, testUuid)
	}

	var testEvents = []entity.Event{{
		Title:            "lorem",
		Datetime:         time.Now(),
		StartDatetime:    time.Date(2022, 9, 8, 11, 45, 10, 1000, time.UTC),
		EndDatetime:      time.Date(2022, 9, 9, 11, 45, 10, 1000, time.UTC),
		Description:      "lorem ipsum sit amet",
		UserID:           testUUIDS[0],
		RemindTimeBefore: 100 * time.Second,
	}, {
		Title:            "от винтааа",
		Datetime:         time.Now(),
		StartDatetime:    time.Date(2022, 5, 8, 11, 45, 10, 1000, time.UTC),
		EndDatetime:      time.Date(2022, 6, 9, 11, 45, 10, 1000, time.UTC),
		Description:      "Потому что в самолете жизнь зависит от.... от... ",
		UserID:           testUUIDS[1],
		RemindTimeBefore: 500 * time.Second,
	},
	}

	var serverAppTest = ServerApp{
		app:    calendar,
		logger: logg,
	}

	var createEventRequests []*http.Request

	for _, val := range testEvents {
		eventJson, _ := json.Marshal(val)
		createEventRequest := &http.Request{
			Method:        http.MethodPost,
			Proto:         "",
			ProtoMajor:    0,
			ProtoMinor:    0,
			Body:          ioutil.NopCloser(strings.NewReader(string(eventJson))),
			ContentLength: 0,
			Close:         false,
			Host:          "test",
			RemoteAddr:    "test",
			RequestURI:    "test",
		}
		createEventRequests = append(createEventRequests, createEventRequest)
	}

	for _, req := range createEventRequests {
		createEventRw := httptest.NewRecorder()
		serverAppTest.createEvent(createEventRw, req)
		assert.Equal(t, createEventRw.Result().StatusCode, http.StatusOK)
	}
	getAllEventsRw := httptest.NewRecorder()
	serverAppTest.getAllEvents(getAllEventsRw, &http.Request{
		Method:        http.MethodGet,
		Proto:         "",
		ProtoMajor:    0,
		ProtoMinor:    0,
		Body:          ioutil.NopCloser(strings.NewReader(``)),
		ContentLength: 0,
		Close:         false,
		Host:          "test",
		RemoteAddr:    "test",
		RequestURI:    "test",
	})
	var testEventsExp []entity.Event
	json.Unmarshal([]byte(getAllEventsRw.Body.String()), &testEventsExp)
	for i, _ := range testEventsExp {
		assert.Equal(t, testEventsExp[i].UserID, testEvents[i].UserID)
		assert.Equal(t, testEventsExp[i].StartDatetime, testEvents[i].StartDatetime)
		assert.Equal(t, testEventsExp[i].StartDatetime, testEvents[i].StartDatetime)
		assert.Equal(t, testEventsExp[i].EndDatetime, testEvents[i].EndDatetime)
		assert.Equal(t, testEventsExp[i].EndDatetime, testEvents[i].EndDatetime)
		assert.Equal(t, testEventsExp[i].Description, testEvents[i].Description)
		assert.Equal(t, testEventsExp[i].Title, testEvents[i].Title)
	}
}
