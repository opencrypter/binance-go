package binance

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func TestSdk_NewOrder(t *testing.T) {
	method, url := "POST", "/api/v3/order"

	t.Run("It should convert api response to a FullOrder", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockedClock := NewMockClock(ctrl)
		mockedClient := NewMockClient(ctrl)
		sdk := Sdk{client: mockedClient, clock: mockedClock}

		expectedTimestamp := time.Now().Unix()
		mockedClock.EXPECT().Now().Return(&expectedTimestamp)

		expectedRequest := newRequest(method, url).
			Param("symbol", "BTCUSDT").
			Param("side", "SELL").
			Param("type", "MARKET").
			Param("quantity", "10.00000000").
			Param("newOrderRespType", "FULL").
			Param("timestamp", strconv.FormatInt(expectedTimestamp, 10)).
			Sign()

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(validFullOrderJson(), nil)

		request := NewOrderRequest("BTCUSDT", "SELL", "MARKET", 10)
		response, _ := sdk.NewOrder(request)

		assert.Equal(t, validFullOrderResponse(), response)
	})

	t.Run("It should read optional parameters", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockedClock := NewMockClock(ctrl)
		mockedClient := NewMockClient(ctrl)
		sdk := Sdk{client: mockedClient, clock: mockedClock}

		expectedTimestamp := time.Now().Unix()
		mockedClock.EXPECT().Now().Return(&expectedTimestamp)

		expectedRequest := newRequest(method, url).
			Param("symbol", "BTCUSDT").
			Param("side", "SELL").
			Param("type", "MARKET").
			Param("quantity", "10.00000000").
			Param("timeInForce", "GTC").
			Param("price", "0.10000000").
			Param("newClientOrderId", "6gCrw2kRUAF9CvJDGP16IP").
			Param("stopPrice", "0.10000000").
			Param("icebergQty", "0.10000000").
			Param("newOrderRespType", "ACK").
			Param("recvWindow", "2").
			Param("timestamp", strconv.FormatInt(expectedTimestamp, 10)).
			Sign()

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(validFullOrderJson(), nil)

		request := NewOrderRequest("BTCUSDT", "SELL", "MARKET", 10).
			TimeInForce("GTC").
			Price(0.1).
			NewClientOrderId("6gCrw2kRUAF9CvJDGP16IP").
			StopPrice(0.1).
			IcebergQuantity(0.1).
			NewOrderResponseType("ACK").
			RecvWindow(2)

		response, _ := sdk.NewOrder(request)

		assert.Equal(t, validFullOrderResponse(), response)
	})

	t.Run("It should return error when api fails", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockedClock := NewMockClock(ctrl)
		mockedClient := NewMockClient(ctrl)
		sdk := Sdk{client: mockedClient, clock: mockedClock}

		expectedTimestamp := time.Now().Unix()
		mockedClock.EXPECT().Now().Return(&expectedTimestamp)

		expectedRequest := newRequest(method, url).
			Param("symbol", "BTCUSDT").
			Param("side", "SELL").
			Param("type", "MARKET").
			Param("quantity", "10.00000000").
			Param("newOrderRespType", "FULL").
			Param("timestamp", strconv.FormatInt(expectedTimestamp, 10)).
			Sign()

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(nil, errors.New("error"))

		request := NewOrderRequest("BTCUSDT", "SELL", "MARKET", 10)
		_, err := sdk.NewOrder(request)

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
			Param("symbol", "BTCUSDT").
			Param("side", "SELL").
			Param("type", "MARKET").
			Param("quantity", "10.00000000").
			Param("newOrderRespType", "FULL").
			Param("timestamp", strconv.FormatInt(expectedTimestamp, 10)).
			Sign()

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(invalidJson(), nil)

		request := NewOrderRequest("BTCUSDT", "SELL", "MARKET", 10)
		_, err := sdk.NewOrder(request)

		assert.Error(t, err)
	})
}

func validFullOrderJson() []byte {
	return []byte(`{
  		"symbol": "BTCUSDT",
  		"orderId": 28,
  		"clientOrderId": "6gCrw2kRUAF9CvJDGP16IP",
  		"transactTime": 1507725176595,
  		"price": "0.1",
  		"origQty": "10.00000000",
  		"executedQty": "10.00000000",
  		"status": "FILLED",
  		"timeInForce": "GTC",
  		"type": "MARKET",
  		"side": "SELL",
  		"fills": [
    		{
      			"price": "4000.00000000",
      			"qty": "10.00000000",
      			"commission": "4.00000000",
      			"commissionAsset": "USDT"
    		}
  		]
	}`)
}

func validFullOrderResponse() *FullOrder {
	return &FullOrder{
		Symbol:           "BTCUSDT",
		OrderId:          28,
		ClientOrderId:    "6gCrw2kRUAF9CvJDGP16IP",
		TransactionTime:  1507725176595,
		Price:            0.1,
		OriginalQuantity: 10,
		ExecutedQuantity: 10,
		Status:           "FILLED",
		TimeInForce:      "GTC",
		Type:             "MARKET",
		Side:             "SELL",
		Fills: []OrderFill{{
			Price:           4000,
			Quantity:        10,
			Commission:      4,
			CommissionAsset: "USDT",
		}},
	}
}
