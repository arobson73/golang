package listener

import (
	"fmt"
	"log"

	"andy/booking/contracts"
	"andy/booking/lib/msgqueue"
	"andy/booking/lib/persistence"

	"gopkg.in/mgo.v2/bson"
)

type EventProcessor struct {
	EventListener msgqueue.EventListener
	Database      persistence.DatabaseHandler
}

func (p *EventProcessor) ProcessEvents() {
	log.Println("listening or events")

	received, errors, err := p.EventListener.Listen("eventCreated", "userCreated")

	if err != nil {
		panic(err)
	}

	for {
		select {
		case evt := <-received:
			fmt.Printf("got event %T: %s\n", evt, evt)
			p.handleEvent(evt)
		case err = <-errors:
			fmt.Printf("got error while receiving event: %s\n", err)
		}
	}
}

func (p *EventProcessor) handleEvent(event msgqueue.Event) {
	switch e := event.(type) {
	case *contracts.EventCreatedEvent:
		log.Printf("event %s created: %s\n", e.ID, e.Name)

		if !bson.IsObjectIdHex(e.ID) {
			log.Printf("event %v did not contain valid object ID", e)
			return
		}

		p.Database.AddEvent(persistence.Event{ID: bson.ObjectIdHex(e.ID), Name: e.Name})
	case *contracts.UserCreatedEvent:
		log.Printf("event %s created: %s\n", e.ID, e.First)
		if !bson.IsObjectIdHex(e.ID) {
			log.Printf("event %v did not contain valid object ID", e)
			return
		}

		p.Database.AddUser(persistence.User{ID: bson.ObjectIdHex(e.ID), First: e.First, Last: e.Last, Age: e.Age, Bookings: e.Bookings})

	case *contracts.LocationCreatedEvent:
		log.Printf("location %s created: %v", e.ID, e)
		// TODO: No persistence for locations, yet
	default:
		log.Printf("unknown event type: %T", e)
	}
}
