package store

import (
	"github.com/mansoorceksport/warehouse-stocking/aggregate"
	"github.com/mansoorceksport/warehouse-stocking/repository/storeinventory"
	memoryStoreInventory "github.com/mansoorceksport/warehouse-stocking/repository/storeinventory/memory"
	"github.com/mansoorceksport/warehouse-stocking/service/stock"
	"github.com/mansoorceksport/warehouse-stocking/service/warehouse"
)

type Configuration func(stock *Store) error

type Store struct {
	stock          *stock.Stock
	storeInventory storeinventory.Repository
}

func NewStore(configuration ...Configuration) *Store {
	store := &Store{}
	for _, c := range configuration {
		err := c(store)
		if err != nil {
			return nil
		}
	}
	return store
}

func WithMemoryStoreInventory() Configuration {
	return func(store *Store) error {
		store.storeInventory = memoryStoreInventory.NewMemoryStoreInventory()
		return nil
	}
}

func WithStockService(wh *warehouse.Warehouse) Configuration {
	return func(store *Store) error {
		st, err := stock.NewStock(wh)
		if err != nil {
			return err
		}
		store.stock = st
		return nil
	}
}

func (s *Store) RequestStock(requestProducts []aggregate.Product) error {
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
