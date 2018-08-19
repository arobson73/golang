package rest

import (
	"net/http"

	"andy/booking/lib/msgqueue"
	"andy/booking/lib/persistence"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func ServeAPI(endpoint string, dbHandler persistence.DatabaseHandler, eventEmitter msgqueue.EventEmitter) error {
	handler := newEventHandler(dbHandler, eventEmitter)

	r := mux.NewRouter()
	eventsrouter := r.PathPrefix("/events").Subrouter()
	eventsrouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.findEventHandler)
	eventsrouter.Methods("GET").Path("").HandlerFunc(handler.allEventHandler)
	eventsrouter.Methods("GET").Path("/{eventID}").HandlerFunc(handler.oneEventHandler)
	eventsrouter.Methods("POST").Path("").HandlerFunc(handler.newEventHandler)

	locationRouter := r.PathPrefix("/locations").Subrouter()
	locationRouter.Methods("GET").Path("").HandlerFunc(handler.allLocationsHandler)
	locationRouter.Methods("POST").Path("").HandlerFunc(handler.newLocationHandler)

	//users not sure about this yet (as in how it fits into grand scheme of things and whether the route is correct etc
	userRouter := r.PathPrefix("/users").Subrouter() //this uses AddUser for db

	userRouter.Methods("POST").Path("").HandlerFunc(handler.newUserHandler)
	userRouter.Methods("GET").Path("/findUser/{firstname}/{secondname}").HandlerFunc(handler.findUserHandler)
	userRouter.Methods("GET").Path("/findUserEmailPass/{email}/{password}").HandlerFunc(handler.findUserEmailPassHandler)
	rc := handlers.CORS()(r)

	return http.ListenAndServe(endpoint, rc)
}
