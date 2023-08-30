package dto

import (
	"fmt"
	"time"
)

type CreateEvent struct {
	UserID      int
	Title       string
	Description string
	Date        time.Time
}

func (request *CreateEvent) Validate() error {
	if request.UserID == 0 {
		return fmt.Errorf("user_id is required")
	}

	if request.Title == "" {
		return fmt.Errorf("title is required")
	}

	if request.Date.IsZero() {
		return fmt.Errorf("date is required")
	}

	return nil
}

type UpdateEvent struct {
	EventID     int
	Title       string
	Description string
	Date        time.Time
}

func (request *UpdateEvent) Validate() error {
	if request.EventID == 0 {
		return fmt.Errorf("event_id is required")
	}

	return nil
}
