package binance

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSdk_Depth(t *testing.T) {
	method, url := "GET", "/api/v1/depth"

	t.Run("It should convert api response to a Depth", func(t *testing.T) {
		mockedClient := NewMockClient(gomock.NewController(t))
		sdk := Sdk{client: mockedClient}

		expectedRequest := newRequest(method, url).
			Param("symbol", "ETHBTC").
			Param("limit", "100")

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(validDepthJson(), nil)

		response, _ := sdk.Depth(NewDepthQuery("ETHBTC"))

		assert.Equal(t, validDepthResponse(), response)
	})

	t.Run("It should read optional parameters", func(t *testing.T) {
		mockedClient := NewMockClient(gomock.NewController(t))
		sdk := Sdk{client: mockedClient}

		expectedRequest := newRequest(method, url).
			Param("symbol", "ETHBTC").
			Param("limit", "500")

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(validDepthJson(), nil)

		query := NewDepthQuery("ETHBTC").Limit(500)
		response, _ := sdk.Depth(query)

		assert.Equal(t, validDepthResponse(), response)
	})

	t.Run("It should return error when api fails", func(t *testing.T) {
		mockedClient := NewMockClient(gomock.NewController(t))
		sdk := Sdk{client: mockedClient}

		expectedRequest := newRequest(method, url).
			Param("symbol", "ETHBTC").
			Param("limit", "100")

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(nil, errors.New("error"))

		_, err := sdk.Depth(NewDepthQuery("ETHBTC"))

		assert.Error(t, err)
	})

	t.Run("It should return error when response cannot be mapped", func(t *testing.T) {
		mockedClient := NewMockClient(gomock.NewController(t))
		sdk := Sdk{client: mockedClient}

		expectedRequest := newRequest(method, url).
			Param("symbol", "ETHBTC").
			Param("limit", "100")

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(invalidJson(), nil)

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
