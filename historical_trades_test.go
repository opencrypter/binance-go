package binance

import (
	"errors"
	"github.com/isd4n/binance-go-sdk/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestSdk_HistoricalTrades(t *testing.T) {
	t.Run("It should convert api response to a Trade slice", func(t *testing.T) {
		clientMock := &mocks.Client{}
		sdk := Sdk{client: clientMock}

		expected := validTradesJson()

		clientMock.On("Get", mock.MatchedBy(func(path string) bool {
			return path == "/api/v1/historicalTrades?symbol=ETHBTC&limit=500"
		})).Return(expected, nil)

		response, _ := sdk.HistoricalTrades(NewHistoricalTradesQuery("ETHBTC"))

		assert.Equal(t, validTradesResponse(), response)
	})

	t.Run("It should read optional parameters", func(t *testing.T) {
		clientMock := &mocks.Client{}
		sdk := Sdk{client: clientMock}

		expected := validTradesJson()

		clientMock.On("Get", mock.MatchedBy(func(path string) bool {
			return path == "/api/v1/historicalTrades?symbol=ETHBTC&limit=350&fromId=2300"
		})).Return(expected, nil)

		query := NewHistoricalTradesQuery("ETHBTC").Limit(350).FromId(2300)
		response, _ := sdk.HistoricalTrades(query)

		assert.Equal(t, validTradesResponse(), response)
	})

	t.Run("It should return error when api fails", func(t *testing.T) {
		clientMock := &mocks.Client{}
		sdk := Sdk{client: clientMock}

		expectedError := errors.New("error")

		clientMock.On("Get", mock.MatchedBy(func(path string) bool {
			return path == "/api/v1/historicalTrades?symbol=ETHBTC&limit=500"
		})).Return(nil, expectedError)

		_, err := sdk.HistoricalTrades(NewHistoricalTradesQuery("ETHBTC"))

		assert.Equal(t, expectedError, err)
	})

	t.Run("It should return error when response cannot be mapped", func(t *testing.T) {
		clientMock := &mocks.Client{}
		sdk := Sdk{client: clientMock}

		clientMock.On("Get", mock.MatchedBy(func(path string) bool {
			return path == "/api/v1/historicalTrades?symbol=ETHBTC&limit=500"
		})).Return(invalidTradesJson(), nil)

		_, err := sdk.HistoricalTrades(NewHistoricalTradesQuery("ETHBTC"))

		assert.Error(t, err)
	})
}
