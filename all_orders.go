package binance

type getAllOrdersQuery struct {
	symbol     *string
	orderId    *int64
	limit      *int64
	recvWindow *int64
	timestamp  *int64
}

func NewGetAllOrdersQuery(symbol string) *getAllOrdersQuery {
	return &getAllOrdersQuery{
		symbol: &symbol,
	}
}

func (r *getAllOrdersQuery) OrderId(value int64) *getAllOrdersQuery {
	r.orderId = &value
	return r
}

func (r *getAllOrdersQuery) Limit(value int64) *getAllOrdersQuery {
	r.limit = &value
	return r
}

func (r *getAllOrdersQuery) RecvWindow(value int64) *getAllOrdersQuery {
	r.recvWindow = &value
	return r
}

func (sdk Sdk) GetAllOrders(query *getAllOrdersQuery) ([]Order, error) {
	req := newRequest("GET", "/api/v3/allOrders").
		StringParam("symbol", query.symbol).
		Int64Param("orderId", query.orderId).
		Int64Param("limit", query.limit).
		Int64Param("recvWindow", query.recvWindow).
		Int64Param("timestamp", sdk.clock.Now()).
		Sign()

	responseContent, err := sdk.client.Do(req)
	if err != nil {
		return nil, err
	}

	return parseOrdersResponse(responseContent)
}
