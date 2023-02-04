package store

import (
	"errors"
	"github.com/mansoorceksport/warehouse-stocking/aggregate"
	storeRepository "github.com/mansoorceksport/warehouse-stocking/repository/store"
	memoryStoreRepository "github.com/mansoorceksport/warehouse-stocking/repository/store/memory"
	storeInventoryRepository "github.com/mansoorceksport/warehouse-stocking/repository/storeinventory"
	memoryStoreInventory "github.com/mansoorceksport/warehouse-stocking/repository/storeinventory/memory"
	"github.com/mansoorceksport/warehouse-stocking/service/warehouse"
)

var (
	ErrRequestProductEmpty = errors.New("request product empty")
)

type Configuration func(stock *Store) error

type Store struct {
	storeRepository          storeRepository.Repository
	warehouseService         *warehouse.Warehouse
	storeInventoryRepository storeInventoryRepository.Repository
}

func NewStore(configuration ...Configuration) *Store {
	s := &Store{}
	for _, c := range configuration {
		err := c(s)
		if err != nil {
			return nil
		}
	}
	return s
}

func WithMemoryStoreInventoryRepository() Configuration {
	return func(store *Store) error {
		store.storeInventoryRepository = memoryStoreInventory.NewMemoryStoreInventory()
		return nil
	}
}

func WithMemoryStoreRepository() Configuration {
	return func(store *Store) error {
		store.storeRepository = memoryStoreRepository.NewStoreMemoryRepository()
		return nil
	}
}

func WithWarehouseService(wh *warehouse.Warehouse) Configuration {
	return func(store *Store) error {
		store.warehouseService = wh
		return nil
	}
}

func (s *Store) RequestStock(requestProducts []aggregate.Product) error {
	if len(requestProducts) == 0 {
		return ErrRequestProductEmpty
	}

	err := s.warehouseService.ProcessStockRequest(requestProducts)
	if err != nil {
		return err
	}

	for _, requestProduct := range requestProducts {
		sp, err := s.storeInventoryRepository.GetByID(requestProduct.GetID())
		if err != nil {
			p, err := aggregate.NewProduct(requestProduct.GetName(), requestProduct.GetQuantity(), 0.0)
			p.SetID(requestProduct.GetID())
			if err != nil {
				return nil
			}

			err = s.storeInventoryRepository.Add(p)
			if err != nil {
				return err
			}
		} else {
			sp.AddQuantity(requestProduct.GetQuantity())
			err := s.storeInventoryRepository.Update(sp)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
