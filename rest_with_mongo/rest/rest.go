package rest

import (
	"net/http"

	"github.com/PacktPublishing/Cloud-Native-programming-with-Golang/Chapter02/myevents/src/lib/persistence"

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
