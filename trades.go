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

func parseTradesResponse(jsonContent []byte) ([]Trade, error) {
	response := make([]Trade, 0)
	err := json.Unmarshal(jsonContent, &response)

	return response, err
}

func (sdk *Sdk) Trades(query *tradesQuery) ([]Trade, error) {
	url := "/api/v1/trades" + "?symbol=" + query.symbol + "&limit=" + strconv.Itoa(query.limit)

	responseContent, err := sdk.client.Get(url)
	if err != nil {
		return nil, err
	}

	return parseTradesResponse(responseContent)
}
