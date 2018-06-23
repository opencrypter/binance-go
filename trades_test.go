package binance

import (
	"errors"
	"github.com/isd4n/binance-go-sdk/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestSdk_Trades(t *testing.T) {
	t.Run("It should convert api response to a Trade slice", func(t *testing.T) {
		clientMock := &mocks.Client{}
		sdk := Sdk{client: clientMock}

		expected := validTradesJson()

		clientMock.On("Get", mock.MatchedBy(func(path string) bool {
			return path == "/api/v1/trades?symbol=ETHBTC&limit=500"
		})).Return(expected, nil)

		response, _ := sdk.Trades(NewTradesQuery("ETHBTC"))

		assert.Equal(t, validTradesResponse(), response)
	})

	t.Run("It should read optional parameters", func(t *testing.T) {
		clientMock := &mocks.Client{}
		sdk := Sdk{client: clientMock}

		expected := validTradesJson()

		clientMock.On("Get", mock.MatchedBy(func(path string) bool {
			return path == "/api/v1/trades?symbol=ETHBTC&limit=350"
		})).Return(expected, nil)

		query := NewTradesQuery("ETHBTC").Limit(350)
		response, _ := sdk.Trades(query)

		assert.Equal(t, validTradesResponse(), response)
	})

	t.Run("It should return error when api fails", func(t *testing.T) {
		clientMock := &mocks.Client{}
		sdk := Sdk{client: clientMock}

		expectedError := errors.New("error")

		clientMock.On("Get", mock.MatchedBy(func(path string) bool {
			return path == "/api/v1/trades?symbol=ETHBTC&limit=500"
		})).Return(nil, expectedError)

		_, err := sdk.Trades(NewTradesQuery("ETHBTC"))

		assert.Equal(t, expectedError, err)
	})

	t.Run("It should return error when response cannot be mapped", func(t *testing.T) {
		clientMock := &mocks.Client{}
		sdk := Sdk{client: clientMock}

		clientMock.On("Get", mock.MatchedBy(func(path string) bool {
			return path == "/api/v1/trades?symbol=ETHBTC&limit=500"
		})).Return(invalidTradesJson(), nil)

		_, err := sdk.Trades(NewTradesQuery("ETHBTC"))

		assert.Error(t, err)
	})
}

func invalidTradesJson() []byte {
	return []byte(`<h1>Page Not available</h1>`)
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
