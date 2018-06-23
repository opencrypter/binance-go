package binance

import (
	"encoding/json"
	"strconv"
)

type depthResponse struct {
	LastUpdateId int
	Bids         [][]interface{}
	Asks         [][]interface{}
}
type Depth struct {
	LastUpdateId int
	Bids         []DepthOrder
	Asks         []DepthOrder
}

type DepthOrder struct {
	Price    float64
	Quantity float64
}

func parseDepthResponse(jsonContent []byte) (*depthResponse, error) {
	response := &depthResponse{}
	err := json.Unmarshal(jsonContent, &response)

	return response, err
}

type depthQuery struct {
	symbol string
	limit  int
}

func convertDepthOrders(sliceOfDepthOrders [][]interface{}) []DepthOrder {
	depthOrders := make([]DepthOrder, 0)
	for _, bid := range sliceOfDepthOrders {
		price, _ := strconv.ParseFloat(bid[0].(string), 64)
		quantity, _ := strconv.ParseFloat(bid[1].(string), 64)

		depthOrders = append(depthOrders, DepthOrder{Price: price, Quantity: quantity})
	}

	return depthOrders
}

// Returns the required query for the Depth endpoint.
func NewDepthQuery(symbol string) *depthQuery {
	return &depthQuery{
		symbol: symbol,
		limit:  100,
	}
}

// Sets the optional limit parameter that by default is 100.
func (d *depthQuery) Limit(limit int) *depthQuery {
	d.limit = limit

	return d
}

// Market depth for a specific symbol. An example of a symbol is "ETHBTC"
//
// The query is optional and it specifies the depth limit.
// Accepted values for the limit: [5, 10, 20, 50, 100 (default), 500, 1000]
//
// Caution: setting limit=0 can return a lot of data.
func (sdk Sdk) Depth(query *depthQuery) (*Depth, error) {
	url := "/api/v1/depth" + "?symbol=" + query.symbol + "&limit=" + strconv.Itoa(query.limit)

	responseContent, err := sdk.client.Get(url)
	if err != nil {
		return nil, err
	}

	response, err := parseDepthResponse(responseContent)
	if err != nil {
		return nil, err
	}

	return &Depth{
		LastUpdateId: response.LastUpdateId,
		Bids:         convertDepthOrders(response.Bids),
		Asks:         convertDepthOrders(response.Asks),
	}, nil
}
