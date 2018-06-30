package binance

import (
	"errors"
	"github.com/isd4n/binance-go-sdk/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestSdk_SymbolPriceTicker(t *testing.T) {
	t.Run("It should convert api response to a SymbolPrice", func(t *testing.T) {
		clientMock := &mocks.Client{}
		sdk := Sdk{client: clientMock}

		expected := validSymbolPriceTickerJson()

		clientMock.On("Get", mock.MatchedBy(func(path string) bool {
			return path == "/api/v3/ticker/price?symbol=LTCBTC"
		})).Return(expected, nil)

		response, _ := sdk.SymbolPriceTicker(NewSymbolPriceTickerQuery("LTCBTC"))

		assert.Equal(t, validSymbolPriceTickerResponse(), response)
	})

	t.Run("It should return error when api fails", func(t *testing.T) {
		clientMock := &mocks.Client{}
		sdk := Sdk{client: clientMock}

		expectedError := errors.New("error")

		clientMock.On("Get", mock.MatchedBy(func(path string) bool {
			return path == "/api/v3/ticker/price?symbol=LTCBTC"
		})).Return(nil, expectedError)

		_, err := sdk.SymbolPriceTicker(NewSymbolPriceTickerQuery("LTCBTC"))

		assert.Equal(t, expectedError, err)
	})

	t.Run("It should return error when response cannot be mapped", func(t *testing.T) {
		clientMock := &mocks.Client{}
		sdk := Sdk{client: clientMock}

		clientMock.On("Get", mock.MatchedBy(func(path string) bool {
			return path == "/api/v3/ticker/price?symbol=LTCBTC"
		})).Return(invalidJson(), nil)

		_, err := sdk.SymbolPriceTicker(NewSymbolPriceTickerQuery("LTCBTC"))

		assert.Error(t, err)
	})
}

func TestSdk_AllSymbolPriceTickers(t *testing.T) {
	t.Run("It should convert api response to a SymbolPrice slice", func(t *testing.T) {
		clientMock := &mocks.Client{}
		sdk := Sdk{client: clientMock}

		expected := validSymbolPriceTickerSliceJson()

		clientMock.On("Get", mock.MatchedBy(func(path string) bool {
			return path == "/api/v3/ticker/price"
		})).Return(expected, nil)

		response, _ := sdk.AllSymbolPriceTickers()

		assert.Equal(t, validSymbolPriceTickerSliceResponse(), response)
	})

	t.Run("It should return error when api fails", func(t *testing.T) {
		clientMock := &mocks.Client{}
		sdk := Sdk{client: clientMock}

		expectedError := errors.New("error")

		clientMock.On("Get", mock.MatchedBy(func(path string) bool {
			return path == "/api/v3/ticker/price"
		})).Return(nil, expectedError)

		_, err := sdk.AllSymbolPriceTickers()

		assert.Equal(t, expectedError, err)
	})

	t.Run("It should return error when response cannot be mapped", func(t *testing.T) {
		clientMock := &mocks.Client{}
		sdk := Sdk{client: clientMock}

		clientMock.On("Get", mock.MatchedBy(func(path string) bool {
			return path == "/api/v3/ticker/price"
		})).Return(invalidJson(), nil)

		_, err := sdk.AllSymbolPriceTickers()

		assert.Error(t, err)
	})
}

func validSymbolPriceTickerJson() []byte {
	return []byte(`{
  		"symbol": "LTCBTC",
  		"price": "4.00000200"
	}`)
}

func validSymbolPriceTickerResponse() *SymbolPrice {
	return &SymbolPrice{
		Symbol: "LTCBTC",
		Price:  4.000002,
	}
}

func validSymbolPriceTickerSliceJson() []byte {
	return []byte(`[
  		{
    		"symbol": "LTCBTC",
    		"price": "4.00000200"
  		},
  		{
    		"symbol": "ETHBTC",
    		"price": "0.07946600"
  		}
	]`)
}

func validSymbolPriceTickerSliceResponse() []SymbolPrice {
	return []SymbolPrice{
		{
			Symbol: "LTCBTC",
			Price:  4.000002,
		},
		{
			Symbol: "ETHBTC",
			Price:  0.079466,
		},
	}
}
