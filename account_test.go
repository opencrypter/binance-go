package binance

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func TestSdk_Account(t *testing.T) {
	method, url := "GET", "/api/v3/account"

	t.Run("It should convert api response to an account", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockedClock := NewMockClock(ctrl)
		mockedClient := NewMockClient(ctrl)
		sdk := Sdk{client: mockedClient, clock: mockedClock}

		expectedTimestamp := time.Now().Unix()
		mockedClock.EXPECT().Now().Return(&expectedTimestamp)

		expectedRequest := newRequest(method, url).
			Param("timestamp", strconv.FormatInt(expectedTimestamp, 10)).
			Sign()

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(validAccountJson(), nil)

		query := NewAccountQuery()
		response, _ := sdk.Account(query)

		assert.Equal(t, validAccountResponse(), response)
	})

	t.Run("It should read optional parameters", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockedClock := NewMockClock(ctrl)
		mockedClient := NewMockClient(ctrl)
		sdk := Sdk{client: mockedClient, clock: mockedClock}

		expectedTimestamp := time.Now().Unix()
		mockedClock.EXPECT().Now().Return(&expectedTimestamp)

		expectedRequest := newRequest(method, url).
			Param("recvWindow", "2").
			Param("timestamp", strconv.FormatInt(expectedTimestamp, 10)).
			Sign()

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(validAccountJson(), nil)

		query := NewAccountQuery().RecvWindow(2)
		response, _ := sdk.Account(query)

		assert.Equal(t, validAccountResponse(), response)
	})

	t.Run("It should return error when api fails", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockedClock := NewMockClock(ctrl)
		mockedClient := NewMockClient(ctrl)
		sdk := Sdk{client: mockedClient, clock: mockedClock}

		expectedTimestamp := time.Now().Unix()
		mockedClock.EXPECT().Now().Return(&expectedTimestamp)

		expectedRequest := newRequest(method, url).
			Param("timestamp", strconv.FormatInt(expectedTimestamp, 10)).
			Sign()

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(nil, errors.New("error"))

		query := NewAccountQuery()
		_, err := sdk.Account(query)

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
			Param("timestamp", strconv.FormatInt(expectedTimestamp, 10)).
			Sign()

		mockedClient.
			EXPECT().
			Do(expectedRequest).
			MinTimes(1).
			Return(invalidJson(), nil)

		query := NewAccountQuery()
		_, err := sdk.Account(query)

		assert.Error(t, err)
	})
}

func validAccountJson() []byte {
	return []byte(`{
  		"makerCommission": 15,
  		"takerCommission": 15,
  		"buyerCommission": 0,
  		"sellerCommission": 0,
  		"canTrade": true,
  		"canWithdraw": true,
  		"canDeposit": true,
  		"updateTime": 123456789,
  		"balances": [
    		{
      			"asset": "BTC",
      			"free": "4723846.89208129",
      			"locked": "0.00000000"
    		},
    		{
      			"asset": "LTC",
      			"free": "4763368.68006011",
      			"locked": "0.00000000"
    		}
		]
	}`)
}

func validAccountResponse() *Account {
	return &Account{
		MakerCommission:  15,
		TakerCommission:  15,
		BuyerCommission:  0,
		SellerCommission: 0,
		CanTrade:         true,
		CanWithdraw:      true,
		CanDeposit:       true,
		UpdateTime:       123456789,
		Balances: []Balance{
			{
				Asset:  "BTC",
				Free:   4723846.89208129,
				Locked: 0,
			}, {
				Asset:  "LTC",
				Free:   4763368.68006011,
				Locked: 0,
			},
		},
	}
}
