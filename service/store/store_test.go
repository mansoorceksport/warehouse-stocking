package store

import (
	"github.com/mansoorceksport/warehouse-stocking/aggregate"
	"github.com/mansoorceksport/warehouse-stocking/service/warehouse"
	"testing"
)

var store *Store
var products []aggregate.Product
var orders []aggregate.Product

func TestMain(m *testing.M) {
	apple, _ := aggregate.NewProduct("apple", 100, 0.99)
	orange, _ := aggregate.NewProduct("orange", 100, 0.99)
	grapes, _ := aggregate.NewProduct("grapes", 10, 0.99)
	products = append(products, apple, orange, grapes)
	wh := warehouse.NewWarehouse(
		warehouse.WithMemoryWarehouse(products),
	)

	orderApple, _ := aggregate.NewProduct("apple", 2, 0.99)
	orderApple.SetID(products[0].GetID())
	orderOrange, _ := aggregate.NewProduct("orange", 2, 0.99)
	orderOrange.SetID(products[1].GetID())
	orders = append(orders, orderApple, orderOrange)

	store = NewStore(
		WithMemoryStoreInventory(),
		WithStockService(wh),
	)
	m.Run()
}

func TestStore_RequestStock(t *testing.T) {
	err := store.RequestStock(orders)
	if err != nil {
		t.Fatal(err)
	}
}
