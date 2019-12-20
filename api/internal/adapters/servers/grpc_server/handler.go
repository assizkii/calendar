package grpc_server

import (
	"context"
	"errors"
	"fmt"
	"github.com/assizkii/calendar/api/internal/domain/entities"
	"github.com/assizkii/calendar/api/internal/domain/interfaces"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type EventServiceServer struct {
	storage interfaces.EventStorage
}


func (s *EventServiceServer) CreateEvent(ctx context.Context, req *entities.EventCreateRequest) (*entities.EventCreateResponse, error) {

	event := req.GetEvent()

	eventData := entities.Event{
		Id:          uuid.New().String(),
		Title:       event.GetTitle(),
		Description: event.GetDescription(),
		Start:       event.GetStart(),
		OwnerId: event.GetOwnerId(),
		EndTime: event.GetEndTime(),
	}

	id, err := s.storage.Add(eventData)
	if err != nil {
		// return internal gRPC error
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	return &entities.EventCreateResponse{Id: id}, nil
}

func (s *EventServiceServer) UpdateEvent(ctx context.Context, req *entities.EventUpdateRequest) (*entities.EventUpdateResponse, error) {

	event := req.GetEvent()

	eventData := entities.Event{
		Title:       event.GetTitle(),
		Description: event.GetDescription(),
		Start:       event.GetStart(),
		OwnerId: event.GetOwnerId(),
		EndTime: event.GetEndTime(),
	}

	err := s.storage.Update(event.GetId(), eventData)
	if err != nil {
		// return internal gRPC error
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal update error: %v", err),
		)
	}

	return &entities.EventUpdateResponse{Id: event.GetId()}, nil
}

func (s *EventServiceServer) DeleteEvent(ctx context.Context, req *entities.EventDeleteRequest) (*entities.EventDeleteResponse, error) {

	eventId := req.GetId()

	err := s.storage.Delete(eventId)

	if err != nil {
		// return internal gRPC error
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	return &entities.EventDeleteResponse{Success: true}, nil

}

func (s *EventServiceServer) EventList(req *entities.EventListRequest, stream entities.EventService_EventListServer) error {

	t := time.Now()
	timeStart := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)

	timeEnd := timeStart
	period := req.Period

	switch period.String() {
	case "DAY":
		timeEnd = timeStart.AddDate(0, 0, 1).Add(time.Nanosecond * -1)
	case "WEEK":
		timeEnd = timeStart.AddDate(0, 0, 7).Add(time.Nanosecond * -1)
	case "MONTH":
		timeEnd = timeStart.AddDate(0, 1, 0).Add(time.Nanosecond * -1)
	default:
		return status.Errorf(codes.Internal, fmt.Sprintf("Steam error: %v", errors.New("you must set correct period")))
	}



	eventList, err := s.storage.FilterByDate(timeStart, timeEnd)
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Database error: %v", err))
	}
	for _, event := range eventList {

		err := stream.SendMsg(&entities.EventListResponse{
			Event: &event,
		})

		if err != nil {
			return status.Errorf(codes.Internal, fmt.Sprintf("Steam error: %v", err))
		}
	}

	return nil
}
