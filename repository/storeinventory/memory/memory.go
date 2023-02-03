package storeinventory

import (
	"github.com/google/uuid"
	"github.com/mansoorceksport/warehouse-stocking/aggregate"
	"github.com/mansoorceksport/warehouse-stocking/repository/storeinventory"
	"sync"
)

type Memory struct {
	storeProducts map[uuid.UUID]aggregate.Product
	sync.Mutex
}

func NewMemoryStoreInventory() *Memory {
	return &Memory{
		storeProducts: make(map[uuid.UUID]aggregate.Product),
	}
}

func (m *Memory) GetByID(id uuid.UUID) (aggregate.Product, error) {
	if p, ok := m.storeProducts[id]; ok {
		return p, nil
	}

	return aggregate.Product{}, storeinventory.ErrProductNotFound
}

func (m *Memory) GetAll() []aggregate.Product {
	var products []aggregate.Product
	for _, p := range m.storeProducts {
		products = append(products, p)
	}
	return products
}

func (m *Memory) Add(product aggregate.Product) error {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.storeProducts[product.GetID()]; ok {
		return storeinventory.ErrProductAlreadyExists
	}
	m.storeProducts[product.GetID()] = product

	return nil
}

func (m *Memory) Update(product aggregate.Product) error {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.storeProducts[product.GetID()]; !ok {
		return storeinventory.ErrProductNotFound
	}
	m.storeProducts[product.GetID()] = product
	return nil
}

func (m *Memory) Delete(id uuid.UUID) error {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.storeProducts[id]; !ok {
		return storeinventory.ErrProductNotFound
	}
	delete(m.storeProducts, id)
	return nil
}

func (m *Memory) StoreInventory() string {
	return "Store Inventory"
}
