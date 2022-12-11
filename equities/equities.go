package equities

import (
	"fmt"
)

// ProcessOrder
func ProcessOrder(es *OrderRequest) error {
	fmt.Println("ProcessOrder GetOrderid, GetUserid:", es.GetOrderid(), es.GetUserid())
	return nil
}
