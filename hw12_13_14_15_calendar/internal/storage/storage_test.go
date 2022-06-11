package storage

// import (
//	"context"
//	"github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/service/entity"
//	"github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/storage/mocks"
//	"github.com/golang/mock/gomock"
//	"github.com/google/uuid"
//	"testing"
//	"time"
//)
//
// func TestSqlStorage(t *testing.T) {
//	events := []entity.Event{
//		{
//			UUID:             uuid.MustParse("8ba2478f-02c4-40b0-b7f5-ae665cbeaa39"),
//			Title:            "Разбуди меня пидрила пушистая",
//			Datetime:         time.Date(2022, 5, 1, 12, 30, 11, 0, time.UTC),
//			StartDatetime:    time.Date(2022, 6, 11, 12, 30, 11, 0, time.UTC),
//			EndDatetime:      time.Date(2022, 6, 12, 12, 30, 11, 0, time.UTC),
//			Description:      "",
//			UserId:           uuid.MustParse("7bc3a1ef-504f-4de1-a472-2a9bed22712c"),
//			RemindTimeBefore: time.Hour,
//		},
//		{
//			UUID:             uuid.MustParse("69b5d204-775b-4e13-8dae-35878826a2ca"),
//			Title:            "Потребовать еды у кожанного",
//			Datetime:         time.Date(2022, 5, 1, 12, 30, 11, 0, time.UTC),
//			StartDatetime:    time.Date(2022, 6, 11, 12, 30, 11, 0, time.UTC),
//			EndDatetime:      time.Date(2022, 6, 12, 12, 30, 11, 0, time.UTC),
//			Description:      "",
//			UserId:           uuid.MustParse("acedd8ff-2c19-409e-b75f-5e14b5552791"),
//			RemindTimeBefore: time.Hour,
//		},
//	}
//
//	ctx := context.Background()
//
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	type StorageCreateMock struct {
//		err error
//	}
//	type StorageDeleteMock struct {
//		err error
//	}
//	type StorageUpdateMock struct {
//		err error
//	}
//	type StorageEventsListDateRangeMock struct {
//		result []entity.Event
//		err    error
//	}
//	type dbRetailerRepositoryMocks struct {
//		StorageCreateMock              *StorageCreateMock
//		getOffersActiveMock            *StorageDeleteMock
//		getOffersStocksInStoreMock     *StorageUpdateMock
//		StorageEventsListDateRangeMock *StorageEventsListDateRangeMock
//	}
//	var tests = []struct {
//		name     string
//		repoMock dbRetailerRepositoryMocks
//		want     *entity.Event
//		wantErr  *string
//	}{{
//		name: "OK",
//		repoMock: dbRetailerRepositoryMocks{
//			StorageCreateMock: &StorageCreateMock{
//				err: nil,
//			},
//			getOffersActiveMock: &StorageDeleteMock{
//				err: nil,
//			},
//			getOffersStocksInStoreMock: &StorageUpdateMock{
//				err: nil,
//			},
//			StorageEventsListDateRangeMock: &StorageEventsListDateRangeMock{
//				result: events,
//				err:    nil,
//			},
//		},
//	}}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			storageMock := mocks.NewMockStorage(ctrl)
//			if tt.repoMock.StorageCreateMock != nil {
//				storageMock.
//					EXPECT().
//					Create(ctx, events[0]).
//					Return(tt.repoMock.StorageCreateMock.err)
//			}
//		})
//	}
// }
