package warehouse

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/mansoorceksport/warehouse-stocking/aggregate"
	"testing"
)

var warehouseId uuid.UUID
var warehouse *Warehouse
var products []aggregate.Product
var orders []aggregate.Product

func TestMain(m *testing.M) {
	apple, _ := aggregate.NewProduct("apple", 100, 0.99)
	orange, _ := aggregate.NewProduct("orange", 100, 0.99)
	grapes, _ := aggregate.NewProduct("grapes", 10, 0.99)
	products = append(products, apple, orange, grapes)
	var err error
	warehouse, err = NewWarehouse(
		WithMemoryWarehouse(),
		WithMemoryWarehouseInventory(products),
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
	err = warehouse.warehouseRepository.Add(w)
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
	err := warehouse.ProcessStockRequest(orders)
	if err != nil {
		t.Fatal(err)
	}
	printWarehouseStock()
}

func printWarehouseStock() {
	warehouseInventoryProducts := warehouse.warehouseInventoryRepository.GetAll()
	w := warehouse.warehouseRepository.GetById(warehouseId)
	fmt.Printf("|========%s========|\n", w.GetName())
	for _, p := range warehouseInventoryProducts {
		fmt.Printf("Warehouse: %s quantity is %d \n", p.GetName(), p.GetQuantity())
	}
}
