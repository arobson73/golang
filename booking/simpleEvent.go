package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"gopkg.in/mgo.v2/bson"
)

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

func main() {
	//4 events as follows:
	eventNames := []string{"music gig", "cinema", "football match", "olympics"}
	eventDuration := []int{240, 120, 90, 180}
	eventStartDates := []int64{(time.Date(2018, 7, 30, 12, 0, 0, 0, time.UTC)).Unix(), (time.Date(2018, 8, 1, 14, 0, 0, 0, time.UTC).Unix()), (time.Date(2018, 8, 2, 15, 0, 0, 0, time.UTC)).Unix(), (time.Date(2018, 8, 3, 12, 0, 0, 0, time.UTC)).Unix()}
	eventEndDate := []int64{(time.Date(2018, 7, 30, 23, 30, 0, 0, time.UTC)).Unix(), (time.Date(2018, 8, 1, 17, 0, 0, 0, time.UTC).Unix()), (time.Date(2018, 8, 2, 16, 45, 0, 0, time.UTC)).Unix(), (time.Date(2018, 8, 3, 22, 30, 0, 0, time.UTC)).Unix()}
	eventIDs := []bson.ObjectId{bson.NewObjectId(), bson.NewObjectId(), bson.NewObjectId(), bson.NewObjectId()}
	locationIDs := []bson.ObjectId{bson.NewObjectId(), bson.NewObjectId(), bson.NewObjectId(), bson.NewObjectId()}
	locationName := []string{"Wembly", "Odeon", "Madjeski Stadium", "London Olympic Track"}
	locationAddress := []string{"London", "Guildford", "Reading", "East London"}
	locationCountry := []string{"UK", "UK", "UK", "UK"}
	locationOpenTime := []int{12, 14, 15, 12}
	locationCloseTime := []int{23, 17, 17, 23}
	locationHallsEvent1 := []Hall{{"Wembly east", "Gate 12", 6000}, {"Wembly south east", "Gate 13", 7000}}
	locationHallsEvent2 := []Hall{{"Screen 1", "Guildford Odeon", 300}, {"Screen 2", "Guildford Odeon", 300}}
	locationHallsEvent3 := []Hall{{"East Stand", "Gate 1", 3000}, {"South Stand", "Gate 2", 5000}}
	locationHallsEvent4 := []Hall{{"Main Gates", "Gate 1-20", 30000}, {"South Stand", "Gate 2", 5000}}
	locHalls := [][]Hall{locationHallsEvent1, locationHallsEvent2, locationHallsEvent3, locationHallsEvent4}

	events := []Event{}
	locs := []Location{}
	for i := 0; i < len(eventNames); i++ {
		var event Event
		event.ID = eventIDs[i]
		event.Name = eventNames[i]
		event.Duration = eventDuration[i]
		event.StartDate = eventStartDates[i]
		event.EndDate = eventEndDate[i]
		var loc Location
		loc.ID = locationIDs[i]
		loc.Name = locationName[i]
		loc.Address = locationAddress[i]
		loc.Country = locationCountry[i]
		loc.OpenTime = locationOpenTime[i]
		loc.CloseTime = locationCloseTime[i]
		var halls = []Hall{}
		halls = locHalls[i]
		loc.Halls = halls
		event.Location = loc
		events = append(events, event)
		locs = append(locs, loc)

	}
	//write each event to a file
	for i := 0; i < len(events); i++ {

		data, err := json.Marshal(events[i])
		if err != nil {
			fmt.Println("error:", err)
		}
		name := "out"
		num := strconv.Itoa(i)
		name += num + ".json"
		//name = append(name, num)

		fo, errf := os.Create(name)
		if errf != nil {
			panic(err)
		}
		defer fo.Close()
		fmt.Fprintf(fo, string(data[:]))

	}
	//write each location to a file
	for i := 0; i < len(events); i++ {
		data, err := json.Marshal(locs[i])
		if err != nil {
			fmt.Println("error:", err)
		}
		name := "loc"
		num := strconv.Itoa(i)
		name += num + ".json"
		//name = append(name, num)

		fo, errf := os.Create(name)
		if errf != nil {
			panic(err)
		}
		defer fo.Close()
		fmt.Fprintf(fo, string(data[:]))

	}
	//this just writes whole json to stdout as well
	b, err := json.Marshal(events)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}
