package binance

import (
	"errors"
	"github.com/isd4n/binance-go/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestSdk_ExchangeInfo(t *testing.T) {
	t.Run("It should convert api response to an ExchangeInfo", func(t *testing.T) {
		clientMock := &mocks.Client{}
		sdk := Sdk{client: clientMock}

		expected := validExchangeInfoJson()

		clientMock.On("Get", mock.MatchedBy(func(path string) bool {
			return path == "/api/v1/exchangeInfo"
		})).Return(expected, nil)

		response, _ := sdk.ExchangeInfo()
		assert.Equal(t, validExchangeInfoResponse(), response)
	})

	t.Run("It should return error when api fails", func(t *testing.T) {
		clientMock := &mocks.Client{}
		sdk := Sdk{client: clientMock}

		expectedError := errors.New("error")

		clientMock.On("Get", mock.MatchedBy(func(path string) bool {
			return path == "/api/v1/exchangeInfo"
		})).Return(nil, expectedError)

		_, err := sdk.ExchangeInfo()

		assert.Equal(t, expectedError, err)
	})

	t.Run("It should return error when response cannot be mapped", func(t *testing.T) {
		clientMock := &mocks.Client{}
		sdk := Sdk{client: clientMock}

		clientMock.On("Get", mock.MatchedBy(func(path string) bool {
			return path == "/api/v1/exchangeInfo"
		})).Return(invalidJson(), nil)

		_, err := sdk.ExchangeInfo()

		assert.Error(t, err)
	})
}

func validExchangeInfoJson() []byte {
	return []byte(`{
		"timezone": "UTC",
  		"serverTime": 1508631584636,
  		"rateLimits": [{
      		"rateLimitType": "REQUESTS",
      		"interval": "MINUTE",
      		"limit": 1200
    	},
    	{
      		"rateLimitType": "ORDERS",
			"interval": "SECOND",
      		"limit": 10
		},
    	{
      		"rateLimitType": "ORDERS",
      		"interval": "DAY",
      		"limit": 100000
		}],
  		"exchangeFilters": [],
  		"symbols": [{
    		"symbol": "ETHBTC",
    		"status": "TRADING",
    		"baseAsset": "ETH",
    		"baseAssetPrecision": 8,
    		"quoteAsset": "BTC",
    		"quotePrecision": 8,
    		"orderTypes": ["LIMIT", "MARKET"],
    		"icebergAllowed": false,
    		"filters": [{
      			"filterType": "PRICE_FILTER",
      			"minPrice": "0.00000100",
      			"maxPrice": "100000.00000000",
      			"tickSize": "0.00000100"
    		}, {
      			"filterType": "LOT_SIZE",
      			"minQty": "0.00100000",
      			"maxQty": "100000.00000000",
      			"stepSize": "0.00100000"
    		}, {
      			"filterType": "MIN_NOTIONAL",
      			"minNotional": "0.00100000"
    		}]
  		}]
	}`)
}

func validExchangeInfoResponse() *ExchangeInfo {
	return &ExchangeInfo{
		Timezone:   "UTC",
		ServerTime: 1508631584636,
		RateLimits: []RateLimits{
			{
				RateLimitType: "REQUESTS",
				Interval:      "MINUTE",
				Limit:         1200,
			},
			{
				RateLimitType: "ORDERS",
				Interval:      "SECOND",
				Limit:         10,
			},
			{
				RateLimitType: "ORDERS",
				Interval:      "DAY",
				Limit:         100000,
			},
		},
		ExchangeFilters: []string{},
		Symbols: []Symbol{
			{
				Symbol:             "ETHBTC",
				Status:             "TRADING",
				BaseAsset:          "ETH",
				BaseAssetPrecision: 8,
				QuoteAsset:         "BTC",
				QuotePrecision:     8,
				OrderTypes: []string{
					"LIMIT",
					"MARKET",
				},
				IcebergAllowed: false,
				Filters: []Filter{
					{
						FilterType: "PRICE_FILTER",
						MinPrice:   0.000001,
						MaxPrice:   100000,
						TickSize:   0.000001,
					},
					{
						FilterType:  "LOT_SIZE",
						MinQuantity: 0.001,
						MaxQuantity: 100000,
						StepSize:    0.001,
					},
					{
						FilterType:  "MIN_NOTIONAL",
						MinNotional: 0.001,
					},
				},
			},
		},
	}
}
