package contracts

import "time"

// EventBookedEvent is emitted whenever an event is booked
type EventBookedEvent struct {
	EventID string    `json:"eventId"`
	UserID  string    `json:"userId"`
	Seats   int       `json:"seats"`
	Date    time.Time `json:"date"`
	Name    string    `json:"name"`
}

// EventName returns the event's name
func (c *EventBookedEvent) EventName() string {
	return "eventBooked"
}
