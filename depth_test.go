package binance

import (
	"errors"
	"github.com/isd4n/binance-go/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestSdk_Depth(t *testing.T) {
	t.Run("It should convert api response to a Depth", func(t *testing.T) {
		clientMock := &mocks.Client{}
		sdk := Sdk{client: clientMock}

		expected := validDepthJson()

		clientMock.On("Get", mock.MatchedBy(func(path string) bool {
			return path == "/api/v1/depth?symbol=ETHBTC&limit=100"
		})).Return(expected, nil)

		response, _ := sdk.Depth(NewDepthQuery("ETHBTC"))

		assert.Equal(t, validDepthResponse(), response)
	})

	t.Run("It should read optional parameters", func(t *testing.T) {
		clientMock := &mocks.Client{}
		sdk := Sdk{client: clientMock}

		expected := validDepthJson()

		clientMock.On("Get", mock.MatchedBy(func(path string) bool {
			return path == "/api/v1/depth?symbol=ETHBTC&limit=500"
		})).Return(expected, nil)

		query := NewDepthQuery("ETHBTC").Limit(500)
		response, _ := sdk.Depth(query)

		assert.Equal(t, validDepthResponse(), response)
	})

	t.Run("It should return error when api fails", func(t *testing.T) {
		clientMock := &mocks.Client{}
		sdk := Sdk{client: clientMock}

		expectedError := errors.New("error")

		clientMock.On("Get", mock.MatchedBy(func(path string) bool {
			return path == "/api/v1/depth?symbol=ETHBTC&limit=100"
		})).Return(nil, expectedError)

		_, err := sdk.Depth(NewDepthQuery("ETHBTC"))

		assert.Equal(t, expectedError, err)
	})

	t.Run("It should return error when response cannot be mapped", func(t *testing.T) {
		clientMock := &mocks.Client{}
		sdk := Sdk{client: clientMock}

		clientMock.On("Get", mock.MatchedBy(func(path string) bool {
			return path == "/api/v1/depth?symbol=ETHBTC&limit=100"
		})).Return(invalidJson(), nil)

		_, err := sdk.Depth(NewDepthQuery("ETHBTC"))

		assert.Error(t, err)
	})
}

func validDepthJson() []byte {
	return []byte(`{
  		"lastUpdateId": 1027024,
  		"bids": [
    		[
      			"4.00000000",     
      			"431.00000000",   
      			[]                
    		]
  		],
  		"asks": [
    		[
      			"4.00000200",
      			"12.05000000",
      			[]
    		]
  		]
	}`)
}

func validDepthResponse() *Depth {
	return &Depth{
		LastUpdateId: 1027024,
		Bids: []DepthOrder{
			{Price: 4.0, Quantity: 431.0},
		},
		Asks: []DepthOrder{
			{Price: 4.000002, Quantity: 12.05},
		},
	}
}
