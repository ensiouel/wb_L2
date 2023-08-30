package service

import (
	"calendar/internal/dto"
	"calendar/internal/model"
	"calendar/internal/storage"
	"context"
	"time"
)

type EventService interface {
	Create(ctx context.Context, request dto.CreateEvent) (model.Event, error)
	Update(ctx context.Context, request dto.UpdateEvent) (model.Event, error)
	Delete(ctx context.Context, eventID int) error
	EventsForDay(ctx context.Context, day time.Time) ([]model.Event, error)
	EventsForWeek(ctx context.Context, week time.Time) ([]model.Event, error)
	EventsForMonth(ctx context.Context, month time.Time) ([]model.Event, error)
}

type EventServiceImpl struct {
	storage storage.EventStorage
}

func NewEventService(storage storage.EventStorage) *EventServiceImpl {
	return &EventServiceImpl{storage: storage}
}

func (service *EventServiceImpl) Create(ctx context.Context, request dto.CreateEvent) (model.Event, error) {
	event := model.Event{
		UserID:      request.UserID,
		Title:       request.Title,
		Description: request.Description,
		Date:        request.Date,
	}

	id, err := service.storage.Create(ctx, event)
	if err != nil {
		return model.Event{}, err
	}

	event.ID = id

	return event, nil
}

func (service *EventServiceImpl) Update(ctx context.Context, request dto.UpdateEvent) (model.Event, error) {
	event, err := service.storage.GetByID(ctx, request.EventID)
	if err != nil {
		return model.Event{}, err
	}

	if request.Title != "" {
		event.Title = request.Title
	}

	if request.Description != "" {
		event.Description = request.Description
	}

	if !request.Date.IsZero() {
		event.Date = request.Date
	}

	err = service.storage.Update(ctx, event)
	if err != nil {
		return model.Event{}, err
	}

	return event, nil
}

func (service *EventServiceImpl) Delete(ctx context.Context, eventID int) error {
	return service.storage.Delete(ctx, eventID)
}

func (service *EventServiceImpl) EventsForDay(ctx context.Context, day time.Time) ([]model.Event, error) {
	return service.storage.EventsForDay(ctx, day)
}

func (service *EventServiceImpl) EventsForWeek(ctx context.Context, week time.Time) ([]model.Event, error) {
	return service.storage.EventsForWeek(ctx, week)
}

func (service *EventServiceImpl) EventsForMonth(ctx context.Context, month time.Time) ([]model.Event, error) {
	return service.storage.EventsForMonth(ctx, month)
}
