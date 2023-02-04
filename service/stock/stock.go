package stock

import (
	"github.com/mansoorceksport/warehouse-stocking/aggregate"
	"github.com/mansoorceksport/warehouse-stocking/service/warehouse"
)

//TODO what is the need of stock service? the store service can call the warehouse service directly to request stock.
type Stock struct {
	*warehouse.Warehouse
}

func NewStock(wh *warehouse.Warehouse) (*Stock, error) {
	stock := &Stock{
		Warehouse: wh,
	}

	return stock, nil
}

func (s *Stock) Request(requestProducts []aggregate.Product) error {
	err := s.Warehouse.ProcessStock(requestProducts)
	if err != nil {
		return err
	}

	return nil
}
