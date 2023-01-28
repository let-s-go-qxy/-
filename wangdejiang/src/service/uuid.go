package service

import (
	"github.com/gofrs/uuid"
	"log"
)

// Create a Version 4 UUID, panicking on error.
// Use this form to initialize package-level variables.
var u1 = uuid.Must(uuid.NewV4())

// CreateUuid Create a Version 4 UUID.
func CreateUuid() string {
	v4, _ := uuid.NewV4()
	return v4.String()
}

// ParseUuid Parse a UUID from a string.
func ParseUuid(str string) uuid.UUID {
	s := "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
	u, err := uuid.FromString(s)
	if err != nil {
		log.Fatalf("failed to parse UUID %q: %v", s, err)
	}
	return u
}
