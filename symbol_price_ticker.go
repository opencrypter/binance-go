package binance

import (
	"encoding/json"
)

// Symbol price ticker data transfer object (DTO)
type SymbolPrice struct {
	Symbol string
	Price  float64 `json:"price,string"`
}

type symbolPriceTickerQuery struct {
	symbol string
}

// Required query for SymbolPriceTicker.
func NewSymbolPriceTickerQuery(symbol string) *symbolPriceTickerQuery {
	return &symbolPriceTickerQuery{symbol: symbol}
}

// Latest price for a symbol.
func (sdk Sdk) SymbolPriceTicker(query *symbolPriceTickerQuery) (*SymbolPrice, error) {
	request := newRequest("GET", "/api/v3/ticker/price").Param("symbol", query.symbol)
	response, err := sdk.client.Do(request)
	if err != nil {
		return nil, err
	}

	return parseSymbolPriceTickerResponse(response)
}

// Latest price for all symbols.
func (sdk Sdk) AllSymbolPriceTickers() ([]SymbolPrice, error) {
	request := newRequest("GET", "/api/v3/ticker/price")
	response, err := sdk.client.Do(request)
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
