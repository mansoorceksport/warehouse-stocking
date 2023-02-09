package memory

import (
	"context"
	"github.com/google/uuid"
	"github.com/mansoorceksport/warehouse-stocking/aggregate"
	"github.com/mansoorceksport/warehouse-stocking/repository/depot"
	"sync"
)

type Memory struct {
	warehouses map[uuid.UUID]aggregate.Warehouse
	sync.Mutex
}

func NewMemoryWareHouse() depot.Repository {
	return &Memory{
		warehouses: make(map[uuid.UUID]aggregate.Warehouse),
	}
}

func (m *Memory) Add(_ context.Context, aw aggregate.Warehouse) error {
	m.Lock()
	defer m.Unlock()
	m.warehouses[aw.GetID()] = aw
	return nil
}

func (m *Memory) GetById(_ context.Context, uuid uuid.UUID) (aggregate.Warehouse, error) {
	return m.warehouses[uuid], nil
}
