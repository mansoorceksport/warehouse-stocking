package warehouse

import (
	"github.com/mansoorceksport/warehouse-stocking/aggregate"
	"github.com/mansoorceksport/warehouse-stocking/repository/warehouseinventory"
	"sync"
)

type Warehouse struct {
	warehouseInventory warehouseinventory.Repository
	sync.Mutex
}

func NewWarehouse(repository warehouseinventory.Repository) *Warehouse {
	return &Warehouse{
		warehouseInventory: repository,
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
