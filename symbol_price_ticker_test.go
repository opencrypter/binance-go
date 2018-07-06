package binance

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSdk_SymbolPriceTicker(t *testing.T) {
	method, url := "GET", "/api/v3/ticker/price"

	t.Run("It should convert api response to a SymbolPrice", func(t *testing.T) {
		mockedClient := NewMockClient(gomock.NewController(t))
		sdk := Sdk{client: mockedClient}

		expectedRequest := newRequest(method, url).Param("symbol", "LTCBTC")

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(validSymbolPriceTickerJson(), nil)

		response, _ := sdk.SymbolPriceTicker(NewSymbolPriceTickerQuery("LTCBTC"))

		assert.Equal(t, validSymbolPriceTickerResponse(), response)
	})

	t.Run("It should return error when api fails", func(t *testing.T) {
		mockedClient := NewMockClient(gomock.NewController(t))
		sdk := Sdk{client: mockedClient}

		expectedRequest := newRequest(method, url).Param("symbol", "LTCBTC")

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(nil, errors.New("error"))

		_, err := sdk.SymbolPriceTicker(NewSymbolPriceTickerQuery("LTCBTC"))

		assert.Error(t, err)
	})

	t.Run("It should return error when response cannot be mapped", func(t *testing.T) {
		mockedClient := NewMockClient(gomock.NewController(t))
		sdk := Sdk{client: mockedClient}

		expectedRequest := newRequest(method, url).Param("symbol", "LTCBTC")

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(invalidJson(), nil)

		_, err := sdk.SymbolPriceTicker(NewSymbolPriceTickerQuery("LTCBTC"))

		assert.Error(t, err)
	})
}

func TestSdk_AllSymbolPriceTickers(t *testing.T) {
	method, url := "GET", "/api/v3/ticker/price"

	t.Run("It should convert api response to a SymbolPrice slice", func(t *testing.T) {
		mockedClient := NewMockClient(gomock.NewController(t))
		sdk := Sdk{client: mockedClient}

		expectedRequest := newRequest(method, url)

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(validSymbolPriceTickerSliceJson(), nil)

		response, _ := sdk.AllSymbolPriceTickers()

		assert.Equal(t, validSymbolPriceTickerSliceResponse(), response)
	})

	t.Run("It should return error when api fails", func(t *testing.T) {
		mockedClient := NewMockClient(gomock.NewController(t))
		sdk := Sdk{client: mockedClient}

		expectedRequest := newRequest(method, url)

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(nil, errors.New("error"))

		_, err := sdk.AllSymbolPriceTickers()

		assert.Error(t, err)
	})

	t.Run("It should return error when response cannot be mapped", func(t *testing.T) {
		mockedClient := NewMockClient(gomock.NewController(t))
		sdk := Sdk{client: mockedClient}

		expectedRequest := newRequest(method, url)

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(invalidJson(), nil)

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
