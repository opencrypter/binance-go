package binance

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSdk_Trades(t *testing.T) {
	method, url := "GET", "/api/v1/historicalTrades"

	t.Run("It should convert api response to a Trade slice", func(t *testing.T) {
		mockedClient := NewMockClient(gomock.NewController(t))
		sdk := Sdk{client: mockedClient}

		expectedRequest := newRequest(method, url).
			Param("symbol", "ETHBTC").
			Param("limit", "500")

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(validTradesJson(), nil)

		response, _ := sdk.Trades(NewTradesQuery("ETHBTC"))

		assert.Equal(t, validTradesResponse(), response)
	})

	t.Run("It should read optional parameters", func(t *testing.T) {
		mockedClient := NewMockClient(gomock.NewController(t))
		sdk := Sdk{client: mockedClient}

		expectedRequest := newRequest(method, url).
			Param("symbol", "ETHBTC").
			Param("limit", "350").
			Param("fromId", "2300")

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(validTradesJson(), nil)

		query := NewTradesQuery("ETHBTC").Limit(350).FromId(2300)
		response, _ := sdk.Trades(query)

		assert.Equal(t, validTradesResponse(), response)
	})

	t.Run("It should return error when api fails", func(t *testing.T) {
		mockedClient := NewMockClient(gomock.NewController(t))
		sdk := Sdk{client: mockedClient}

		expectedRequest := newRequest(method, url).
			Param("symbol", "ETHBTC").
			Param("limit", "500")

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(validTradesJson(), errors.New("error"))
		_, err := sdk.Trades(NewTradesQuery("ETHBTC"))

		assert.Error(t, err)
	})

	t.Run("It should return error when response cannot be mapped", func(t *testing.T) {
		mockedClient := NewMockClient(gomock.NewController(t))
		sdk := Sdk{client: mockedClient}

		expectedRequest := newRequest(method, url).
			Param("symbol", "ETHBTC").
			Param("limit", "500")

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(invalidJson(), nil)

		_, err := sdk.Trades(NewTradesQuery("ETHBTC"))

		assert.Error(t, err)
	})
}

func validTradesJson() []byte {
	return []byte(`[
 		{
   		"id": 28457,
   		"price": "4.00000100",
   		"qty": "12.00000000",
   		"time": 1499865549590,
   		"isBuyerMaker": true,
   		"isBestMatch": true
 		}
	]`)
}

func validTradesResponse() []Trade {
	return []Trade{{
		Id:           28457,
		Price:        4.000001,
		Quantity:     12.0,
		Time:         1499865549590,
		IsBuyerMaker: true,
		IsBestMatch:  true,
	}}
}
