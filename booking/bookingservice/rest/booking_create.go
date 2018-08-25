package rest

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"andy/booking/contracts"
	"andy/booking/lib/msgqueue"
	"andy/booking/lib/persistence"

	"github.com/gorilla/mux"
)

type eventRef struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
}

type createBookingRequest struct {
	Seats int    `json:"seats"`
	Name  string `json:"name"`
}

type createBookingResponse struct {
	ID    string   `json:"id"`
	Event eventRef `json:"event"`
}

type CreateBookingHandler struct {
	eventEmitter msgqueue.EventEmitter
	database     persistence.DatabaseHandler
}

func newBookingHandler(databaseHandler persistence.DatabaseHandler, eventEmitter msgqueue.EventEmitter) *CreateBookingHandler {
	return &CreateBookingHandler{
		eventEmitter: eventEmitter,
		database:     databaseHandler,
	}
}

func (h *CreateBookingHandler) bookingsForUserHandler(w http.ResponseWriter, r *http.Request) {
	routeVars := mux.Vars(r)
	userID, ok := routeVars["userID"]
	log.Printf("userID=%s\n", userID)
	if !ok {
		w.WriteHeader(400)
		fmt.Fprint(w, "missing route parameter 'userID")
		return
	}
	userIDMongo, _ := hex.DecodeString(userID)
	bookings, err := h.database.FindBookingsForUser(userIDMongo)
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprintf(w, "bookings for %s could not be loaded: %s", userIDMongo, err)
		return
	}
	//	log.Printf("\nbookings are:\n")
	//log.Printf("len bookings =%d", len(bookings))
	for _, b := range bookings {
		log.Printf("%d \n", b.Date)

	}
	//	log.Printf("\ndone")
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	err = json.NewEncoder(w).Encode(bookings)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Error occured while trying to encode bookings to JSON %s", err)
	}

}

func (h *CreateBookingHandler) bookingHandler(w http.ResponseWriter, r *http.Request) {
	//func (h *CreateBookingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	routeVars := mux.Vars(r)
	eventID, ok := routeVars["eventID"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprint(w, "missing route parameter 'eventID'")
		return
	}
	userID, ok2 := routeVars["userID"]
	if !ok2 {
		w.WriteHeader(400)
		fmt.Fprint(w, "missing route parameter 'userID'")
		return
	}
	//	fmt.Fprintf(w, "%s %s", eventID, userID)

	eventIDMongo, _ := hex.DecodeString(eventID)
	//log.Println("eventIDMongo=", string(eventIDMongo[:]))
	//	log.Println("eventIDMongoS=", eventID)
	//log.Println("eventIDMongoSS=", hex.EncodeToString(eventIDMongo))

	event, err := h.database.FindEvent(eventIDMongo)
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprintf(w, "event %s could not be loaded: %s", eventID, err)
		return
	}
	//userIDMongo, _ := hex.DecodeString(userID) //note we assume the user id is a valid one

	//from the json post
	bookingRequest := createBookingRequest{}
	err = json.NewDecoder(r.Body).Decode(&bookingRequest)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "could not decode JSON body: %s", err)
		return
	}

	if bookingRequest.Seats <= 0 {
		w.WriteHeader(400)
		fmt.Fprintf(w, "seat number must be positive (was %d)", bookingRequest.Seats)
		return
	}

	eventIDAsBytes, _ := event.ID.MarshalText()
	//	log.Println("eventIDAsBytes=", string(eventIDAsBytes[:]))
	//log.Println("eventIDAsBytesS=", event.ID)
	//log.Println("eventIDAsBytesSS=", event.ID.String())
	//log.Println("eventIDAsBytesSSS=", event.ID.Hex())

	tn := time.Now()
	booking := persistence.Booking{
		Date:    tn.Unix(),
		EventID: eventIDAsBytes,
		Seats:   bookingRequest.Seats,
		Name:    bookingRequest.Name,
	}

	//need to get userid. in real use case user would be logged in,
	//so this userid could be supplied from the url

	msg := contracts.EventBookedEvent{
		EventID: event.ID.Hex(),
		//UserID:  "someUserID",

		UserID: userID,
		Seats:  booking.Seats,
		Date:   tn,
		Name:   booking.Name,
	}
	h.eventEmitter.Emit(&msg)

	//err = h.database.AddBookingForUser([]byte("someUserID"), booking)

	userIDMongo, _ := hex.DecodeString(userID)
	err = h.database.AddBookingForUser(userIDMongo, booking)

	if err != nil {
		log.Println("error adding booking to db")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	json.NewEncoder(w).Encode(&booking)
}
