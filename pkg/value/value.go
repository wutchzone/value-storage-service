package value

import (
	"time"

	"github.com/mongodb/mongo-go-driver/bson/objectid"
)

type key int

const (
	// ValueKey for context
	ValueKey key = iota
	// FilterKey for context
	FilterKey
)

// Filter for searcing in the DB
type Filter struct {
	Asc  bool
	UID  string
	From time.Time
	To   time.Time
}

// Value that is represented like a struct
type Value struct {
	ID         objectid.ObjectID `bson:"_id,omitempty"`
	Key        string
	Value      string
	DeviceID   string `json:",omitempty"`
	Location   Location
	ModifiedAt time.Time
	CreatedAt  time.Time
}

// Location where value was captured
type Location struct {
	Lat string `json:",omitempty"`
	Lon string `json:",omitempty"`
}

// NewValue return new struct
func NewValue(key string, value string, loc Location) *Value {
	// var l *Location
	// if loc.Lat == "" || loc.Lon == "" {
	// 	l = &Location{
	// 		Lat: loc.Lat,
	// 		Lon: loc.Lon,
	// 	}
	// }
	return &Value{
		Key:        key,
		Value:      value,
		ID:         objectid.New(),
		Location:   loc,
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	}
}
