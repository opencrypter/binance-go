package binance

import (
	"encoding/json"
)

type CancelledOrder struct {
	Symbol            string `json:"symbol"`
	OrigClientOrderId string `json:"origClientOrderId"`
	OrderId           int64  `json:"orderId"`
	ClientOrderId     string `json:"clientOrderId"`
}

type cancelOrderRequest struct {
	symbol            *string
	orderId           *int64
	origClientOrderId *string
	newClientOrderId  *string
	recvWindow        *int64
	timestamp         *int64
}

func NewCancelOrderRequest(symbol string) *cancelOrderRequest {
	return &cancelOrderRequest{
		symbol: &symbol,
	}
}

func (r *cancelOrderRequest) OrderId(value int64) *cancelOrderRequest {
	r.orderId = &value
	return r
}

func (r *cancelOrderRequest) OrigClientOrderId(value string) *cancelOrderRequest {
	r.origClientOrderId = &value
	return r
}

func (r *cancelOrderRequest) RecvWindow(value int64) *cancelOrderRequest {
	r.recvWindow = &value
	return r
}

func (sdk Sdk) CancelOrder(request *cancelOrderRequest) (*CancelledOrder, error) {
	req := newRequest("DELETE", "/api/v3/order").
		StringParam("symbol", request.symbol).
		Int64Param("orderId", request.orderId).
		StringParam("origClientOrderId", request.origClientOrderId).
		Int64Param("recvWindow", request.recvWindow).
		Int64Param("timestamp", sdk.clock.Now()).
		Sign()

	responseContent, err := sdk.client.Do(req)
	if err != nil {
		return nil, err
	}

	return parseCancelledOrderResponse(responseContent)
}

func parseCancelledOrderResponse(jsonContent []byte) (*CancelledOrder, error) {
	response := &CancelledOrder{}
	err := json.Unmarshal(jsonContent, &response)
	return response, err
}
