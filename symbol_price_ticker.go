package binance

import "encoding/json"

type SymbolPrice struct {
	Symbol string
	Price  float64 `json:"price,string"`
}

type symbolPriceTickerQuery struct {
	symbol string
}

func NewSymbolPriceTickerQuery(symbol string) *symbolPriceTickerQuery {
	return &symbolPriceTickerQuery{symbol: symbol}
}

func (sdk Sdk) SymbolPriceTicker(query *symbolPriceTickerQuery) (*SymbolPrice, error) {
	url := "/api/v3/ticker/price?symbol=" + query.symbol

	response, err := sdk.client.Get(url)
	if err != nil {
		return nil, err
	}

	return parseSymbolPriceTickerResponse(response)
}

func (sdk Sdk) AllSymbolPriceTickers() ([]SymbolPrice, error) {
	url := "/api/v3/ticker/price"

	response, err := sdk.client.Get(url)
	if err != nil {
		return nil, err
	}

	return parseAllSymbolPriceTickersResponse(response)
}

func parseSymbolPriceTickerResponse(jsonContent []byte) (*SymbolPrice, error) {
	response := &SymbolPrice{}
	err := json.Unmarshal(jsonContent, &response)
	return response, err
}

func parseAllSymbolPriceTickersResponse(jsonContent []byte) ([]SymbolPrice, error) {
	response := make([]SymbolPrice, 0)
	err := json.Unmarshal(jsonContent, &response)
	return response, err
}
