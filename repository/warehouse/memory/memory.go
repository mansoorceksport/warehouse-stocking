package memory

import (
	"github.com/google/uuid"
	"github.com/mansoorceksport/warehouse-stocking/aggregate"
	"github.com/mansoorceksport/warehouse-stocking/repository/warehouse"
	"sync"
)

type Memory struct {
	warehouses map[uuid.UUID]aggregate.Warehouse
	sync.Mutex
}

func NewMemoryWareHouse() warehouse.Repository {
	return &Memory{
		warehouses: make(map[uuid.UUID]aggregate.Warehouse),
	}
}

func (m *Memory) Add(aw aggregate.Warehouse) error {
	m.Lock()
	defer m.Unlock()
	m.warehouses[aw.GetID()] = aw
	return nil
}

func (m *Memory) GetById(uuid uuid.UUID) aggregate.Warehouse {
	return m.warehouses[uuid]
}
