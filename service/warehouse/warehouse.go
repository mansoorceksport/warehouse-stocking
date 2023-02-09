package warehouse

import (
	"context"
	"github.com/mansoorceksport/warehouse-stocking/aggregate"
	"github.com/mansoorceksport/warehouse-stocking/common/postgres"
	warehouseRepository "github.com/mansoorceksport/warehouse-stocking/repository/depot"
	memoryWarehouseRepository "github.com/mansoorceksport/warehouse-stocking/repository/depot/memory"
	postgresDepot "github.com/mansoorceksport/warehouse-stocking/repository/depot/postgres"
	"github.com/mansoorceksport/warehouse-stocking/repository/warehouseinventory"
	memoryWarehouseInventoryRepository "github.com/mansoorceksport/warehouse-stocking/repository/warehouseinventory/memory"
	postgresWarehouseInventoryRepository "github.com/mansoorceksport/warehouse-stocking/repository/warehouseinventory/postgres"
	"sync"
)

type Configuration func(wh *Warehouse) error

type Warehouse struct {
	warehouseRepository          warehouseRepository.Repository
	warehouseInventoryRepository warehouseinventory.Repository
	sync.Mutex
}

func NewWarehouse(configuration ...Configuration) (*Warehouse, error) {
	wh := &Warehouse{}

	for _, c := range configuration {
		err := c(wh)
		if err != nil {
			return nil, err
		}
	}
	return wh, nil
}

func WithMemoryDepot() Configuration {
	return func(wh *Warehouse) error {
		wh.warehouseRepository = memoryWarehouseRepository.NewMemoryWareHouse()
		return nil
	}
}

func WithPostgresDepot(pg *postgres.Postgres) Configuration {
	return func(wh *Warehouse) error {
		wh.warehouseRepository = postgresDepot.NewPostgresWarehouse(pg)
		return nil
	}
}

func WithMemoryWarehouseInventory(ctx context.Context, products []aggregate.Product) Configuration {
	return func(wh *Warehouse) error {
		wh.warehouseInventoryRepository = memoryWarehouseInventoryRepository.NewMemoryWarehouseInventory()
		for _, product := range products {
			err := wh.warehouseInventoryRepository.Add(ctx, product)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

func WithPostgresWarehouse(pg *postgres.Postgres) Configuration {
	return func(wh *Warehouse) error {
		wh.warehouseInventoryRepository = postgresWarehouseInventoryRepository.New(pg)
		return nil
	}
}

func (w *Warehouse) AddWarehouse(aw aggregate.Warehouse) error {
	ctx := context.Background()
	err := w.warehouseRepository.Add(ctx, aw)
	if err != nil {
		return err
	}
	return nil
}

func (w *Warehouse) ProcessStockRequest(ctx context.Context, requestProducts []aggregate.Product) error {
	w.Lock()
	defer w.Unlock()

	var processedProduct []aggregate.Product
	for _, requestProduct := range requestProducts {
		warehouseProduct, err := w.warehouseInventoryRepository.GetByID(ctx, requestProduct.GetID())
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
		err := w.warehouseInventoryRepository.Update(ctx, pp)
		if err != nil {
			return err
		}
	}

	return nil
}
