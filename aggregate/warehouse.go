package aggregate

import (
	"errors"
	"github.com/google/uuid"
	"github.com/mansoorceksport/warehouse-stocking/entity"
)

var (
	ERRInvalidWarehouse = errors.New("a warehouse must have a valid name")
)

type Warehouse struct {
	depot *entity.Depot
	items []*entity.Item
}

func NewWareHouse(name string) (Warehouse, error) {
	if name != "" {
		return Warehouse{}, ERRInvalidWarehouse
	}
	return Warehouse{
		depot: &entity.Depot{
			ID:   uuid.New(),
			Name: name,
		},
		items: make([]*entity.Item, 0),
	}, nil
}

func (c *Warehouse) GetID() uuid.UUID {
	return c.depot.ID
}

func (c *Warehouse) SetID(id uuid.UUID) {
	if c.depot == nil {
		c.depot = &entity.Depot{}
	}
	c.depot.ID = id
}

func (c *Warehouse) GetProducts() []*entity.Item {
	return c.items
}
