package depot

import (
	"context"
	"github.com/google/uuid"
	"github.com/mansoorceksport/warehouse-stocking/aggregate"
)

type Repository interface {
	Add(ctx context.Context, aw aggregate.Warehouse) error
	GetById(ctx context.Context, uuid uuid.UUID) (aggregate.Warehouse, error)
}
