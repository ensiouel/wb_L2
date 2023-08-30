package storage

import (
	"calendar/internal/model"
	"context"
	"errors"
	"github.com/ensiouel/apperror"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type EventStorage interface {
	Create(ctx context.Context, event model.Event) (int, error)
	GetByID(ctx context.Context, eventID int) (model.Event, error)
	Update(ctx context.Context, event model.Event) error
	Delete(ctx context.Context, eventID int) error
	EventsForDay(ctx context.Context, day time.Time) ([]model.Event, error)
	EventsForWeek(ctx context.Context, week time.Time) ([]model.Event, error)
	EventsForMonth(ctx context.Context, month time.Time) ([]model.Event, error)
}

type EventStorageImpl struct {
	db *pgxpool.Pool
}

func NewEventStorage(pool *pgxpool.Pool) *EventStorageImpl {
	return &EventStorageImpl{db: pool}
}

func (storage *EventStorageImpl) Create(ctx context.Context, event model.Event) (int, error) {
	var id int
	err := storage.db.QueryRow(ctx, `
INSERT INTO event (user_id, title, description, date)
VALUES ($1, $2, $3, $4)
RETURNING id
`, event.UserID, event.Title, event.Description, event.Date).Scan(&id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, apperror.NotFound.WithError(err)
		}

		return 0, apperror.Internal.WithError(err)
	}

	return id, nil
}

func (storage *EventStorageImpl) GetByID(ctx context.Context, eventID int) (model.Event, error) {
	var event model.Event
	err := storage.db.QueryRow(ctx, `
SELECT id, user_id, title, description, date
FROM event
WHERE id = $1
`, eventID).Scan(&event.ID, &event.UserID, &event.Title, &event.Description, &event.Date)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return event, apperror.NotFound.WithError(err)
		}

		return event, apperror.Internal.WithError(err)
	}

	return event, nil
}

func (storage *EventStorageImpl) Update(ctx context.Context, event model.Event) error {
	_, err := storage.db.Exec(ctx, `
UPDATE event
SET title       = $1,
    description = $2,
    date        = $3
WHERE id = $4
`, event.Title, event.Description, event.Date, event.ID)
	if err != nil {
		return apperror.Internal.WithError(err)
	}

	return nil
}

func (storage *EventStorageImpl) Delete(ctx context.Context, eventID int) error {
	_, err := storage.db.Exec(ctx, `
DELETE
FROM event
WHERE id = $1
`, eventID)
	if err != nil {
		return apperror.Internal.WithError(err)
	}

	return nil
}

func (storage *EventStorageImpl) EventsForDay(ctx context.Context, day time.Time) ([]model.Event, error) {
	return storage.eventsForUnit(ctx, "day", day)
}

func (storage *EventStorageImpl) EventsForWeek(ctx context.Context, week time.Time) ([]model.Event, error) {
	return storage.eventsForUnit(ctx, "week", week)
}

func (storage *EventStorageImpl) EventsForMonth(ctx context.Context, month time.Time) ([]model.Event, error) {
	return storage.eventsForUnit(ctx, "month", month)
}

func (storage *EventStorageImpl) eventsForUnit(ctx context.Context, unit string, date time.Time) ([]model.Event, error) {
	rows, err := storage.db.Query(ctx, `
SELECT id, user_id, title, description, date
FROM event
WHERE DATE_TRUNC($1, date) = DATE_TRUNC($1, $2::TIMESTAMPTZ)
`, unit, date)
	if err != nil {
		return []model.Event{}, apperror.Internal.WithError(err)
	}

	var events []model.Event

	for rows.Next() {
		var event model.Event
		err = rows.Scan(&event.ID, &event.UserID, &event.Title, &event.Description, &event.Date)
		if err != nil {
			return []model.Event{}, apperror.Internal.WithError(err)
		}

		events = append(events, event)
	}

	return events, nil
}
