package rest

import (
	"net/http"

	"andy/booking/lib/msgqueue"
	"andy/booking/lib/persistence"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func ServeAPI(listenAddr string, database persistence.DatabaseHandler, eventEmitter msgqueue.EventEmitter) error {
	handler := newBookingHandler(database, eventEmitter)
	r := mux.NewRouter()
	bookingrouter := r.PathPrefix("/bookings").Subrouter()
	bookingrouter.Methods("POST").Path("/{eventID}/{userID}").HandlerFunc(handler.bookingHandler)

	bookingrouter.Methods("GET").Path("/{userID}").HandlerFunc(handler.bookingsForUserHandler)
	//	r.Methods("post").Path("/events/{eventID}/{userID}/bookings").Handler(&CreateBookingHandler{eventEmitter, database})
	rc := handlers.CORS()(r)
	return http.ListenAndServe(listenAddr, rc)
	/*
		srv := http.Server{
			Handler:      rc,
			Addr:         listenAddr,
			WriteTimeout: 2 * time.Second,
			ReadTimeout:  1 * time.Second,
		}
	*/
	/*
		err := srv.ListenAndServe()
		if err != nil {
			fmt.Printf("%+v\n", err)
		}*/
}
