package memory

import (
	"github.com/google/uuid"
	"github.com/mansoorceksport/warehouse-stocking/aggregate"
	"github.com/mansoorceksport/warehouse-stocking/repository/store"
	"sync"
)

type Memory struct {
	Stores map[uuid.UUID]aggregate.Store
	sync.Mutex
}

func NewStoreMemoryRepository() store.Repository {
	return &Memory{
		Stores: make(map[uuid.UUID]aggregate.Store),
	}
}

func (m *Memory) Add(s aggregate.Store) error {
	m.Lock()
	defer m.Unlock()
	m.Stores[s.GetID()] = s
	return nil
}

func (m *Memory) GetById(uuid uuid.UUID) aggregate.Store {
	return m.Stores[uuid]
}
