package binance

import (
	"encoding/json"
)

type getOpenOrdersQuery struct {
	symbol     *string
	recvWindow *int64
	timestamp  *int64
}

func NewGetOpenOrdersQuery(symbol string) *getOpenOrdersQuery {
	return &getOpenOrdersQuery{
		symbol: &symbol,
	}
}

func (r *getOpenOrdersQuery) RecvWindow(value int64) *getOpenOrdersQuery {
	r.recvWindow = &value
	return r
}

func (sdk Sdk) GetOpenOrders(query *getOpenOrdersQuery) ([]Order, error) {
	req := newRequest("GET", "/api/v3/openOrders").
		StringParam("symbol", query.symbol).
		Int64Param("recvWindow", query.recvWindow).
		Int64Param("timestamp", sdk.clock.Now()).
		Sign()

	responseContent, err := sdk.client.Do(req)
	if err != nil {
		return nil, err
	}

	return parseOrdersResponse(responseContent)
}

func parseOrdersResponse(jsonContent []byte) ([]Order, error) {
	response := make([]Order, 0)
	err := json.Unmarshal(jsonContent, &response)
	return response, err
}
