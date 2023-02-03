package warehouseinventory

import (
	"github.com/google/uuid"
	"github.com/mansoorceksport/warehouse-stocking/aggregate"
	"github.com/mansoorceksport/warehouse-stocking/repository/warehouseinventory"
	"sync"
)

type Memory struct {
	products map[uuid.UUID]aggregate.Product
	sync.Mutex
}

func NewMemoryWarehouseInventory() *Memory {
	return &Memory{
		products: make(map[uuid.UUID]aggregate.Product),
	}
}

func (m *Memory) GetByID(id uuid.UUID) (aggregate.Product, error) {
	if p, ok := m.products[id]; ok {
		return p, nil
	}

	return aggregate.Product{}, warehouseinventory.ErrProductNotFound
}

func (m *Memory) GetAll() []aggregate.Product {
	var products []aggregate.Product
	for _, p := range m.products {
		products = append(products, p)
	}
	return products
}

func (m *Memory) Add(product aggregate.Product) error {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.products[product.GetID()]; ok {
		return warehouseinventory.ErrProductAlreadyExists
	}
	m.products[product.GetID()] = product

	return nil
}

func (m *Memory) Update(product aggregate.Product) error {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.products[product.GetID()]; !ok {
		return warehouseinventory.ErrProductNotFound
	}
	m.products[product.GetID()] = product
	return nil
}

func (m *Memory) Delete(id uuid.UUID) error {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.products[id]; !ok {
		return warehouseinventory.ErrProductNotFound
	}
	delete(m.products, id)
	return nil
}