package binance

import (
	"encoding/json"
)

type Order struct {
	Symbol           string  `json:"symbol"`
	OrderId          int64   `json:"orderId"`
	ClientOrderId    string  `json:"clientOrderId"`
	Price            float64 `json:"price,string"`
	OriginalQuantity float64 `json:"origQty,string"`
	ExecutedQuantity float64 `json:"executedQty,string"`
	Status           string  `json:"status"`
	TimeInForce      string  `json:"timeInForce"`
	Type             string  `json:"type"`
	Side             string  `json:"side"`
	StopPrice        float64 `json:"stopPrice,string"`
	IcebergQuantity  float64 `json:"icebergQty,string"`
	Time             int64   `json:"time"`
	IsWorking        bool    `json:"isWorking"`
}

type getOrderQuery struct {
	symbol            *string
	orderId           *int64
	origClientOrderId *string
	recvWindow        *int64
	timestamp         *int64
}

func NewGetOrderQuery(symbol string) *getOrderQuery {
	return &getOrderQuery{
		symbol: &symbol,
	}
}

func (r *getOrderQuery) OrderId(value int64) *getOrderQuery {
	r.orderId = &value
	return r
}

func (r *getOrderQuery) OrigClientOrderId(value string) *getOrderQuery {
	r.origClientOrderId = &value
	return r
}

func (r *getOrderQuery) RecvWindow(value int64) *getOrderQuery {
	r.recvWindow = &value
	return r
}

func (sdk Sdk) GetOrder(query *getOrderQuery) (*Order, error) {
	req := newRequest("GET", "/api/v3/order").
		StringParam("symbol", query.symbol).
		Int64Param("orderId", query.orderId).
		StringParam("origClientOrderId", query.origClientOrderId).
		Int64Param("recvWindow", query.recvWindow).
		Int64Param("timestamp", sdk.clock.Now()).
		Sign()

	responseContent, err := sdk.client.Do(req)
	if err != nil {
		return nil, err
	}

	return parseOrderResponse(responseContent)
}

func parseOrderResponse(jsonContent []byte) (*Order, error) {
	response := &Order{}
	err := json.Unmarshal(jsonContent, &response)
	return response, err
}
