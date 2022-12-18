package equities

import (
	"fmt"
)

// ProcessOrder
func ProcessOrder(es *OrderRequest) (bool, error) {

	fmt.Println("AT@Streaming: GetUserid, GetOrderid:", es.GetUserid(), es.GetOrderid())

	return true, nil
}
