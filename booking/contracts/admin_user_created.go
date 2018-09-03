package contracts

import "andy/booking/lib/persistence"

//UserCreatedEvent is emitted whenever a new user is created
type AdminUserCreatedEvent struct {
	ID       string              `json:"id"`
	First    string              `json:"first"`
	Last     string              `json:"last"`
	Email    string              `json:"email"`
	Password string              `json:"password"`
	Company  string              `json:"company"`
	Events   []persistence.Event `json:"events"`
}

func (u *AdminUserCreatedEvent) EventName() string {
	return "adminUserCreated"
}
