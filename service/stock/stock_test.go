package stock

import (
	"fmt"
	"github.com/mansoorceksport/warehouse-stocking/aggregate"
	"testing"
)

var stock *Stock
var products []aggregate.Product
var orders []aggregate.Product

func TestMain(t *testing.M) {

	apple, _ := aggregate.NewProduct("apple", 100, 0.99)

	orange, _ := aggregate.NewProduct("orange", 100, 0.99)

	grapes, _ := aggregate.NewProduct("grapes", 10, 0.99)

	products = append(products, apple, orange, grapes)
	s, _ := NewStock(WithMemoryStoreInventory(), WithMemoryWarehouseInventory())
	stock = s
	for _, p := range products {
		_ = stock.warehouseInventory.Add(p)
	}

	orderApple, _ := aggregate.NewProduct("apple", 2, 0.99)

	orderApple.SetID(products[0].GetID())

	orderOrange, _ := aggregate.NewProduct("orange", 2, 0.99)
	orderOrange.SetID(products[1].GetID())

	orders = append(orders, orderApple, orderOrange)
	t.Run()
}

func TestRequest(t *testing.T) {
	err := stock.Request(orders)
	if err != nil {
		t.Fatal(err)
	}
	printWarehouseStock(stock)
	printStoreStock(stock)
}

func printWarehouseStock(s *Stock) {
	warehouseInventoryProducts := s.warehouseInventory.GetAll()
	fmt.Printf("===============\n")
	for _, p := range warehouseInventoryProducts {
		fmt.Printf("Warehouse: %s quantity is %d \n", p.GetName(), p.GetQuantity())
	}
}

func printStoreStock(s *Stock) {
	storeInventoryProducts := s.storeInventory.GetAll()
	fmt.Printf("===============\n")
	for _, p := range storeInventoryProducts {
		fmt.Printf("Store: %s quantity is %d \n", p.GetName(), p.GetQuantity())
	}
}
