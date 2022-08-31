package internalhttp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/cptCarrotIronfoundersson/hw12_13_14_15_calendar/configs/config"
	"github.com/cptCarrotIronfoundersson/hw12_13_14_15_calendar/internal/app"
	"github.com/cptCarrotIronfoundersson/hw12_13_14_15_calendar/internal/logger"
	"github.com/cptCarrotIronfoundersson/hw12_13_14_15_calendar/internal/service/entity"
	"github.com/google/uuid"
)

type Server struct {
	Host string
	Port string
}

type Logger interface {
	Info(args ...interface{})
	Error(args ...interface{})
}

type ServerApp struct {
	app    app.Application
	logger Logger
}

func (a ServerApp) createEvent(rw http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	decoder := json.NewDecoder(req.Body)
	var event entity.Event
	err := decoder.Decode(&event)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		a.logger.Error("server.createEvent " + err.Error())
		return
	}
	err = a.app.CreateEvent(req.Context(), event)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		a.logger.Error("server.createEvent " + err.Error())
		return
	}
	rw.WriteHeader(http.StatusOK)
}

func (a ServerApp) UpdateEvent(rw http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	decoder := json.NewDecoder(req.Body)
	var event entity.Event
	err := decoder.Decode(&event)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		a.logger.Error("server.UpdateEvent " + err.Error())
		return
	}
	if event.UUID == uuid.Nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("UUID was not specified"))
		return
	}
	err = a.app.UpdateEvent(context.Background(), event)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		a.logger.Error("server.UpdateEvent " + err.Error())
		return
	}
}

func (a ServerApp) DeleteEvent(rw http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	decoder := json.NewDecoder(req.Body)
	var event entity.Event
	err := decoder.Decode(&event)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		a.logger.Error("server.DeleteEvent " + err.Error())
		return
	}
	err = a.app.DeleteEvent(context.Background(), event.UUID)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		a.logger.Error("server.DeleteEvent " + err.Error())
		return
	}
}

func (a ServerApp) getEventsByDay(rw http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	events, err := a.app.GetEventsBetween(context.Background(), time.Now(), time.Now().AddDate(0, 0, 1))
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		a.logger.Error("server.getEventsByDay " + err.Error())
		return
	}
	eventsJSON, err := json.MarshalIndent(events, "", "\t")
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		a.logger.Error("server.getEventsByDay " + err.Error())
		return
	}
	rw.Write(eventsJSON)
}

func (a ServerApp) getEventsByWeek(rw http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	events, err := a.app.GetEventsBetween(context.Background(), time.Now(), time.Now().AddDate(0, 0, 7))
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		a.logger.Error("server.getEventsByWeek " + err.Error())
		return
	}
	eventsJSON, err := json.MarshalIndent(events, "", "\t")
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		a.logger.Error("server.getEventsByWeek " + err.Error())
		return
	}
	rw.Write(eventsJSON)
}

func (a ServerApp) getEventsByMonth(rw http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	events, err := a.app.GetEventsBetween(context.Background(), time.Now(), time.Now().AddDate(0, 1, 0))
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		a.logger.Error("server.getEventsByMounth " + err.Error())
		return
	}
	eventsJSON, err := json.MarshalIndent(events, "", "\t")
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		a.logger.Error("server.getEventsByMounth " + err.Error())
		return
	}
	rw.Write(eventsJSON)
}

func (a ServerApp) getAllEvents(rw http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		events, err := a.app.GetAllEvents(context.Background())
		if err != nil {
			fmt.Println(err)
		}
		eventsJSON, err := json.MarshalIndent(events, "", "\t")
		if err != nil {
			a.logger.Error(err)
		}
		rw.Write(eventsJSON)
	} else {
		rw.WriteHeader(http.StatusBadRequest)
	}
}

func NewServer(logger *logger.Logger, config *config.Config, app app.Application) *Server {
	logger.Info(fmt.Sprintf("Server startded:  Host %v, Port %v", config.HTTPServer.Host, config.GRPCServer.Port))
	application := ServerApp{
		app:    app,
		logger: logger,
	}
	eventCreate := middlewareChainApply(logger, http.HandlerFunc(application.createEvent),
		[]middleware{EnsureAppJSONMiddleware, loggingMiddleware})
	getAllEvents := middlewareChainApply(logger, http.HandlerFunc(application.getAllEvents),
		[]middleware{EnsureAppJSONMiddleware, loggingMiddleware})
	eventUpdate := middlewareChainApply(logger, http.HandlerFunc(application.UpdateEvent),
		[]middleware{EnsureAppJSONMiddleware, loggingMiddleware})
	eventDelete := middlewareChainApply(logger, http.HandlerFunc(application.DeleteEvent),
		[]middleware{EnsureAppJSONMiddleware, loggingMiddleware})
	getEventsByDay := middlewareChainApply(logger, http.HandlerFunc(application.getEventsByDay),
		[]middleware{EnsureAppJSONMiddleware, loggingMiddleware})
	getEventsByWeek := middlewareChainApply(logger, http.HandlerFunc(application.getEventsByWeek),
		[]middleware{loggingMiddleware})
	getEventsByMonth := middlewareChainApply(logger, http.HandlerFunc(application.getEventsByMonth),
		[]middleware{EnsureAppJSONMiddleware, loggingMiddleware})

	http.Handle("/event/create", eventCreate)
	http.Handle("/event/get_all", getAllEvents)
	http.Handle("/event/update", eventUpdate)
	http.Handle("/event/delete", eventDelete)
	http.Handle("/event/get_events_by_day", getEventsByDay)
	http.Handle("/event/get_events_by_week", getEventsByWeek)
	http.Handle("/event/get_events_by_month", getEventsByMonth)
	return &Server{
		Host: config.GRPCServer.Host,
		Port: config.GRPCServer.Port,
	}
}

func (s *Server) Start(ctx context.Context) error {
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", s.Host, s.Port), nil)
	fmt.Println(err)
	<-ctx.Done()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	os.Exit(1)
	return nil
}
