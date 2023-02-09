package warehouseinventory

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/mansoorceksport/warehouse-stocking/aggregate"
)

var (
	ErrProductNotFound      = errors.New("product not found")
	ErrProductAlreadyExists = errors.New("product already exists")
)

type Repository interface {
	GetByID(ctx context.Context, id uuid.UUID) (aggregate.Product, error)
	GetAll(ctx context.Context) []aggregate.Product
	Add(ctx context.Context, product aggregate.Product) error
	Update(ctx context.Context, product aggregate.Product) error
	Delete(ctx context.Context, id uuid.UUID) error
}
