package equities

import (
	"fmt"
)

// ProcessOrder
func ProcessOrder(es *OrderRequest) (int64, bool, error) {
	pQty := es.GetQuantity() - int64((float64(es.GetQuantity()) * 0.10))

	fmt.Println("AT@Streaming: GetUserid, GetOrderid, GetQuantity, ProcessedQty:", es.GetUserid(), es.GetOrderid(), es.GetQuantity(), pQty)

	return pQty, true, nil
}
