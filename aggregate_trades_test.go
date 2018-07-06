package binance

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSdk_CompressedTrades(t *testing.T) {
	method, url := "GET", "/api/v1/aggTrades"

	t.Run("It should convert api response to a compressed Trade slice", func(t *testing.T) {
		mockedClient := NewMockClient(gomock.NewController(t))
		sdk := Sdk{client: mockedClient}
		expectedRequest := newRequest(method, url).Param("symbol", "ETHBTC")

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(validCompressedTradesJson(), nil)

		response, _ := sdk.CompressedTrades(NewCompressedTradesQuery("ETHBTC"))

		assert.Equal(t, validCompressedTradesResponse(), response)
	})

	t.Run("It should read optional parameters", func(t *testing.T) {
		mockedClient := NewMockClient(gomock.NewController(t))
		sdk := Sdk{client: mockedClient}
		expectedRequest := newRequest(method, url).
			Param("symbol", "ETHBTC").
			Param("limit", "10").
			Param("fromId", "1").
			Param("startTime", "1498793709153").
			Param("endTime", "1498793709163")

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(validCompressedTradesJson(), nil)

		query := NewCompressedTradesQuery("ETHBTC").Limit(10).FromId(1).StartTime(1498793709153).EndTime(1498793709163)
		response, _ := sdk.CompressedTrades(query)

		assert.Equal(t, validCompressedTradesResponse(), response)
	})

	t.Run("It should return error when api fails", func(t *testing.T) {
		mockedClient := NewMockClient(gomock.NewController(t))
		sdk := Sdk{client: mockedClient}
		expectedRequest := newRequest(method, url).Param("symbol", "ETHBTC")

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(nil, errors.New("error"))

		_, err := sdk.CompressedTrades(NewCompressedTradesQuery("ETHBTC"))

		assert.Error(t, err)
	})

	t.Run("It should return error when response cannot be mapped", func(t *testing.T) {
		mockedClient := NewMockClient(gomock.NewController(t))
		sdk := Sdk{client: mockedClient}
		expectedRequest := newRequest(method, url).Param("symbol", "ETHBTC")

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(invalidJson(), nil)

		_, err := sdk.CompressedTrades(NewCompressedTradesQuery("ETHBTC"))

		assert.Error(t, err)
	})
}

func validCompressedTradesJson() []byte {
	return []byte(`[
  		{
    		"a": 26129,         
    		"p": "0.01633102",  
    		"q": "4.70443515",  
    		"f": 27781,         
    		"l": 27781,         
    		"T": 1498793709153, 
    		"m": true,          
    		"M": true           
  		}
	]`)
}

func validCompressedTradesResponse() []CompressedTrade {
	return []CompressedTrade{{
		Id:           26129,
		Price:        0.01633102,
		Quantity:     4.70443515,
		FirstTradeId: 27781,
		LastTradeId:  27781,
		Time:         1498793709153,
		IsBuyerMaker: true,
		IsBestMatch:  true,
	}}
}
