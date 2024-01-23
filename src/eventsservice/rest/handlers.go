package rest

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/souravsk/myevent/src/lib/persistence"
)

// this will be the handler for all /events requests
type eventServiceHandler struct {
	dbhandler persistence.DatabaseHandler // this handler can handle all db related functions
}

// NewEventHandler returns a new instance of eventServiceHandler
func NewEventHandler(databasehandler persistence.DatabaseHandler) *eventServiceHandler {
	return &eventServiceHandler{ // return the address of the struct
		dbhandler: databasehandler, // set the handler
	}
}

// FindEventHandler returns a single event by id or name
func (eh *eventServiceHandler) FindEventHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	criteria, ok := vars["SearchCriteria"] // SearchCriteria is a variable in the route
	if !ok {
		w.WriteHeader(400)
		fmt.Fprint(w, `{"error": "No search criteria found, you can either
						search by id via /id/4
						to search by name via /name/coldplayconcert}"`)
		return
	}

	searchkey, ok := vars["search"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprint(w, `{"error": "No search keys found, you can either search
						by id via /id/4
						to search by name via /name/coldplayconcert}"`)
		return
	}

	var event persistence.Event // this will hold the event data
	var err error               // this will hold the error in case something goes wrong

	switch strings.ToLower(criteria) {
	case "name":
		event, err = eh.dbhandler.FindEventByName(searchkey) // search by name
	case "id":
		id, err := hex.DecodeString(searchkey) // search by id
		if err == nil {
			event, err = eh.dbhandler.FindEvent(id)
		}
	}
	if err != nil {
		fmt.Fprintf(w, `{"error": "%s"}`, err)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	json.NewEncoder(w).Encode(&event) // encode the event data into JSON and send it back
}

// AllEventHandler returns all events
func (eh *eventServiceHandler) AllEventHandler(w http.ResponseWriter, r *http.Request) {
	events, err := eh.dbhandler.FindAllAvailableEvents()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error": "Error occured while trying to find all available events %s"}`, err)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	err = json.NewEncoder(w).Encode(&events)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error": "Error occured while trying encode events to JSON %s"}`, err)
	}
}

// NewEventHandler adds a new event
func (eh *eventServiceHandler) NewEventHandler(w http.ResponseWriter, r *http.Request) {
	event := persistence.Event{}
	err := json.NewDecoder(r.Body).Decode(&event)
	if nil != err {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error": "error occured while decoding event data %s"}`, err)
		return
	}
	id, err := eh.dbhandler.AddEvent(event)
	if nil != err {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error": "error occured while persisting event %d %s"}`, id, err)
		return
	}
	fmt.Fprint(w, `{"id":%d}`, id) // return the id of the newly added event
}
