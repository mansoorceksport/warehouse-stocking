package warehouseinventory

import (
	"context"
	"github.com/google/uuid"
	"github.com/mansoorceksport/warehouse-stocking/aggregate"
	"github.com/mansoorceksport/warehouse-stocking/repository/warehouseinventory"
	"sync"
)

type WarehouseInventory struct {
	warehouseProducts map[uuid.UUID]aggregate.Product
	sync.Mutex
}

func NewMemoryWarehouseInventory() *WarehouseInventory {
	return &WarehouseInventory{
		warehouseProducts: make(map[uuid.UUID]aggregate.Product),
	}
}

func (m *WarehouseInventory) GetByID(_ context.Context, id uuid.UUID) (aggregate.Product, error) {
	if p, ok := m.warehouseProducts[id]; ok {
		return p, nil
	}

	return aggregate.Product{}, warehouseinventory.ErrProductNotFound
}

func (m *WarehouseInventory) GetAll(_ context.Context) []aggregate.Product {
	var products []aggregate.Product
	for _, p := range m.warehouseProducts {
		products = append(products, p)
	}
	return products
}

func (m *WarehouseInventory) Add(_ context.Context, product aggregate.Product) error {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.warehouseProducts[product.GetID()]; ok {
		return warehouseinventory.ErrProductAlreadyExists
	}
	m.warehouseProducts[product.GetID()] = product

	return nil
}

func (m *WarehouseInventory) Update(_ context.Context, product aggregate.Product) error {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.warehouseProducts[product.GetID()]; !ok {
		return warehouseinventory.ErrProductNotFound
	}
	m.warehouseProducts[product.GetID()] = product
	return nil
}

func (m *WarehouseInventory) Delete(_ context.Context, id uuid.UUID) error {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.warehouseProducts[id]; !ok {
		return warehouseinventory.ErrProductNotFound
	}
	delete(m.warehouseProducts, id)
	return nil
}
