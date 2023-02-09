package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/mansoorceksport/warehouse-stocking/aggregate"
	"github.com/mansoorceksport/warehouse-stocking/common/postgres"
	"github.com/mansoorceksport/warehouse-stocking/repository/warehouseinventory"
)

type WarehouseInventory struct {
	postgres *postgres.Postgres
}

func New(postgres *postgres.Postgres) warehouseinventory.Repository {
	return &WarehouseInventory{postgres: postgres}
}

func (w *WarehouseInventory) GetByID(_ context.Context, id uuid.UUID) (aggregate.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (w *WarehouseInventory) GetAll(_ context.Context) []aggregate.Product {
	//TODO implement me
	panic("implement me")
}

func (w *WarehouseInventory) Add(_ context.Context, product aggregate.Product) error {
	//TODO implement me
	panic("implement me")
}

func (w *WarehouseInventory) Update(_ context.Context, product aggregate.Product) error {
	//TODO implement me
	panic("implement me")
}

func (w *WarehouseInventory) Delete(_ context.Context, id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
