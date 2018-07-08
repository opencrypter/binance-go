package binance

import "encoding/json"

type AccountTrade struct {
	Id              int64   `json:"id"`
	OrderId         int64   `json:"orderId"`
	Price           float64 `json:"price,string"`
	Quantity        float64 `json:"qty,string"`
	Commission      float64 `json:"commission,string"`
	CommissionAsset string  `json:"commissionAsset"`
	Time            int64   `json:"time"`
	IsBuyer         bool    `json:"isBuyer"`
	IsMaker         bool    `json:"isMaker"`
	IsBestMatch     bool    `json:"isBestMatch"`
}

type myTradesQuery struct {
	symbol     *string
	limit      *int64
	fromId     *int64
	recvWindow *int64
}

func (r *myTradesQuery) Limit(value int64) *myTradesQuery {
	r.limit = &value
	return r
}

func (r *myTradesQuery) FromId(value int64) *myTradesQuery {
	r.fromId = &value
	return r
}

func (r *myTradesQuery) RecvWindow(value int64) *myTradesQuery {
	r.recvWindow = &value
	return r
}

func NewMyTradesQuery(symbol string) *myTradesQuery {
	return &myTradesQuery{symbol: &symbol}
}

func (sdk Sdk) MyTrades(query *myTradesQuery) ([]AccountTrade, error) {
	req := newRequest("GET", "/api/v3/myTrades").
		StringParam("symbol", query.symbol).
		Int64Param("recvWindow", query.recvWindow).
		Int64Param("limit", query.limit).
		Int64Param("fromId", query.fromId).
		Int64Param("timestamp", sdk.clock.Now()).
		Sign()

	responseContent, err := sdk.client.Do(req)
	if err != nil {
		return nil, err
	}

	return parseAccountTradesResponse(responseContent)
}

func parseAccountTradesResponse(jsonContent []byte) ([]AccountTrade, error) {
	response := make([]AccountTrade, 0)
	err := json.Unmarshal(jsonContent, &response)
	return response, err
}
