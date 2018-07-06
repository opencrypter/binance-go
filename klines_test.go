package binance

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSdk_KLines(t *testing.T) {
	method, url := "GET", "/api/v1/klines"

	t.Run("It should convert api response to a compressed KLine slice", func(t *testing.T) {
		mockedClient := NewMockClient(gomock.NewController(t))
		sdk := Sdk{client: mockedClient}

		expectedRequest := newRequest(method, url).
			Param("symbol", "ETHBTC").
			Param("interval", string(Interval1m))

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(validKLinesJson(), nil)

		response, _ := sdk.KLines(NewKLinesQuery("ETHBTC", Interval1m))

		assert.Equal(t, validKLinesResponse(), response)
	})

	t.Run("It should read optional parameters", func(t *testing.T) {
		mockedClient := NewMockClient(gomock.NewController(t))
		sdk := Sdk{client: mockedClient}

		expectedRequest := newRequest(method, url).
			Param("symbol", "ETHBTC").
			Param("interval", string(Interval1m)).
			Param("limit", "10").
			Param("startTime", "1498793709153").
			Param("endTime", "1498793709163")

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(validKLinesJson(), nil)

		query := NewKLinesQuery("ETHBTC", Interval1m).Limit(10).StartTime(1498793709153).EndTime(1498793709163)
		response, _ := sdk.KLines(query)

		assert.Equal(t, validKLinesResponse(), response)
	})

	t.Run("It should return error when api fails", func(t *testing.T) {
		mockedClient := NewMockClient(gomock.NewController(t))
		sdk := Sdk{client: mockedClient}

		expectedRequest := newRequest(method, url).
			Param("symbol", "ETHBTC").
			Param("interval", string(Interval1m))

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(nil, errors.New("error"))

		_, err := sdk.KLines(NewKLinesQuery("ETHBTC", Interval1m))

		assert.Error(t, err)
	})

	t.Run("It should return error when response cannot be mapped", func(t *testing.T) {
		mockedClient := NewMockClient(gomock.NewController(t))
		sdk := Sdk{client: mockedClient}

		expectedRequest := newRequest(method, url).
			Param("symbol", "ETHBTC").
			Param("interval", string(Interval1m))

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(invalidJson(), nil)

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
