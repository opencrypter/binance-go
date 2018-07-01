package binance

import "encoding/json"

type OrderBookTicker struct {
	Symbol      string  `json:"symbol"`
	BidPrice    float64 `json:"bidPrice,string"`
	BidQuantity float64 `json:"bidQty,string"`
	AskPrice    float64 `json:"askPrice,string"`
	AskQuantity float64 `json:"askQty,string"`
}

type symbolOrderBookTickerQuery struct {
	symbol string
}

func NewSymbolOrderBookTickerQuery(symbol string) *symbolOrderBookTickerQuery {
	return &symbolOrderBookTickerQuery{symbol: symbol}
}

func (sdk Sdk) SymbolOrderBookTicker(query *symbolOrderBookTickerQuery) (*OrderBookTicker, error) {
	url := "/api/v3/ticker/bookTicker?symbol=" + query.symbol

	response, err := sdk.client.Get(url)
	if err != nil {
		return nil, err
	}

	return parseSymbolOrderBookTickerResponse(response)
}

func (sdk Sdk) AllSymbolOrderBookTickers() ([]OrderBookTicker, error) {
	url := "/api/v3/ticker/bookTicker"

	response, err := sdk.client.Get(url)
	if err != nil {
		return nil, err
	}

	return parseAllSymbolOrderBookTickersResponse(response)
}

func parseSymbolOrderBookTickerResponse(jsonContent []byte) (*OrderBookTicker, error) {
	response := &OrderBookTicker{}
	err := json.Unmarshal(jsonContent, &response)
	return response, err
}

func parseAllSymbolOrderBookTickersResponse(jsonContent []byte) ([]OrderBookTicker, error) {
	response := make([]OrderBookTicker, 0)
	err := json.Unmarshal(jsonContent, &response)
	return response, err
}
