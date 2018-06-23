package binance

import "strconv"

type historicalTradesQuery struct {
	symbol string
	limit  int
	fromId int
}

func NewHistoricalTradesQuery(symbol string) *historicalTradesQuery {
	return &historicalTradesQuery{
		symbol: symbol,
		limit:  500,
	}
}

func (t *historicalTradesQuery) Limit(limit int) *historicalTradesQuery {
	t.limit = limit

	return t
}

func (t *historicalTradesQuery) FromId(fromId int) *historicalTradesQuery {
	t.fromId = fromId

	return t
}

func (sdk *Sdk) HistoricalTrades(query *historicalTradesQuery) ([]Trade, error) {
	url := "/api/v1/historicalTrades" + "?symbol=" + query.symbol + "&limit=" + strconv.Itoa(query.limit)

	if query.fromId > 0 {
		url += "&fromId=" + strconv.Itoa(query.fromId)
	}

	responseContent, err := sdk.client.Get(url)
	if err != nil {
		return nil, err
	}

	return parseTradesResponse(responseContent)
}
