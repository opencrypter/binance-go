package binance

import (
	"errors"
	"github.com/isd4n/binance-go/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestSdk_SymbolOrderBookTicker(t *testing.T) {
	t.Run("It should convert api response to a SymbolPrice", func(t *testing.T) {
		clientMock := &mocks.Client{}
		sdk := Sdk{client: clientMock}

		expected := valiSymbolOrderBookTickerJson()

		clientMock.On("Get", mock.MatchedBy(func(path string) bool {
			return path == "/api/v3/ticker/bookTicker?symbol=LTCBTC"
		})).Return(expected, nil)

		response, _ := sdk.SymbolOrderBookTicker(NewSymbolOrderBookTickerQuery("LTCBTC"))

		assert.Equal(t, validSymbolOrderBookTickerResponse(), response)
	})

	t.Run("It should return error when api fails", func(t *testing.T) {
		clientMock := &mocks.Client{}
		sdk := Sdk{client: clientMock}

		expectedError := errors.New("error")

		clientMock.On("Get", mock.MatchedBy(func(path string) bool {
			return path == "/api/v3/ticker/bookTicker?symbol=LTCBTC"
		})).Return(nil, expectedError)

		_, err := sdk.SymbolOrderBookTicker(NewSymbolOrderBookTickerQuery("LTCBTC"))

		assert.Equal(t, expectedError, err)
	})

	t.Run("It should return error when response cannot be mapped", func(t *testing.T) {
		clientMock := &mocks.Client{}
		sdk := Sdk{client: clientMock}

		clientMock.On("Get", mock.MatchedBy(func(path string) bool {
			return path == "/api/v3/ticker/bookTicker?symbol=LTCBTC"
		})).Return(invalidJson(), nil)

		_, err := sdk.SymbolOrderBookTicker(NewSymbolOrderBookTickerQuery("LTCBTC"))

		assert.Error(t, err)
	})
}

func TestSdk_AllSymbolOrderBookTickers(t *testing.T) {
	t.Run("It should convert api response to a SymbolPrice slice", func(t *testing.T) {
		clientMock := &mocks.Client{}
		sdk := Sdk{client: clientMock}

		expected := validAllSymbolOrderBookTickersJson()

		clientMock.On("Get", mock.MatchedBy(func(path string) bool {
			return path == "/api/v3/ticker/bookTicker"
		})).Return(expected, nil)

		response, _ := sdk.AllSymbolOrderBookTickers()

		assert.Equal(t, validAllSymbolOrderBookTickersResponse(), response)
	})

	t.Run("It should return error when api fails", func(t *testing.T) {
		clientMock := &mocks.Client{}
		sdk := Sdk{client: clientMock}

		expectedError := errors.New("error")

		clientMock.On("Get", mock.MatchedBy(func(path string) bool {
			return path == "/api/v3/ticker/bookTicker"
		})).Return(nil, expectedError)

		_, err := sdk.AllSymbolOrderBookTickers()

		assert.Equal(t, expectedError, err)
	})

	t.Run("It should return error when response cannot be mapped", func(t *testing.T) {
		clientMock := &mocks.Client{}
		sdk := Sdk{client: clientMock}

		clientMock.On("Get", mock.MatchedBy(func(path string) bool {
			return path == "/api/v3/ticker/bookTicker"
		})).Return(invalidJson(), nil)

		_, err := sdk.AllSymbolOrderBookTickers()

		assert.Error(t, err)
	})
}

func valiSymbolOrderBookTickerJson() []byte {
	return []byte(`{
  		"symbol": "LTCBTC",
  		"bidPrice": "4.00000000",
  		"bidQty": "431.00000000",
  		"askPrice": "4.00000200",
  		"askQty": "9.00000000"
	}`)
}

func validSymbolOrderBookTickerResponse() *OrderBookTicker {
	return &OrderBookTicker{
		Symbol:      "LTCBTC",
		BidPrice:    4,
		BidQuantity: 431,
		AskPrice:    4.000002,
		AskQuantity: 9,
	}
}

func validAllSymbolOrderBookTickersJson() []byte {
	return []byte(`[
  		{
    		"symbol": "LTCBTC",
    		"bidPrice": "4.00000000",
    		"bidQty": "431.00000000",
    		"askPrice": "4.00000200",
    		"askQty": "9.00000000"
  		},
  		{
    		"symbol": "ETHBTC",
    		"bidPrice": "0.07946700",
    		"bidQty": "9.00000000",
    		"askPrice": "100000.00000000",
    		"askQty": "1000.00000000"
  		}
	]`)
}

func validAllSymbolOrderBookTickersResponse() []OrderBookTicker {
	return []OrderBookTicker{
		{
			Symbol:      "LTCBTC",
			BidPrice:    4,
			BidQuantity: 431,
			AskPrice:    4.000002,
			AskQuantity: 9,
		},
		{
			Symbol:      "ETHBTC",
			BidPrice:    0.079467,
			BidQuantity: 9,
			AskPrice:    100000,
			AskQuantity: 1000,
		},
	}
}
