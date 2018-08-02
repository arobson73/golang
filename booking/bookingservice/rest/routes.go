package rest

import (
	"fmt"
	"net/http"
	"time"

	"andy/booking/lib/msgqueue"
	"andy/booking/lib/persistence"

	"github.com/gorilla/mux"
)

func ServeAPI(listenAddr string, database persistence.DatabaseHandler, eventEmitter msgqueue.EventEmitter) {
	r := mux.NewRouter()
	r.Methods("post").Path("/events/{eventID}/{userID}/bookings").Handler(&CreateBookingHandler{eventEmitter, database})

	srv := http.Server{
		Handler:      r,
		Addr:         listenAddr,
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  1 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
}
