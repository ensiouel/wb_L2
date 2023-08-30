package handler

import (
	"calendar/internal/dto"
	"calendar/internal/model"
	"calendar/internal/service"
	"calendar/pkg/httputil"
	"context"
	"github.com/ensiouel/apperror"
	"net/http"
	"strconv"
	"time"
)

const (
	unitDay   = "day"
	unitWeek  = "week"
	unitMonth = "month"
)

type EventHandler struct {
	eventService service.EventService
}

func NewEventHandler(eventService service.EventService) *EventHandler {
	return &EventHandler{eventService: eventService}
}

func (handler *EventHandler) Register(mux *http.ServeMux) {
	mux.Handle("/create_event", httputil.Logger(httputil.POST(http.HandlerFunc(handler.create))))
	mux.Handle("/update_event", httputil.Logger(httputil.POST(http.HandlerFunc(handler.update))))
	mux.Handle("/delete_event", httputil.Logger(httputil.POST(http.HandlerFunc(handler.delete))))

	mux.Handle("/events_for_day", httputil.Logger(httputil.GET(http.HandlerFunc(handler.eventsForUnit(unitDay)))))
	mux.Handle("/events_for_week", httputil.Logger(httputil.GET(http.HandlerFunc(handler.eventsForUnit(unitWeek)))))
	mux.Handle("/events_for_month", httputil.Logger(httputil.GET(http.HandlerFunc(handler.eventsForUnit(unitMonth)))))
}

func (handler *EventHandler) create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		httputil.Error(w, apperror.BadRequest.WithMessage("failed to parse form"))
		return
	}

	var userID int
	userID, err = strconv.Atoi(r.Form.Get("user_id"))
	if err != nil {
		httputil.Error(w, apperror.BadRequest.WithMessage("failed to parse user_id"))
		return
	}

	title := r.Form.Get("title")
	description := r.Form.Get("description")

	var date time.Time
	date, err = time.Parse(time.RFC3339, r.Form.Get("date"))
	if err != nil {
		httputil.Error(w, apperror.BadRequest.WithMessage("failed to parse date"))
		return
	}

	request := dto.CreateEvent{
		UserID:      userID,
		Title:       title,
		Description: description,
		Date:        date,
	}

	if err = request.Validate(); err != nil {
		httputil.Error(w, apperror.BadRequest.WithMessage(err.Error()))
		return
	}

	var event model.Event
	event, err = handler.eventService.Create(r.Context(), request)
	if err != nil {
		httputil.Error(w, err)
		return
	}

	httputil.JSON(w, http.StatusOK, map[string]any{
		"result": event,
	})
}

func (handler *EventHandler) update(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		httputil.Error(w, apperror.BadRequest.WithMessage("failed to parse form"))
		return
	}

	var eventID int
	eventID, err = strconv.Atoi(r.Form.Get("event_id"))
	if err != nil {
		httputil.Error(w, apperror.BadRequest.WithMessage("failed to parse event_id"))
		return
	}

	title := r.Form.Get("title")
	description := r.Form.Get("description")

	var date time.Time
	date, err = time.Parse(time.RFC3339, r.Form.Get("date"))
	if err != nil {
		httputil.Error(w, apperror.BadRequest.WithMessage("failed to parse date"))
		return
	}

	request := dto.UpdateEvent{
		EventID:     eventID,
		Title:       title,
		Description: description,
		Date:        date,
	}

	if err = request.Validate(); err != nil {
		httputil.Error(w, apperror.BadRequest.WithMessage(err.Error()))
		return
	}

	var event model.Event
	event, err = handler.eventService.Update(r.Context(), request)
	if err != nil {
		httputil.Error(w, err)
		return
	}

	httputil.JSON(w, http.StatusOK, map[string]any{
		"result": event,
	})
}

func (handler *EventHandler) delete(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		httputil.Error(w, apperror.BadRequest.WithMessage("failed to parse form"))
		return
	}

	var eventID int
	eventID, err = strconv.Atoi(r.Form.Get("event_id"))
	if err != nil {
		httputil.Error(w, apperror.BadRequest.WithMessage("failed to parse event_id"))
		return
	}

	err = handler.eventService.Delete(r.Context(), eventID)
	if err != nil {
		httputil.Error(w, err)
		return
	}

	httputil.JSON(w, http.StatusOK, map[string]any{
		"result": 1,
	})
}

func (handler *EventHandler) eventsForUnit(unit string) func(w http.ResponseWriter, r *http.Request) {
	var eventsGetter func(context.Context, time.Time) ([]model.Event, error)
	switch unit {
	case unitDay:
		eventsGetter = handler.eventService.EventsForDay
	case unitWeek:
		eventsGetter = handler.eventService.EventsForWeek
	case unitMonth:
		eventsGetter = handler.eventService.EventsForMonth
	}

	return func(w http.ResponseWriter, r *http.Request) {
		date, err := time.Parse(time.DateOnly, r.URL.Query().Get("date"))
		if err != nil {
			httputil.Error(w, apperror.BadRequest.WithMessage("failed to parse date"))
			return
		}

		var events []model.Event
		events, err = eventsGetter(r.Context(), date)
		if err != nil {
			httputil.Error(w, err)
			return
		}

		httputil.JSON(w, http.StatusOK, map[string]any{
			"result": events,
		})
	}
}
