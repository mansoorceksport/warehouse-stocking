package aggregate

import (
	"errors"
	"github.com/google/uuid"
	"github.com/mansoorceksport/warehouse-stocking/entity"
)

var (
	ErrProductInvalid = errors.New("a product have a valid name")
)

type Product struct {
	item *entity.Item
}

func NewProduct(name string, quantity int, price float64) (Product, error) {
	if name == "" {
		return Product{}, ErrProductInvalid
	}

	return Product{
		item: &entity.Item{
			ID:       uuid.New(),
			Name:     name,
			Quantity: quantity,
			Price:    price,
		},
	}, nil
}

func (p Product) SetQuantity(q int) {
	p.item.Quantity = q
}

func (p Product) GetQuantity() int {
	return p.item.Quantity
}

func (p Product) SetName(name string) {
	p.item.Name = name
}

func (p Product) GetName() string {
	return p.item.Name
}

func (p Product) GetID() uuid.UUID {
	return p.item.ID
}

func (p Product) SetID(id uuid.UUID) {
	p.item.ID = id
}
