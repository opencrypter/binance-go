package binance

import (
	"errors"
	"github.com/isd4n/binance-go/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestSdk_CompressedTrades(t *testing.T) {
	t.Run("It should convert api response to a compressed Trade slice", func(t *testing.T) {
		clientMock := &mocks.Client{}
		sdk := Sdk{client: clientMock}

		expected := validCompressedTradesJson()

		clientMock.On("Get", mock.MatchedBy(func(path string) bool {
			return path == "/api/v1/aggTrades?symbol=ETHBTC"
		})).Return(expected, nil)

		response, _ := sdk.CompressedTrades(NewCompressedTradesQuery("ETHBTC"))

		assert.Equal(t, validCompressedTradesResponse(), response)
	})

	t.Run("It should read optional parameters", func(t *testing.T) {
		clientMock := &mocks.Client{}
		sdk := Sdk{client: clientMock}

		expected := validCompressedTradesJson()

		clientMock.On("Get", mock.MatchedBy(func(path string) bool {
			return path == "/api/v1/aggTrades"+
				"?symbol=ETHBTC&limit=10&fromId=1&startTime=1498793709153&endTime=1498793709163"
		})).Return(expected, nil)

		query := NewCompressedTradesQuery("ETHBTC").Limit(10).FromId(1).StartTime(1498793709153).EndTime(1498793709163)
		response, _ := sdk.CompressedTrades(query)

		assert.Equal(t, validCompressedTradesResponse(), response)
	})

	t.Run("It should return error when api fails", func(t *testing.T) {
		clientMock := &mocks.Client{}
		sdk := Sdk{client: clientMock}

		expectedError := errors.New("error")

		clientMock.On("Get", mock.MatchedBy(func(path string) bool {
			return path == "/api/v1/aggTrades?symbol=ETHBTC"
		})).Return(nil, expectedError)

		_, err := sdk.CompressedTrades(NewCompressedTradesQuery("ETHBTC"))

		assert.Equal(t, expectedError, err)
	})

	t.Run("It should return error when response cannot be mapped", func(t *testing.T) {
		clientMock := &mocks.Client{}
		sdk := Sdk{client: clientMock}

		clientMock.On("Get", mock.MatchedBy(func(path string) bool {
			return path == "/api/v1/aggTrades?symbol=ETHBTC"
		})).Return(invalidJson(), nil)

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
