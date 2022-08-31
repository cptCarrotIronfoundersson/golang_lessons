package grpc

import (
	"context"
	"fmt"
	"github.com/cptCarrotIronfoundersson/hw12_13_14_15_calendar/configs/config"
	"github.com/cptCarrotIronfoundersson/hw12_13_14_15_calendar/internal/app"
	"github.com/cptCarrotIronfoundersson/hw12_13_14_15_calendar/internal/logger"
	pbgrpc "github.com/cptCarrotIronfoundersson/hw12_13_14_15_calendar/internal/server/grpc/pb"
	"github.com/cptCarrotIronfoundersson/hw12_13_14_15_calendar/internal/service/entity"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
	"time"
)

type ServerApp struct {
	pbgrpc.UnimplementedCalendarServer
	app app.Application
}

type Server struct {
	Host     string
	Port     string
	listener net.Listener
	server   *grpc.Server
	app      app.Application
}

func (a ServerApp) CreateEvent(ctx context.Context, event *pbgrpc.EventCreate) (*pbgrpc.Result, error) {
	UserUUID, err := uuid.Parse(event.UserID)
	if err != nil {
		return nil, err
	}
	ev := entity.Event{
		Title:            event.Title,
		Datetime:         time.Unix(event.Datetime.Seconds, int64(event.Datetime.Nanos)),
		StartDatetime:    time.Unix(event.StartDatetime.Seconds, int64(event.StartDatetime.Nanos)),
		EndDatetime:      time.Unix(event.EndDatetime.Seconds, int64(event.EndDatetime.Nanos)),
		Description:      event.Description,
		UserID:           UserUUID,
		RemindTimeBefore: event.RemindTimeBefore.AsDuration(),
	}

	fmt.Println(event.RemindTimeBefore.Nanos)
	err = a.app.CreateEvent(ctx, ev)

	if err != nil {
		return nil, err
	}
	return &pbgrpc.Result{Success: true}, nil
}

func (a ServerApp) UpdateEvent(ctx context.Context, event *pbgrpc.Event) (*pbgrpc.Result, error) {
	eventUUID, err := uuid.Parse(event.UUID)
	if err != nil {
		return nil, err
	}
	UserUUID, err := uuid.Parse(event.UUID)
	if err != nil {
		return nil, err
	}
	ev := entity.Event{
		UUID:             eventUUID,
		Title:            event.Title,
		Datetime:         time.Unix(event.Datetime.Seconds, int64(event.Datetime.Nanos)),
		StartDatetime:    time.Unix(event.StartDatetime.Seconds, int64(event.StartDatetime.Nanos)),
		EndDatetime:      time.Unix(event.EndDatetime.Seconds, int64(event.EndDatetime.Nanos)),
		Description:      event.Description,
		UserID:           UserUUID,
		RemindTimeBefore: time.Duration(event.RemindTimeBefore.Nanos),
	}
	err = a.app.UpdateEvent(ctx, ev)

	if err != nil {
		return nil, err
	}
	return &pbgrpc.Result{Success: true}, nil
}

func (a ServerApp) DeleteEvent(ctx context.Context, event *pbgrpc.Event) (*pbgrpc.Result, error) {
	eventUUID, err := uuid.Parse(event.UUID)
	if err != nil {
		return nil, err
	}
	err = a.app.DeleteEvent(ctx, eventUUID)

	if err != nil {
		return nil, err
	}
	return &pbgrpc.Result{Success: true}, nil
}

func (a ServerApp) GetEventsByDay(ctx context.Context, datetime *pbgrpc.Datetime) (*pbgrpc.Events, error) {
	startDatetime := time.Unix(datetime.Timestamp.Seconds, 0)
	events, err := a.app.GetEventsBetween(ctx, startDatetime, startDatetime.AddDate(0, 0, 1))
	if err != nil {
		return nil, err
	}
	var grpcEvents pbgrpc.Events
	for _, ev := range events {

		grpcEvents.Events = append(grpcEvents.Events, &pbgrpc.Event{
			UUID:  ev.UUID.String(),
			Title: ev.Title,
			Datetime: &timestamppb.Timestamp{
				Seconds: ev.Datetime.Unix(),
				Nanos:   0,
			},
			StartDatetime: &timestamppb.Timestamp{
				Seconds: ev.StartDatetime.Unix(),
				Nanos:   0,
			},
			EndDatetime: &timestamppb.Timestamp{
				Seconds: ev.EndDatetime.Unix(),
				Nanos:   0,
			},
			Description:      ev.Description,
			UserID:           ev.UserID.String(),
			RemindTimeBefore: &durationpb.Duration{Seconds: int64(ev.RemindTimeBefore.Seconds())},
		})
	}
	return &grpcEvents, nil
}

func (a ServerApp) GetEventsByWeek(ctx context.Context, datetime *pbgrpc.Datetime) (*pbgrpc.Events, error) {
	startDatetime := time.Unix(datetime.Timestamp.Seconds, 0)
	events, err := a.app.GetEventsBetween(ctx, startDatetime, startDatetime.AddDate(0, 0, 7))
	if err != nil {
		return nil, err
	}
	var grpcEvents pbgrpc.Events
	for _, ev := range events {

		grpcEvents.Events = append(grpcEvents.Events, &pbgrpc.Event{
			UUID:  ev.UUID.String(),
			Title: ev.Title,
			Datetime: &timestamppb.Timestamp{
				Seconds: ev.Datetime.Unix(),
				Nanos:   0,
			},
			StartDatetime: &timestamppb.Timestamp{
				Seconds: ev.StartDatetime.Unix(),
				Nanos:   0,
			},
			EndDatetime: &timestamppb.Timestamp{
				Seconds: ev.EndDatetime.Unix(),
				Nanos:   0,
			},
			Description:      ev.Description,
			UserID:           ev.UserID.String(),
			RemindTimeBefore: &durationpb.Duration{Seconds: int64(ev.RemindTimeBefore.Seconds())},
		})
	}
	return &grpcEvents, nil
}

func (a ServerApp) GetEventsByMonth(ctx context.Context, datetime *pbgrpc.Datetime) (*pbgrpc.Events, error) {
	startDatetime := time.Unix(datetime.Timestamp.Seconds, 0)
	events, err := a.app.GetEventsBetween(ctx, startDatetime, startDatetime.AddDate(0, 1, 0))
	if err != nil {
		return nil, err
	}
	var grpcEvents pbgrpc.Events
	for _, ev := range events {

		grpcEvents.Events = append(grpcEvents.Events, &pbgrpc.Event{
			UUID:  ev.UUID.String(),
			Title: ev.Title,
			Datetime: &timestamppb.Timestamp{
				Seconds: ev.Datetime.Unix(),
				Nanos:   0,
			},
			StartDatetime: &timestamppb.Timestamp{
				Seconds: ev.StartDatetime.Unix(),
				Nanos:   0,
			},
			EndDatetime: &timestamppb.Timestamp{
				Seconds: ev.EndDatetime.Unix(),
				Nanos:   0,
			},
			Description:      ev.Description,
			UserID:           ev.UserID.String(),
			RemindTimeBefore: &durationpb.Duration{Seconds: int64(ev.RemindTimeBefore.Seconds())},
		})
	}
	return &grpcEvents, nil
}
func (a ServerApp) GetAllEvents(ctx context.Context, Empty *pbgrpc.Empty) (*pbgrpc.Events, error) {
	events, err := a.app.GetAllEvents(ctx)
	if err != nil {
		return nil, err
	}
	var grpcEvents pbgrpc.Events
	for _, ev := range events {

		grpcEvents.Events = append(grpcEvents.Events, &pbgrpc.Event{
			UUID:  ev.UUID.String(),
			Title: ev.Title,
			Datetime: &timestamppb.Timestamp{
				Seconds: ev.Datetime.Unix(),
				Nanos:   0,
			},
			StartDatetime: &timestamppb.Timestamp{
				Seconds: ev.StartDatetime.Unix(),
				Nanos:   0,
			},
			EndDatetime: &timestamppb.Timestamp{
				Seconds: ev.EndDatetime.Unix(),
				Nanos:   0,
			},
			Description: ev.Description,
			UserID:      ev.UserID.String(),
			RemindTimeBefore: &durationpb.Duration{Seconds: int64(ev.RemindTimeBefore.Seconds()),
				Nanos: int32(ev.RemindTimeBefore.Nanoseconds())},
		})
	}

	return &grpcEvents, nil
}

func NewServer(logger *logger.Logger, config *config.Config, app app.Application) *Server {
	return &Server{
		Host:   config.GrpcServer.Host,
		Port:   config.GrpcServer.Port,
		server: grpc.NewServer(),
		app:    app,
	}
}

func (s *Server) Start(ctx context.Context) error {
	lsn, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.Host, s.Port))
	s.listener = lsn
	if err != nil {
		log.Fatal("listener error ", err)
	}
	server := grpc.NewServer()
	serverApp := new(ServerApp)
	serverApp.app = s.app
	pbgrpc.RegisterCalendarServer(server, serverApp)
	if err := server.Serve(lsn); err != nil {
		return err
	}
	<-ctx.Done()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	err := s.listener.Close()

	if err != nil {
		log.Fatal(err)
	}
	s.server.GracefulStop()
	log.Println("Server successfully stopped")
	return nil
}
