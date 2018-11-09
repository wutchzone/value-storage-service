package value

import (
	"time"

	"github.com/mongodb/mongo-go-driver/bson/objectid"
)

// Value that is represented like a struct
type Value struct {
	ID         objectid.ObjectID `bson:"_id,omitempty"`
	Key        string
	Value      string
	DeviceID   string
	Location   Location
	ModifiedAt time.Time
	CreatedAt  time.Time
}

// Location where value was captured
type Location struct {
	Lat string
	Lon string
}

// NewValue return new struct
func NewValue(key string, value string) *Value {
	return &Value{
		Key:   key,
		Value: value,
		ID:    objectid.New(),
	}
}
