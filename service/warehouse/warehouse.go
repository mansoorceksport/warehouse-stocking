package warehouse

import (
	"github.com/mansoorceksport/warehouse-stocking/aggregate"
	"github.com/mansoorceksport/warehouse-stocking/repository/warehouseinventory"
	warehouseInventory "github.com/mansoorceksport/warehouse-stocking/repository/warehouseinventory/memory"
	"sync"
)

type Configuration func(wh *Warehouse) error

type Warehouse struct {
	warehouseInventory warehouseinventory.Repository
	sync.Mutex
}

func NewWarehouse(configuration ...Configuration) *Warehouse {
	wh := &Warehouse{}
	for _, c := range configuration {
		err := c(wh)
		if err != nil {
			return nil
		}
	}
	return wh
}

func WithMemoryWarehouse(products []aggregate.Product) Configuration {
	return func(wh *Warehouse) error {
		wh.warehouseInventory = warehouseInventory.NewMemoryWarehouseInventory()
		for _, product := range products {
			err := wh.warehouseInventory.Add(product)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

func (w *Warehouse) ProcessStock(requestProducts []aggregate.Product) error {
	w.Lock()
	defer w.Unlock()

	var processedProduct []aggregate.Product
	for _, requestProduct := range requestProducts {
		warehouseProduct, err := w.warehouseInventory.GetByID(requestProduct.GetID())
		if err != nil {
			return err
		}

		err = warehouseProduct.ReduceQuantity(requestProduct.GetQuantity())
		if err != nil {
			return err
		}
		processedProduct = append(processedProduct, warehouseProduct)
	}

	for _, pp := range processedProduct {
		err := w.warehouseInventory.Update(pp)
		if err != nil {
			return err
		}
	}

	return nil
}
