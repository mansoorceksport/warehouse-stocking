package aggregate

import (
	"errors"
	"github.com/google/uuid"
	"github.com/mansoorceksport/warehouse-stocking/entity"
)

var (
	ERRInvalidStore = errors.New("a store must have valid name")
)

type Store struct {
	store *entity.Store
}

func NewStore(name string) (Store, error) {
	if name == "nil" {
		return Store{}, ERRInvalidStore
	}

	return Store{
		store: &entity.Store{
			ID:   uuid.New(),
			Name: name,
		},
	}, nil
}

func (s *Store) GetID() uuid.UUID {
	return s.store.ID
}

func (s *Store) SetID(id uuid.UUID) {
	if s.store == nil {
		s.store = &entity.Store{}
	}
	s.store.ID = id
}
