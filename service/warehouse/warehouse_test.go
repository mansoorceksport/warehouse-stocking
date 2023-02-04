package warehouse

import (
	"fmt"
	"github.com/mansoorceksport/warehouse-stocking/aggregate"
	"testing"
)

var warehouse *Warehouse
var products []aggregate.Product
var orders []aggregate.Product

func TestMain(m *testing.M) {
	apple, _ := aggregate.NewProduct("apple", 100, 0.99)
	orange, _ := aggregate.NewProduct("orange", 100, 0.99)
	grapes, _ := aggregate.NewProduct("grapes", 10, 0.99)
	products = append(products, apple, orange, grapes)
	warehouse = NewWarehouse(WithMemoryWarehouse(products))

	orderApple, _ := aggregate.NewProduct("apple", 2, 0.99)
	orderApple.SetID(products[0].GetID())
	orderOrange, _ := aggregate.NewProduct("orange", 2, 0.99)
	orderOrange.SetID(products[1].GetID())
	orders = append(orders, orderApple, orderOrange)

	m.Run()
}

func TestWarehouse_ProcessStock(t *testing.T) {
	err := warehouse.ProcessStock(orders)
	if err != nil {
		t.Fatal(err)
	}
	printWarehouseStock()
}

func printWarehouseStock() {
	warehouseInventoryProducts := warehouse.warehouseInventory.GetAll()
	fmt.Printf("===============\n")
	for _, p := range warehouseInventoryProducts {
		fmt.Printf("Warehouse: %s quantity is %d \n", p.GetName(), p.GetQuantity())
	}
}
