package main

import (
	"encoding/json"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

type Hall struct {
	Name     string `json:"name"`
	Location string `json:"location,omitempty"`
	Capacity int    `json:"capacity"`
}

type Location struct {
	ID bson.ObjectId `json:"id" bson:"_id,omniempty"`
	//ID        int
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

var (
	//locs *Location
	hall Hall
	/*
		//Hall Data
		hallLocation string
		hallName     string
		hallCapacity int
		//Location
		locationName      string
		locationAddress   string
		locationCountry   string
		locationOpenTime  int
		locationCloseTime int
		locationHalls     []Hall
	*/
)

func main() {
	jsonHalls := []byte(`[
		{"Name":"York", "Location":"Derby","Capacity":100},
		{"Name":"Jubilee", "Location":"Reading","Capacity":1000}
		]`)
	jsonLoc := `{ 
			"ID" : 123, "Name":"London","Address":"Palace Road","Country":"UK","OpenTime":9,"CloseTime":10,"Halls":
		[{"Name":"York", "Location":"Derby","Capacity":100},
		{"Name":"Jubilee", "Location":"Reading","Capacity":1000}
	]}`
	var halls []Hall
	err := json.Unmarshal(jsonHalls, &halls)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", halls)
	var loc Location
	err = json.Unmarshal([]byte(jsonLoc), &loc)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", loc)

	/*
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
	*/
}
