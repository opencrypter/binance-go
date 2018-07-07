package binance

import (
	"encoding/json"
)

type FullOrder struct {
	Symbol           string      `json:"symbol"`
	OrderId          int         `json:"orderId"`
	ClientOrderId    string      `json:"clientOrderId"`
	TransactionTime  int         `json:"transactTime"`
	Price            float64     `json:"price,string"`
	OriginalQuantity float64     `json:"origQty,string"`
	ExecutedQuantity float64     `json:"executedQty,string"`
	Status           string      `json:"status"`
	TimeInForce      string      `json:"timeInForce"`
	Type             string      `json:"type"`
	Side             string      `json:"side"`
	Fills            []OrderFill `json:"fills"`
}

type OrderFill struct {
	Price           float64 `json:"price,string"`
	Quantity        float64 `json:"qty,string"`
	Commission      float64 `json:"commission,string"`
	CommissionAsset string  `json:"commissionAsset"`
}

type newOrderRequest struct {
	symbol               *string
	side                 *string
	orderType            *string
	quantity             *float64
	timestamp            *int64
	timeInForce          *string
	price                *float64
	newClientOrderId     *string
	stopPrice            *float64
	icebergQuantity      *float64
	newOrderResponseType *string
	recvWindow           *int64
}

func NewOrderRequest(symbol string, side string, orderType string, quantity float64) *newOrderRequest {
	responseType := "FULL"

	return &newOrderRequest{
		symbol:               &symbol,
		side:                 &side,
		orderType:            &orderType,
		quantity:             &quantity,
		newOrderResponseType: &responseType,
	}
}

func (r *newOrderRequest) TimeInForce(value string) *newOrderRequest {
	r.timeInForce = &value
	return r
}

func (r *newOrderRequest) Price(value float64) *newOrderRequest {
	r.price = &value
	return r
}

func (r *newOrderRequest) NewClientOrderId(value string) *newOrderRequest {
	r.newClientOrderId = &value
	return r
}

func (r *newOrderRequest) StopPrice(value float64) *newOrderRequest {
	r.stopPrice = &value
	return r
}

func (r *newOrderRequest) IcebergQuantity(value float64) *newOrderRequest {
	r.icebergQuantity = &value
	return r
}

func (r *newOrderRequest) RecvWindow(value int64) *newOrderRequest {
	r.recvWindow = &value
	return r
}

func (r *newOrderRequest) NewOrderResponseType(value string) *newOrderRequest {
	r.newOrderResponseType = &value
	return r
}

func (sdk Sdk) NewOrder(request *newOrderRequest) (*FullOrder, error) {
	req := newRequest("POST", "/api/v3/order").
		StringParam("symbol", request.symbol).
		StringParam("type", request.orderType).
		StringParam("side", request.side).
		Float64Param("quantity", request.quantity).
		StringParam("newOrderRespType", request.newOrderResponseType).
		Int64Param("timestamp", sdk.clock.Now()).
		Float64Param("price", request.price).
		Float64Param("icebergQty", request.icebergQuantity).
		StringParam("newClientOrderId", request.newClientOrderId).
		Float64Param("stopPrice", request.stopPrice).
		Int64Param("recvWindow", request.recvWindow).
		StringParam("timeInForce", request.timeInForce).
		Sign()

	responseContent, err := sdk.client.Do(req)
	if err != nil {
		return nil, err
	}

	return parseNewOrderResponse(responseContent)
}

func parseNewOrderResponse(jsonContent []byte) (*FullOrder, error) {
	response := &FullOrder{}
	err := json.Unmarshal(jsonContent, &response)
	return response, err
}
