package binance

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func TestSdk_MyTrades(t *testing.T) {
	method, url := "GET", "/api/v3/myTrades"

	t.Run("It should convert api response to account trades", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockedClock := NewMockClock(ctrl)
		mockedClient := NewMockClient(ctrl)
		sdk := Sdk{client: mockedClient, clock: mockedClock}

		expectedTimestamp := time.Now().Unix()
		mockedClock.EXPECT().Now().Return(&expectedTimestamp)

		expectedRequest := newRequest(method, url).
			Param("symbol", "ETHBTC").
			Param("timestamp", strconv.FormatInt(expectedTimestamp, 10)).
			Sign()

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(validAccountTradesJson(), nil)

		query := NewMyTradesQuery("ETHBTC")
		response, _ := sdk.MyTrades(query)

		assert.Equal(t, validAccountTradesResponse(), response)
	})

	t.Run("It should read optional parameters", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockedClock := NewMockClock(ctrl)
		mockedClient := NewMockClient(ctrl)
		sdk := Sdk{client: mockedClient, clock: mockedClock}

		expectedTimestamp := time.Now().Unix()
		mockedClock.EXPECT().Now().Return(&expectedTimestamp)

		expectedRequest := newRequest(method, url).
			Param("symbol", "ETHBTC").
			Param("limit", "10").
			Param("fromId", "200").
			Param("recvWindow", "2").
			Param("timestamp", strconv.FormatInt(expectedTimestamp, 10)).
			Sign()

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(validAccountTradesJson(), nil)

		query := NewMyTradesQuery("ETHBTC").Limit(10).RecvWindow(2).FromId(200)
		response, _ := sdk.MyTrades(query)

		assert.Equal(t, validAccountTradesResponse(), response)
	})

	t.Run("It should return error when api fails", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockedClock := NewMockClock(ctrl)
		mockedClient := NewMockClient(ctrl)
		sdk := Sdk{client: mockedClient, clock: mockedClock}

		expectedTimestamp := time.Now().Unix()
		mockedClock.EXPECT().Now().Return(&expectedTimestamp)

		expectedRequest := newRequest(method, url).
			Param("symbol", "ETHBTC").
			Param("timestamp", strconv.FormatInt(expectedTimestamp, 10)).
			Sign()

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(nil, errors.New("error"))

		query := NewMyTradesQuery("ETHBTC")
		_, err := sdk.MyTrades(query)

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
			Param("symbol", "ETHBTC").
			Param("timestamp", strconv.FormatInt(expectedTimestamp, 10)).
			Sign()

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(invalidJson(), nil)

		query := NewMyTradesQuery("ETHBTC")
		_, err := sdk.MyTrades(query)

		assert.Error(t, err)
	})
}

func validAccountTradesJson() []byte {
	return []byte(`[
  		{
    		"id": 28457,
    		"orderId": 100234,
    		"price": "4.00000100",
    		"qty": "12.00000000",
    		"commission": "10.10000000",
    		"commissionAsset": "BNB",
    		"time": 1499865549590,
    		"isBuyer": true,
    		"isMaker": false,
    		"isBestMatch": true
  		}
	]`)
}

func validAccountTradesResponse() []AccountTrade {
	return []AccountTrade{{
		Id:              28457,
		OrderId:         100234,
		Price:           4.000001,
		Quantity:        12,
		Commission:      10.1,
		CommissionAsset: "BNB",
		Time:            1499865549590,
		IsBuyer:         true,
		IsMaker:         false,
		IsBestMatch:     true,
	}}
}
