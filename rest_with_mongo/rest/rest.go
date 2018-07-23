package rest

import (
	"net/http"
	//will need to edit this path depending on your code location
	"github.com/rest_with_mongo/lib/persistence"

	"github.com/gorilla/mux"
)

func ServeAPI(endpoint string, databasehandler persistence.DatabaseHandler) error {
	handler := NewEventHandler(databasehandler)
	r := mux.NewRouter()
	eventsrouter := r.PathPrefix("/events").Subrouter()
	eventsrouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.FindEventHandler)
	eventsrouter.Methods("GET").Path("").HandlerFunc(handler.AllEventHandler)
	eventsrouter.Methods("POST").Path("").HandlerFunc(handler.NewEventHandler)
	eventsrouter.Methods("DELETE").Path("/{Event}").HandlerFunc(handler.DeleteEventHandler)
	return http.ListenAndServe(endpoint, r)
}
