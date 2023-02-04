package warehouse

import (
	"github.com/google/uuid"
	"github.com/mansoorceksport/warehouse-stocking/aggregate"
)

type Repository interface {
	Add(aw aggregate.Warehouse) error
	GetById(uuid uuid.UUID) aggregate.Warehouse
}
