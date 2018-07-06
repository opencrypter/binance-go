package binance

import (
	"encoding/json"
	"time"
)

type ExchangeInfo struct {
	Timezone        string
	ServerTime      time.Duration
	RateLimits      []RateLimits
	ExchangeFilters []string
	Symbols         []Symbol
}

type RateLimits struct {
	RateLimitType string
	Interval      string
	Limit         int
}

type Symbol struct {
	Symbol             string
	Status             string
	BaseAsset          string
	BaseAssetPrecision int
	QuoteAsset         string
	QuotePrecision     int
	OrderTypes         []string
	IcebergAllowed     bool
	Filters            []Filter
}

type Filter struct {
	FilterType  string
	MinPrice    float64 `json:",omitempty,string"`
	MaxPrice    float64 `json:",omitempty,string"`
	TickSize    float64 `json:",omitempty,string"`
	MinQuantity float64 `json:"minQty,omitempty,string"`
	MaxQuantity float64 `json:"maxQty,omitempty,string"`
	StepSize    float64 `json:",omitempty,string"`
	MinNotional float64 `json:",omitempty,string"`
}

func parseExchangeInfo(jsonContent []byte) (*ExchangeInfo, error) {
	exchange := &ExchangeInfo{}
	err := json.Unmarshal(jsonContent, &exchange)

	return exchange, err
}

func (sdk Sdk) ExchangeInfo() (*ExchangeInfo, error) {
	request := newRequest("GET", "/api/v1/exchangeInfo")
	response, err := sdk.client.Do(request)
	if err != nil {
		return nil, err
	}

	exchangeInfo, err := parseExchangeInfo(response)
	if err != nil {
		return nil, err
	}

	return exchangeInfo, nil
}
