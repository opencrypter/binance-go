package binance

import (
	"encoding/json"
)

type Account struct {
	MakerCommission  int64
	TakerCommission  int64
	BuyerCommission  int64
	SellerCommission int64
	CanTrade         bool
	CanWithdraw      bool
	CanDeposit       bool
	UpdateTime       int64
	Balances         []Balance
}

type Balance struct {
	Asset  string
	Free   float64 `json:"free,string"`
	Locked float64 `json:"locked,string"`
}

type accountQuery struct {
	recvWindow *int64
	timestamp  *int64
}

func NewAccountQuery() *accountQuery {
	return &accountQuery{}
}

func (r *accountQuery) RecvWindow(value int64) *accountQuery {
	r.recvWindow = &value
	return r
}

func (sdk Sdk) Account(query *accountQuery) (*Account, error) {
	req := newRequest("GET", "/api/v3/account").
		Int64Param("recvWindow", query.recvWindow).
		Int64Param("timestamp", sdk.clock.Now()).
		Sign()

	responseContent, err := sdk.client.Do(req)
	if err != nil {
		return nil, err
	}

	return parseAccountResponse(responseContent)
}

func parseAccountResponse(jsonContent []byte) (*Account, error) {
	response := &Account{}
	err := json.Unmarshal(jsonContent, &response)
	return response, err
}
