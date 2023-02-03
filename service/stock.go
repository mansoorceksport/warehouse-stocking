package service

import (
	"errors"
	"github.com/mansoorceksport/warehouse-stocking/aggregate"
	"github.com/mansoorceksport/warehouse-stocking/repository/storeinventory"
	memoryStoreInventory "github.com/mansoorceksport/warehouse-stocking/repository/storeinventory/memory"
	"github.com/mansoorceksport/warehouse-stocking/repository/warehouseinventory"
	memoryWarehouseInventory "github.com/mansoorceksport/warehouse-stocking/repository/warehouseinventory/memory"
	"sync"
)

var (
	ErrStockNotAvailable      = errors.New("stock not available")
	ErrFailedToProcessRequest = errors.New("failed to process request, one or more stock is not available")
)

type StockConfiguration func(stock *Stock) error

type Stock struct {
	storeInventory     storeinventory.Repository
	warehouseInventory warehouseinventory.Repository
	sync.Mutex
}

func NewStock(configuration ...StockConfiguration) (*Stock, error) {
	stock := &Stock{}
	for _, sc := range configuration {
		err := sc(stock)
		if err != nil {
			return nil, err
		}
	}

	return stock, nil
}

func WithMemoryStoreInventory() StockConfiguration {
	return func(stock *Stock) error {
		stock.storeInventory = memoryStoreInventory.NewMemoryStockInventory()
		return nil
	}
}

func WithMemoryWarehouseInventory() StockConfiguration {
	return func(stock *Stock) error {
		stock.warehouseInventory = memoryWarehouseInventory.NewMemoryWarehouseInventory()
		return nil
	}
}

func (s *Stock) Order(requestProducts []aggregate.Product) error {
	s.Lock()
	defer s.Unlock()
	var processedProduct []aggregate.Product
	for _, requestProduct := range requestProducts {
		warehouseProduct, err := s.warehouseInventory.GetByID(requestProduct.GetID())
		if err != nil {
			return err
		}

		q := warehouseProduct.GetQuantity() - requestProduct.GetQuantity()
		if q < 0 {
			processedProduct = nil
			return ErrStockNotAvailable
		}
		warehouseProduct.SetQuantity(q)
		processedProduct = append(processedProduct, warehouseProduct)
	}

	if processedProduct == nil {
		return ErrFailedToProcessRequest
	}

	for _, pp := range processedProduct {
		err := s.warehouseInventory.Update(pp)
		if err != nil {
			return err
		}
	}

	return nil
}
