package persistence

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

// User represents a single user
type User struct {
	ID       bson.ObjectId `bson:"_id"`
	First    string
	Last     string
	Age      int
	Bookings []Booking
}

// String returns a string representation of the user
func (u *User) String() string {
	return fmt.Sprintf("id: %s, first_name: %s, last_name: %s, Age: %d, Bookings: %v", u.ID, u.First, u.Last, u.Age, u.Bookings)
}

// Booking represents a single booking
type Booking struct {
	Date    int64
	EventID []byte
	Seats   int
}

// Event represents a single event
type Event struct {
	ID        bson.ObjectId `bson:"_id"`
	Name      string
	Duration  int
	StartDate int64
	EndDate   int64
	Location  Location
}

// Location represents a single location
type Location struct {
	ID        bson.ObjectId `bson:"_id"`
	Name      string
	Address   string
	Country   string
	OpenTime  int
	CloseTime int
	Halls     []Hall
}

// Hall represents a single hall
type Hall struct {
	Name     string `json:"name"`
	Location string `json:"location,omitempty"`
	Capacity int    `json:"capacity"`
}
