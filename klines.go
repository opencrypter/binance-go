package binance

import (
	"encoding/json"
	"strconv"
)

const (
	Interval1m  Interval = "1m"
	Interval3m  Interval = "3m"
	Interval5m  Interval = "5m"
	Interval15m Interval = "15m"
	Interval30m Interval = "30m"
	Interval1h  Interval = "1h"
	Interval2h  Interval = "2h"
	Interval4h  Interval = "4h"
	Interval8h  Interval = "8h"
	Interval12h Interval = "12h"
	Interval1d  Interval = "1d"
	Interval3d  Interval = "3d"
	Interval1w  Interval = "1w"
	Interval1M  Interval = "1M"
)

type Interval string

type KLine struct {
	OpenTime            int
	Open                float64
	High                float64
	Low                 float64
	Close               float64
	Volume              float64
	CloseTime           int
	QuoteAssetVolume    float64
	TradesNumber        int
	TakerBuyBaseVolume  float64
	TakerBuyQuoteVolume float64
}

type kLinesQuery struct {
	symbol    string
	interval  Interval
	limit     int
	startTime int
	endTime   int
}

func NewKLinesQuery(symbol string, interval Interval) *kLinesQuery {
	return &kLinesQuery{symbol: symbol, interval: interval}
}

func (t *kLinesQuery) Limit(limit int) *kLinesQuery {
	t.limit = limit
	return t
}
func (t *kLinesQuery) StartTime(startTime int) *kLinesQuery {
	t.startTime = startTime
	return t
}

func (t *kLinesQuery) EndTime(endTime int) *kLinesQuery {
	t.endTime = endTime
	return t
}

func parseKLinesResponse(jsonContent []byte) ([]KLine, error) {
	elements := make([][]interface{}, 0)
	err := json.Unmarshal(jsonContent, &elements)

	kLines := make([]KLine, 0)
	for _, item := range elements {
		open, _ := strconv.ParseFloat(item[1].(string), 64)
		high, _ := strconv.ParseFloat(item[2].(string), 64)
		low, _ := strconv.ParseFloat(item[3].(string), 64)
		close, _ := strconv.ParseFloat(item[4].(string), 64)
		volume, _ := strconv.ParseFloat(item[5].(string), 64)
		quoteVolume, _ := strconv.ParseFloat(item[7].(string), 64)
		takerBuyBaseVolume, _ := strconv.ParseFloat(item[9].(string), 64)
		takerBuyQuoteVolume, _ := strconv.ParseFloat(item[10].(string), 64)

		kLines = append(kLines, KLine{
			OpenTime:            int(item[0].(float64)),
			Open:                open,
			High:                high,
			Low:                 low,
			Close:               close,
			Volume:              volume,
			CloseTime:           int(item[6].(float64)),
			QuoteAssetVolume:    quoteVolume,
			TradesNumber:        int(item[8].(float64)),
			TakerBuyBaseVolume:  takerBuyBaseVolume,
			TakerBuyQuoteVolume: takerBuyQuoteVolume,
		})
	}

	return kLines, err
}

func (sdk *Sdk) KLines(query *kLinesQuery) ([]KLine, error) {
	request := newRequest("GET", "/api/v1/klines").
		Param("symbol", query.symbol).
		Param("interval", string(query.interval))

	if query.limit > 0 {
		request.Param("limit", strconv.Itoa(query.limit))
	}

	if query.startTime > 0 {
		request.Param("startTime", strconv.Itoa(query.startTime))
	}

	if query.endTime > 0 {
		request.Param("endTime", strconv.Itoa(query.endTime))
	}

	responseContent, err := sdk.client.Do(request)
	if err != nil {
		return nil, err
	}

	return parseKLinesResponse(responseContent)
}
