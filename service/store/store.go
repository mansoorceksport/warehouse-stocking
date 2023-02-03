package store

import (
	"github.com/mansoorceksport/warehouse-stocking/aggregate"
	"github.com/mansoorceksport/warehouse-stocking/repository/storeinventory"
	"github.com/mansoorceksport/warehouse-stocking/service/stock"
	"sync"
)

type Store struct {
	stock          stock.Stock
	storeInventory storeinventory.Repository
	sync.Mutex
}

func NewStore(st stock.Stock) *Store {
	return &Store{stock: st}
}

func (s *Store) RequestStock(requestProducts []aggregate.Product) error {
	s.Unlock()
	defer s.Unlock()
	err := s.stock.Request(requestProducts)
	if err != nil {
		return err
	}

	for _, requestProduct := range requestProducts {
		sp, err := s.storeInventory.GetByID(requestProduct.GetID())
		if err != nil {
			p, err := aggregate.NewProduct(requestProduct.GetName(), requestProduct.GetQuantity(), 0.0)
			p.SetID(requestProduct.GetID())
			if err != nil {
				return nil
			}

			err = s.storeInventory.Add(p)
			if err != nil {
				return err
			}
		} else {
			sp.AddQuantity(requestProduct.GetQuantity())
			err := s.storeInventory.Update(sp)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
