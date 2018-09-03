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
	OpenTime  string
	CloseTime string
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

type Booking struct {
	Date    int64
	EventID []byte
	Seats   int
}

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

func main() {
	//4 events as follows:
	eventNames := []string{"music gig", "cinema", "football match", "olympics"}
	eventDuration := []int{240, 120, 90, 180}
	eventStartDates := []int64{(time.Date(2018, 7, 30, 12, 0, 0, 0, time.UTC)).Unix(), (time.Date(2018, 8, 1, 14, 0, 0, 0, time.UTC).Unix()), (time.Date(2018, 8, 2, 15, 0, 0, 0, time.UTC)).Unix(), (time.Date(2018, 8, 3, 12, 0, 0, 0, time.UTC)).Unix()}
	eventEndDate := []int64{(time.Date(2018, 7, 30, 23, 30, 0, 0, time.UTC)).Unix(), (time.Date(2018, 8, 1, 17, 0, 0, 0, time.UTC).Unix()), (time.Date(2018, 8, 2, 16, 45, 0, 0, time.UTC)).Unix(), (time.Date(2018, 8, 3, 22, 30, 0, 0, time.UTC)).Unix()}
	eventIDs := []bson.ObjectId{bson.NewObjectId(), bson.NewObjectId(), bson.NewObjectId(), bson.NewObjectId()}
	for _, v := range eventIDs {
		fmt.Printf("+%v \n", v)
	}
	/*
		e0, _ := eventIDs[0].MarshalText()
		e1, _ := eventIDs[1].MarshalText()
		e2, _ := eventIDs[2].MarshalText()
		e3, _ := eventIDs[3].MarshalText()
		eventIDsBytes := [][]byte{e0, e1, e2, e3}
	*/
	//	b2, _ := b1.MarshalText()
	//eventIDSAsBytes := [][]bytes{eventIDs[0].MarshalText(), eventIDs[1].MarshalText(), eventIDs[2].MarshalText(), eventIDs[3].MarshalText()}
	locationIDs := []bson.ObjectId{bson.NewObjectId(), bson.NewObjectId(), bson.NewObjectId(), bson.NewObjectId()}
	locationName := []string{"Wembly", "Odeon", "Madjeski Stadium", "London Olympic Track"}
	locationAddress := []string{"London", "Guildford", "Reading", "East London"}
	locationCountry := []string{"UK", "UK", "UK", "UK"}
	locationOpenTime := []string{"12:30", "14:20", "15:10", "12:00"}
	locationCloseTime := []string{"23:00", "17:00", "17:00", "23:00"}
	locationHallsEvent1 := []Hall{{"Wembly east", "Gate 12", 6000}, {"Wembly south east", "Gate 13", 7000}}
	locationHallsEvent2 := []Hall{{"Screen 1", "Guildford Odeon", 300}, {"Screen 2", "Guildford Odeon", 300}}
	locationHallsEvent3 := []Hall{{"East Stand", "Gate 1", 3000}, {"South Stand", "Gate 2", 5000}}
	locationHallsEvent4 := []Hall{{"Main Gates", "Gate 1-20", 30000}, {"South Stand", "Gate 2", 5000}}
	locHalls := [][]Hall{locationHallsEvent1, locationHallsEvent2, locationHallsEvent3, locationHallsEvent4}

	events := []Event{}
	locs := []Location{}
	//users with no bookings, can add booking via rest instead
	users := []User{
		{bson.NewObjectId(), "Jim", "Smith", "jimsmith@yahoo.co.uk", "fer6et", 22, []Booking{
			/*
				{eventStartDates[0], eventIDsBytes[0], 2},
				{eventStartDates[1], eventIDsBytes[1], 3}*/},
		},
		{bson.NewObjectId(), "Paul", "Jones", "pauljones44@hotmail.com", "jonesy123", 44, []Booking{ /*
				{eventStartDates[2], eventIDsBytes[2], 2}*/},
		},
		{bson.NewObjectId(), "Frank", "Roberts", "froberts@google.co.uk", "freddie99", 31, []Booking{ /*
				{eventStartDates[3], eventIDsBytes[3], 1}*/},
		},
		{bson.NewObjectId(), "Gale", "Moor", "galeM43@google.com", "gmrr889", 28, []Booking{ /*
				{eventStartDates[1], eventIDsBytes[1], 4},
				{eventStartDates[2], eventIDsBytes[2], 2},
				{eventStartDates[3], eventIDsBytes[3], 3}*/},
		},
	}
	adminUsers := []AdminUser{
		{bson.NewObjectId(), "James", "Roberts", "jamesrobs@yahoo.co.uk", "james11", "Apple", []Event{
			/*
				{eventStartDates[0], eventIDsBytes[0], 2},
				{eventStartDates[1], eventIDsBytes[1], 3}*/},
		},
		{bson.NewObjectId(), "Mark", "Blake", "markblake4@hotmail.com", "mblake2", "Google", []Event{ /*
				{eventStartDates[2], eventIDsBytes[2], 2}*/},
		},
	}

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
	//create some users
	for i := 0; i < len(users); i++ {
		data, err := json.Marshal(users[i])
		if err != nil {
			fmt.Println("error:", err)
		}
		name := "user"
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
	//create some admin users
	for i := 0; i < len(adminUsers); i++ {
		data, err := json.Marshal(adminUsers[i])
		if err != nil {
			fmt.Println("error:", err)
		}
		name := "adminuser"
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
