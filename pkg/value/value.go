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
	From *time.Time
	To   *time.Time
}

// Value that is represented like a struct
type Value struct {
	ID         objectid.ObjectID `bson:"_id,omitempty"`
	Key        string            `bson:"key"`
	Value      string            `bson:"value"`
	DeviceID   string            `bson:"device_id"`
	Location   Location          `bson:"location"`
	ModifiedAt time.Time         `bson:"modified_at"`
	CreatedAt  time.Time         `bson:"created_at"`
}

// Location where value was captured
type Location struct {
	Lat string
	Lon string
}

// NewValue return new struct
func NewValue(key string, value string, loc Location) *Value {
	return &Value{
		Key:        key,
		Value:      value,
		ID:         objectid.New(),
		Location:   loc,
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	}
}
