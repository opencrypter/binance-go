package binance

import (
	"encoding/json"
	"strconv"
)

type Trade struct {
	Id           int
	Price        float64 `json:"price,string"`
	Quantity     float64 `json:"qty,string"`
	Time         int     `json:"time"`
	IsBuyerMaker bool    `json:"isBuyerMaker"`
	IsBestMatch  bool    `json:"isBestMatch"`
}

type tradesQuery struct {
	symbol string
	limit  int
	fromId int
}

// Returns the required query for the Trades endpoint.
func NewTradesQuery(symbol string) *tradesQuery {
	return &tradesQuery{
		symbol: symbol,
		limit:  500,
	}
}

// Sets the optional limit parameter that by default is 500.
func (t *tradesQuery) Limit(limit int) *tradesQuery {
	t.limit = limit
	return t
}

// TradeId to fetch from. Default gets most recent trades.
func (t *tradesQuery) FromId(fromId int) *tradesQuery {
	t.fromId = fromId
	return t
}

func parseTradesResponse(jsonContent []byte) ([]Trade, error) {
	response := make([]Trade, 0)
	err := json.Unmarshal(jsonContent, &response)
	return response, err
}

func (sdk *Sdk) Trades(query *tradesQuery) ([]Trade, error) {
	request := newRequest("GET", "/api/v1/historicalTrades").
		Param("symbol", query.symbol).
		Param("limit", strconv.Itoa(query.limit))

	if query.fromId > 0 {
		request.Param("fromId", strconv.Itoa(query.fromId))
	}

	responseContent, err := sdk.client.Do(request)
	if err != nil {
		return nil, err
	}

	return parseTradesResponse(responseContent)
}
