package binance

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func TestSdk_CancelOrder(t *testing.T) {
	method, url := "DELETE", "/api/v3/order"

	t.Run("It should convert api response to a cancelled order", func(t *testing.T) {
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
			Return(validCancelledOrderJson(), nil)

		request := NewCancelOrderRequest("LTCBTC")
		response, _ := sdk.CancelOrder(request)

		assert.Equal(t, validCancelledOrderResponse(), response)
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
			Return(validCancelledOrderJson(), nil)

		request := NewCancelOrderRequest("LTCBTC").
			OrderId(1).
			OrigClientOrderId("myOrder1").
			RecvWindow(2)

		response, _ := sdk.CancelOrder(request)

		assert.Equal(t, validCancelledOrderResponse(), response)
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

		request := NewCancelOrderRequest("LTCBTC")
		_, err := sdk.CancelOrder(request)

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

		request := NewCancelOrderRequest("LTCBTC")
		_, err := sdk.CancelOrder(request)

		assert.Error(t, err)
	})
}

func validCancelledOrderJson() []byte {
	return []byte(`{
  		"symbol": "LTCBTC",
  		"origClientOrderId": "myOrder1",
  		"orderId": 1,
  		"clientOrderId": "cancelMyOrder1"
	}`)
}

func validCancelledOrderResponse() *CancelledOrder {
	return &CancelledOrder{
		Symbol:            "LTCBTC",
		OrigClientOrderId: "myOrder1",
		OrderId:           1,
		ClientOrderId:     "cancelMyOrder1",
	}
}
