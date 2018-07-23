package main

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/mgo.v2/bson"
)

func main() {
	type Hall struct {
		Name     string `json:"name"`
		Location string `json:"location,omitempty"`
		Capacity int    `json:"capacity"`
	}

	type Location struct {
		ID        bson.ObjectId `json:"id" bson:"_id,omniempty"`
		Name      string
		Address   string
		Country   string
		OpenTime  int
		CloseTime int
		Halls     []Hall
	}

	type Event struct {
		ID        bson.ObjectId `json:"id" bson:"_id"`
		Name      string
		Duration  int
		StartDate int64
		EndDate   int64
		Location  Location
	}
	var hall = Hall{
		Name:     "Cesar Hall",
		Location: "second floor,room 2210",
		Capacity: 10,
	}
	var loc = Location{
		ID:   bson.NewObjectId(),
		Name: "West Street Opera House",
		Address: "11 West Street, AZ	73846",
		Country:   "U.S.A",
		OpenTime:  7,
		CloseTime: 20,
		Halls:     []Hall{hall},
	}
	var event = Event{
		ID:        bson.NewObjectId(),
		Name:      "opera aida",
		Duration:  120,
		StartDate: 768346784368,
		EndDate:   43988943,
		Location:  loc,
	}

	b, err := json.Marshal(event)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}
