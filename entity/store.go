package entity

import "github.com/google/uuid"

type Store struct {
	ID   uuid.UUID
	Name string
}
