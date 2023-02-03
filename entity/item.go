package entity

import "github.com/google/uuid"

type Item struct {
	ID       uuid.UUID
	Name     string
	Quantity int
	Price    float64
}
