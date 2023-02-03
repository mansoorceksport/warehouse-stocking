package entity

import "github.com/google/uuid"

type Depot struct {
	ID   uuid.UUID
	Name string
}
