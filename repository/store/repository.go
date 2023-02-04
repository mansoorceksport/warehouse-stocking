package store

import (
	"github.com/google/uuid"
	"github.com/mansoorceksport/warehouse-stocking/aggregate"
)

type Repository interface {
	Add(s aggregate.Store) error
	GetById(uuid uuid.UUID) aggregate.Store
}
