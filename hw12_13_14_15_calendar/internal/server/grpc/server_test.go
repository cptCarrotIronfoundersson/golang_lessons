package grpc

import (
	"context"
	"fmt"
	"github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/app"
	"github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/logger"
	pbgrpc "github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/server/grpc/pb"
	memorystorage "github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/storage/memory"
	"github.com/magiconair/properties/assert"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
)

func TestCreateEvents(t *testing.T) {
	storage := memorystorage.New()
	logg := logger.New("DEBUG")
	calendar := app.New(logg, storage)

	server := ServerApp{
		UnimplementedCalendarServer: pbgrpc.UnimplementedCalendarServer{},
		app:                         calendar,
	}

	var testEvents = pbgrpc.Events{
		Events: []*pbgrpc.Event{
			{
				UUID:  "efcc9d68-a658-4a36-92b0-0af8ffa94add",
				Title: "test title1",
				Datetime: &timestamppb.Timestamp{
					Seconds: 1661782225,
					Nanos:   0,
				},
				StartDatetime: &timestamppb.Timestamp{
					Seconds: 1661782225,
					Nanos:   0,
				},
				EndDatetime: &timestamppb.Timestamp{
					Seconds: 1661782225,
					Nanos:   0,
				},
				Description:      "Literally speaking the epoch is Unix time 0 (midnight 1/1/1970), but 'epoch' is often used as ",
				UserID:           "8ae89fb0-e5e0-482b-9b4e-36f3298ea4dc",
				RemindTimeBefore: &durationpb.Duration{Seconds: 100, Nanos: 100},
			},
			{
				Title: "test title2",
				Datetime: &timestamppb.Timestamp{
					Seconds: 1661782225,
					Nanos:   0,
				},
				StartDatetime: &timestamppb.Timestamp{
					Seconds: 1661782225,
					Nanos:   0,
				},
				EndDatetime: &timestamppb.Timestamp{
					Seconds: 1661782225,
					Nanos:   0,
				},
				Description:      "Lorem ipsum sit amet",
				UserID:           "8ae89fb0-e5e0-482b-9b4e-36f3298ea4db",
				RemindTimeBefore: &durationpb.Duration{Seconds: 100, Nanos: 100},
			},
		},
	}

	ctx := context.Background()
	for _, event := range testEvents.Events {
		_, err := server.CreateEvent(ctx, &pbgrpc.EventCreate{
			Title:            event.Title,
			Datetime:         event.Datetime,
			StartDatetime:    event.StartDatetime,
			EndDatetime:      event.EndDatetime,
			Description:      event.Description,
			UserID:           event.UserID,
			RemindTimeBefore: event.RemindTimeBefore,
		})
		if err != nil {
			panic(err)
		}
	}

	var date = pbgrpc.Datetime{
		Timestamp: &timestamppb.Timestamp{
			Seconds: 0,
		},
	}

	date.Timestamp.Seconds = 1
	allEvents, _ := server.GetAllEvents(ctx, nil)

	fmt.Printf("%+v\n", allEvents)
	for i := 0; i < len(allEvents.Events); i++ {
		assert.Equal(t, allEvents.Events[i].UserID, testEvents.Events[i].UserID)
		assert.Equal(t, allEvents.Events[i].StartDatetime.Seconds, testEvents.Events[i].StartDatetime.Seconds)
		assert.Equal(t, allEvents.Events[i].StartDatetime.Nanos, testEvents.Events[i].StartDatetime.Nanos)
		assert.Equal(t, allEvents.Events[i].EndDatetime.Seconds, testEvents.Events[i].EndDatetime.Seconds)
		assert.Equal(t, allEvents.Events[i].EndDatetime.Nanos, testEvents.Events[i].EndDatetime.Nanos)
		assert.Equal(t, allEvents.Events[i].Description, testEvents.Events[i].Description)
		assert.Equal(t, allEvents.Events[i].Title, testEvents.Events[i].Title)
	}
	fmt.Printf("%+v\n", &testEvents)
}
