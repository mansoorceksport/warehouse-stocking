package store

import (
	"fmt"
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
	wh, err := warehouse.NewWarehouse(
		warehouse.WithMemoryWarehouse(),
		warehouse.WithMemoryWarehouseInventory(products),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	w, err := aggregate.NewWareHouse("warehouse one")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = wh.AddWarehouse(w)
	if err != nil {
		return
	}

	orderApple, _ := aggregate.NewProduct("apple", 2, 0.99)
	orderApple.SetID(products[0].GetID())
	orderOrange, _ := aggregate.NewProduct("orange", 2, 0.99)
	orderOrange.SetID(products[1].GetID())
	orders = append(orders, orderApple, orderOrange)

	store = NewStore(
		WithMemoryStoreRepository(),
		WithMemoryStoreInventoryRepository(),
		WithWarehouseService(wh),
	)
	s, err := aggregate.NewStore("jakarta one")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = store.storeRepository.Add(s)
	if err != nil {
		return
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	m.Run()
}

func TestStore_RequestStock(t *testing.T) {
	err := store.RequestStock(orders)
	if err != nil {
		t.Fatal(err)
	}
	printStoreStock()
}

func printStoreStock() {
	storeInventoryProducts := store.storeInventoryRepository.GetAll()
	//w := store.warehouseService..GetById(warehouseId)
	fmt.Printf("|================|\n")
	for _, p := range storeInventoryProducts {
		fmt.Printf("Store: %s quantity is %d \n", p.GetName(), p.GetQuantity())
	}
}
