package binance

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func TestSdk_GetOrder(t *testing.T) {
	method, url := "GET", "/api/v3/order"

	t.Run("It should convert api response to an open order", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockedClock := NewMockClock(ctrl)
		mockedClient := NewMockClient(ctrl)
		sdk := Sdk{client: mockedClient, clock: mockedClock}

		expectedTimestamp := time.Now().Unix()
		mockedClock.EXPECT().Now().Return(&expectedTimestamp)

		expectedRequest := newRequest(method, url).
			Param("symbol", "LTCBTC").
			Param("timestamp", strconv.FormatInt(expectedTimestamp, 10)).
			Sign()

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(validOpenOrderJson(), nil)

		query := NewGetOrderQuery("LTCBTC")
		response, _ := sdk.GetOrder(query)

		assert.Equal(t, validOpenOrderResponse(), response)
	})

	t.Run("It should read optional parameters", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockedClock := NewMockClock(ctrl)
		mockedClient := NewMockClient(ctrl)
		sdk := Sdk{client: mockedClient, clock: mockedClock}

		expectedTimestamp := time.Now().Unix()
		mockedClock.EXPECT().Now().Return(&expectedTimestamp)

		expectedRequest := newRequest(method, url).
			Param("symbol", "LTCBTC").
			Param("orderId", "1").
			Param("origClientOrderId", "myOrder1").
			Param("recvWindow", "2").
			Param("timestamp", strconv.FormatInt(expectedTimestamp, 10)).
			Sign()

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(validOpenOrderJson(), nil)

		query := NewGetOrderQuery("LTCBTC").
			OrderId(1).
			OrigClientOrderId("myOrder1").
			RecvWindow(2)

		response, _ := sdk.GetOrder(query)

		assert.Equal(t, validOpenOrderResponse(), response)
	})

	t.Run("It should return error when api fails", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockedClock := NewMockClock(ctrl)
		mockedClient := NewMockClient(ctrl)
		sdk := Sdk{client: mockedClient, clock: mockedClock}

		expectedTimestamp := time.Now().Unix()
		mockedClock.EXPECT().Now().Return(&expectedTimestamp)

		expectedRequest := newRequest(method, url).
			Param("symbol", "LTCBTC").
			Param("timestamp", strconv.FormatInt(expectedTimestamp, 10)).
			Sign()

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(nil, errors.New("error"))

		query := NewGetOrderQuery("LTCBTC")
		_, err := sdk.GetOrder(query)

		assert.Error(t, err)
	})

	t.Run("It should return error when response cannot be mapped", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockedClock := NewMockClock(ctrl)
		mockedClient := NewMockClient(ctrl)
		sdk := Sdk{client: mockedClient, clock: mockedClock}

		expectedTimestamp := time.Now().Unix()
		mockedClock.EXPECT().Now().Return(&expectedTimestamp)

		expectedRequest := newRequest(method, url).
			Param("symbol", "LTCBTC").
			Param("timestamp", strconv.FormatInt(expectedTimestamp, 10)).
			Sign()

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(invalidJson(), nil)

		query := NewGetOrderQuery("LTCBTC")
		_, err := sdk.GetOrder(query)

		assert.Error(t, err)
	})
}

func validOpenOrderJson() []byte {
	return []byte(`{
  		"symbol": "LTCBTC",
  		"orderId": 1,
  		"clientOrderId": "myOrder1",
		"price": "0.1",
  		"origQty": "1.0",
  		"executedQty": "0.0",
  		"status": "NEW",
  		"timeInForce": "GTC",
  		"type": "LIMIT",
  		"side": "BUY",
  		"stopPrice": "0.0",
  		"icebergQty": "0.0",
  		"time": 1499827319559,
  		"isWorking": true
	}`)
}

func validOpenOrderResponse() *OpenOrder {
	return &OpenOrder{
		Symbol:           "LTCBTC",
		OrderId:          1,
		ClientOrderId:    "myOrder1",
		Price:            0.1,
		OriginalQuantity: 1,
		ExecutedQuantity: 0,
		Status:           "NEW",
		TimeInForce:      "GTC",
		Type:             "LIMIT",
		Side:             "BUY",
		StopPrice:        0.0,
		IcebergQuantity:  0.0,
		Time:             1499827319559,
		IsWorking:        true,
	}
}
