package persistence

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID       bson.ObjectId `bson:"_id"`
	First    string
	Last     string
	Email    string
	Password string
	Age      int
	Bookings []Booking
}
type AdminUser struct {
	ID       bson.ObjectId `bson:"_id"`
	First    string
	Last     string
	Email    string
	Password string
	Company  string
	Events   []Event
}

func (u *User) String() string {
	return fmt.Sprintf("id: %s, first_name: %s, last_name: %s, Age: %d, Bookings: %v", u.ID, u.First, u.Last, u.Email, u.Password, u.Age, u.Bookings)
}

type Booking struct {
	Date    int64
	EventID []byte
	//EventID bson.ObjectId `bson:"_id"`
	Seats int
	Name  string
}

type UserFind struct {
	Email    string
	Password string
}

type Event struct {
	ID        bson.ObjectId `bson:"_id"`
	Name      string        `dynamodbav:"EventName"`
	Duration  int
	StartDate int64
	EndDate   int64
	Location  Location
}

type Location struct {
	ID        bson.ObjectId `bson:"_id"`
	Name      string
	Address   string
	Country   string
	OpenTime  string
	CloseTime string
	Halls     []Hall
}

type Hall struct {
	Name     string `json:"name"`
	Location string `json:"location,omitempty"`
	Capacity int    `json:"capacity"`
}
