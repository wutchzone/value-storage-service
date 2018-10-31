package value

import "github.com/gofrs/uuid"

// Value that is represented like a struct
type Value struct {
	Key   string
	Value string
	UUID  string
}

// NewValue return new struct
func NewValue(key string, value string) *Value {
	return &Value{
		Key:   key,
		Value: value,
		UUID:  uuid.Must(uuid.NewV4()).String(),
	}
}
