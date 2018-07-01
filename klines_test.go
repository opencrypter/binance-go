package binance

import (
	"errors"
	"github.com/isd4n/binance-go/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestSdk_KLines(t *testing.T) {
	t.Run("It should convert api response to a compressed KLine slice", func(t *testing.T) {
		clientMock := &mocks.Client{}
		sdk := Sdk{client: clientMock}

		expected := validKLinesJson()

		clientMock.On("Get", mock.MatchedBy(func(path string) bool {
			return path == "/api/v1/klines?symbol=ETHBTC&interval=1m"
		})).Return(expected, nil)

		response, _ := sdk.KLines(NewKLinesQuery("ETHBTC", Interval1m))

		assert.Equal(t, validKLinesResponse(), response)
	})

	t.Run("It should read optional parameters", func(t *testing.T) {
		clientMock := &mocks.Client{}
		sdk := Sdk{client: clientMock}

		expected := validKLinesJson()

		clientMock.On("Get", mock.MatchedBy(func(path string) bool {
			return path == "/api/v1/klines?symbol=ETHBTC&interval=1m"+
				"&limit=10&startTime=1498793709153&endTime=1498793709163"
		})).Return(expected, nil)

		query := NewKLinesQuery("ETHBTC", Interval1m).Limit(10).StartTime(1498793709153).EndTime(1498793709163)
		response, _ := sdk.KLines(query)

		assert.Equal(t, validKLinesResponse(), response)
	})

	t.Run("It should return error when api fails", func(t *testing.T) {
		clientMock := &mocks.Client{}
		sdk := Sdk{client: clientMock}

		expectedError := errors.New("error")

		clientMock.On("Get", mock.MatchedBy(func(path string) bool {
			return path == "/api/v1/klines?symbol=ETHBTC&interval=1m"
		})).Return(nil, expectedError)

		_, err := sdk.KLines(NewKLinesQuery("ETHBTC", Interval1m))

		assert.Equal(t, expectedError, err)
	})

	t.Run("It should return error when response cannot be mapped", func(t *testing.T) {
		clientMock := &mocks.Client{}
		sdk := Sdk{client: clientMock}

		clientMock.On("Get", mock.MatchedBy(func(path string) bool {
			return path == "/api/v1/klines?symbol=ETHBTC&interval=1m"
		})).Return(invalidJson(), nil)

		_, err := sdk.KLines(NewKLinesQuery("ETHBTC", Interval1m))

		assert.Error(t, err)
	})
}

func validKLinesJson() []byte {
	return []byte(`[
		[	
    		1499040000000,      
    		"0.01634790",       
    		"0.80000000",       
    		"0.01575800",       
    		"0.01577100",       
    		"148976.11427815",  
    		1499644799999,      
    		"2434.19055334",    
    		308,                
    		"1756.87402397",    
    		"28.46694368",      
    		"17928899.62484339" 
  		]
	]`)
}

func validKLinesResponse() []KLine {
	return []KLine{{
		OpenTime:            1499040000000,
		Open:                0.01634790,
		High:                0.8,
		Low:                 0.015758,
		Close:               0.015771,
		Volume:              148976.11427815,
		CloseTime:           1499644799999,
		QuoteAssetVolume:    2434.19055334,
		TradesNumber:        308,
		TakerBuyBaseVolume:  1756.87402397,
		TakerBuyQuoteVolume: 28.46694368,
	}}
}
