package listener

import (
	"andy/booking/contracts"
	"andy/booking/lib/msgqueue"
	"andy/booking/lib/persistence"
	"encoding/hex"
	"fmt"
	"log"

	"gopkg.in/mgo.v2/bson"
)

type BookingProcessor struct {
	BookingListener msgqueue.EventListener
	Database        persistence.DatabaseHandler
}

func (p *BookingProcessor) ProcessEvents() {

	log.Println("listening or events")
	received, errors, err := p.BookingListener.Listen("eventBooked")

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

func (p *BookingProcessor) handleEvent(event msgqueue.Event) {
	switch e := event.(type) {
	case *contracts.EventBookedEvent:
		log.Printf("booking event % created: %s\n", e.EventID, e.EventName)
		//ObjectIdHex(s string) ObjectId
		if !bson.IsObjectIdHex(e.EventID) {
			log.Printf("event %v did not contain valid object ID", e)
			return
		}
		eventId := bson.ObjectIdHex(e.EventID)
		eventbytes, _ := eventId.MarshalText()
		booking := persistence.Booking{
			Date:    e.Date.Unix(),
			EventID: eventbytes,
			Seats:   e.Seats,
		}
		userIDBytes, _ := hex.DecodeString(e.UserID)
		err := p.Database.AddBookingForUser(userIDBytes, booking)
		if err != nil {
			log.Println("error adding booking to db")
		}

	default:
		log.Printf("unknown event type: %T", e)
	}
}
