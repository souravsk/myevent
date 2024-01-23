package persistence

// Event represents a single event
type DatabaseHandler interface {
	AddEvent(Event) ([]byte, error)           // returns the id of the newly added event
	FindEvent([]byte) (Event, error)          // finds an event by its id
	FindEventByName(string) (Event, error)    // finds an event by its name
	FindAllAvailableEvents() ([]Event, error) // returns the list of all events
}
