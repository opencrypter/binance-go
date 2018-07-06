package binance

import (
	"encoding/json"
	"strconv"
)

type CompressedTrade struct {
	Id           int     `json:"a"`
	Price        float64 `json:"p,string"`
	Quantity     float64 `json:"q,string"`
	FirstTradeId int     `json:"f"`
	LastTradeId  int     `json:"l"`
	Time         int     `json:"T"`
	IsBuyerMaker bool    `json:"m"`
	IsBestMatch  bool    `json:"M"`
}

type compressedTradesQuery struct {
	symbol    string
	limit     int
	fromId    int
	startTime int
	endTime   int
}

// Returns the required query for the Trades endpoint.
func NewCompressedTradesQuery(symbol string) *compressedTradesQuery {
	return &compressedTradesQuery{symbol: symbol}
}

// Sets the optional limit parameter that by default is 500.
func (t *compressedTradesQuery) Limit(limit int) *compressedTradesQuery {
	t.limit = limit
	return t
}

// TradeId to fetch from. Default gets most recent trades.
func (t *compressedTradesQuery) FromId(fromId int) *compressedTradesQuery {
	t.fromId = fromId
	return t
}

// Timestamp in ms to get aggregate trades from INCLUSIVE.
func (t *compressedTradesQuery) StartTime(startTime int) *compressedTradesQuery {
	t.startTime = startTime
	return t
}

// Timestamp in ms to get aggregate trades until INCLUSIVE.
func (t *compressedTradesQuery) EndTime(endTime int) *compressedTradesQuery {
	t.endTime = endTime
	return t
}

func parseCompressedTradesResponse(jsonContent []byte) ([]CompressedTrade, error) {
	response := make([]CompressedTrade, 0)
	err := json.Unmarshal(jsonContent, &response)
	return response, err
}

func (sdk *Sdk) CompressedTrades(query *compressedTradesQuery) ([]CompressedTrade, error) {
	request := newRequest("GET", "/api/v1/aggTrades").Param("symbol", query.symbol)

	if query.limit > 0 {
		request.Param("limit", strconv.Itoa(query.limit))
	}

	if query.fromId > 0 {
		request.Param("fromId", strconv.Itoa(query.fromId))
	}

	if query.startTime > 0 {
		request.Param("startTime", strconv.Itoa(query.startTime))
	}

	if query.endTime > 0 {
		request.Param("endTime", strconv.Itoa(query.endTime))
	}

	responseContent, err := sdk.client.Do(request)
	if err != nil {
		return nil, err
	}

	return parseCompressedTradesResponse(responseContent)
}
