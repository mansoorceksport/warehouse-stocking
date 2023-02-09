package warehouse

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/mansoorceksport/warehouse-stocking/aggregate"
	"github.com/mansoorceksport/warehouse-stocking/common/postgres"
	"log"
	"testing"
)

var warehouseId uuid.UUID
var warehouse *Warehouse
var products []aggregate.Product
var orders []aggregate.Product
var ctx context.Context
var pq *postgres.Postgres

func TestMain(m *testing.M) {
	ctx = context.Background()
	apple, _ := aggregate.NewProduct("apple", 100, 0.99)
	orange, _ := aggregate.NewProduct("orange", 100, 0.99)
	grapes, _ := aggregate.NewProduct("grapes", 10, 0.99)
	products = append(products, apple, orange, grapes)
	var err error
	pq = postgres.NewPostgres("postgres://postgres:password@localhost:5432/stocking")
	warehouse, err = NewWarehouse(
		//WithMemoryDepot(),
		WithPostgresWarehouse(pq),
		WithMemoryWarehouseInventory(ctx, products),
		WithPostgresDepot(pq),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	w, err := aggregate.NewWareHouse("warehouse one")
	warehouseId = w.GetID()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = warehouse.warehouseRepository.Add(ctx, w)
	if err != nil {
		return
	}

	orderApple, _ := aggregate.NewProduct("apple", 2, 0.99)
	orderApple.SetID(products[0].GetID())
	orderOrange, _ := aggregate.NewProduct("orange", 2, 0.99)
	orderOrange.SetID(products[1].GetID())
	orders = append(orders, orderApple, orderOrange)

	m.Run()
}

func TestWarehouse_ProcessStock(t *testing.T) {
	err := warehouse.ProcessStockRequest(ctx, orders)
	if err != nil {
		t.Fatal(err)
	}
	printWarehouseStock(ctx)
}

func printWarehouseStock(ctx context.Context) {
	warehouseInventoryProducts := warehouse.warehouseInventoryRepository.GetAll(ctx)
	w, err := warehouse.warehouseRepository.GetById(ctx, warehouseId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("|========%s========|\n", w.GetName())
	for _, p := range warehouseInventoryProducts {
		fmt.Printf("Warehouse: %s quantity is %d \n", p.GetName(), p.GetQuantity())
	}
}
