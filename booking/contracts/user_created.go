package contracts

import "andy/booking/lib/persistence"

//UserCreatedEvent is emitted whenever a new user is created
type UserCreatedEvent struct {
	ID       string                `json:"id"`
	First    string                `json:"first"`
	Last     string                `json:"last"`
	Email    string                `json:"email"`
	Password string                `json:"password"`
	Age      int                   `json:"age"`
	Bookings []persistence.Booking `json:"bookings"`
}

func (u *UserCreatedEvent) EventName() string {
	return "userCreated"
}
