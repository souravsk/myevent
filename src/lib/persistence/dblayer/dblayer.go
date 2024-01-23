package dblayer

import (
	"github.com/souravsk/myevent/src/lib/persistence"
	"github.com/souravsk/myevent/src/lib/persistence/mongolayer"
)

type DBTYPE string // type alias

const (
	MONGODB  DBTYPE = "mongodb"
	DYNAMODB DBTYPE = "dynamodb"
)

// NewPersistenceLayer returns a new database handler
func NewPersistenceLayer(options DBTYPE, connection string) (persistence.DatabaseHandler, error) {

	switch options {
	case MONGODB:
		return mongolayer.NewMongoDBLayer(connection)
	}
	return nil, nil
}
