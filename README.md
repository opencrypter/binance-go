# Binance-Go
[![Build Status](https://travis-ci.org/isd4n/binance-go.svg?branch=master)](https://travis-ci.org/isd4n/binance-go)
[![codecov](https://codecov.io/gh/isd4n/binance-go/branch/master/graph/badge.svg)](https://codecov.io/gh/isd4n/binance-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

An open source sdk for [Binance exchange](https://www.binance.com) written in Golang. There are endpoints/streams that
are not implemented yet. Feel free to collaborate with me if you need them now :)

## Installation
After you've configured your local Go package:
```bash
go get github.com/isd4n/binance-go
```

## Usage
This SDK is based on the official [binance api docs](https://github.com/binance-exchange/binance-official-api-docs)

You only have to call the constructor function in order to use it:

```go
import "github.com/isd4n/binance-go"

sdk := binance.New("Your-api-key", "Your secret api-key")

exchangeInfo, err := sdk.ExchangeInfo()
```

## Available api endpoints
### ExchangeInfo
Current exchange trading rules and symbol information

Official doc: [exchange-information](https://github.com/binance-exchange/binance-official-api-docs/blob/master/rest-api.md#exchange-information)

#### Example
```go
exchangeInfo, err := sdk.ExchangeInfo()
```

### Order book
Order depth for a specific symbol

Official doc: [Order book](https://github.com/binance-exchange/binance-official-api-docs/blob/master/rest-api.md#order-book)

#### Example
```go
// Example with limit by default
query := binance.NewDepthQuery("ETHBTC")
depth, err := sdk.Depth(query)

// Example with limit parameter
query := binance.NewDepthQuery("ETHBTC").Limit(5)
depth, err := sdk.Depth(query)
```

### Historical trades list
Get historical trades for a specific symbol

Official doc: [Recent trades list](https://github.com/binance-exchange/binance-official-api-docs/blob/master/rest-api.md#old-trade-lookup-market_data)

#### Example
```go
// Example to retrieve last trades with the limit by default
query := binance.NewTradesQuery("ETHBTC")
response, err := sdk.Trades(query)

// Example to retrieve historical trades from the given id and a result limit
query := binance.NewTradesQuery("ETHBTC").Limit(350).FromId(3500)
trades, err := sdk.Trades(query)
```

### Compressed/Aggregate trades list
Get compressed, aggregate trades. Trades that fill at the time, from the same order, with the same price will have the quantity aggregated.

Official doc: [Compressed/Aggregate trades list](https://github.com/binance-exchange/binance-official-api-docs/blob/master/rest-api.md#compressedaggregate-trades-list)

#### Example
```go
// Example with parameters by default
query := binance.NewCompressedTradesQuery("ETHBTC")
response, err := sdk.CompressedTrades(query)

// Example with all parameters. Keep in mind that you cannot use all parameters at the same time (read the official doc)
query := binance.NewCompressedTradesQuery("ETHBTC").Limit(10).FromId(1).StartTime(1498793709153).EndTime(1498793709163)
trades, err := sdk.CompressedTrades(query)
```

### Kline/Candlestick data
Kline/candlestick bars for a symbol. Klines are uniquely identified by their open time.

Official doc: [Kline/Candlestick data](https://github.com/binance-exchange/binance-official-api-docs/blob/master/rest-api.md#klinecandlestick-data)

#### Example
```go
// Example with parameters by default
query := binance.NewKLinesQuery("ETHBTC", binance.Interval12h)
response, err := sdk.KLines(query)

// Example with all parameters.
query := NewKLinesQuery("ETHBTC", Interval1m).Limit(10).StartTime(1498793709153).EndTime(1498793709163)
trades, err := sdk.KLines(query)
```

### Symbol price ticker
Latest price for a symbol.

Official doc: [Symbol price ticker](https://github.com/binance-exchange/binance-official-api-docs/blob/master/rest-api.md#symbol-price-ticker)

#### Example
```go
query := binance.NewSymbolPriceTickerQuery("ETHBTC")
response, err := sdk.SymbolPriceTicker(query)
```

### All symbol price tickers
Latest price for all symbols.

Official doc: [Symbol price ticker](https://github.com/binance-exchange/binance-official-api-docs/blob/master/rest-api.md#symbol-price-ticker)

#### Example
```go
response, err := sdk.AllSymbolPriceTickers()
```

## Available web socket streams:
Not available yet.

## Tests
All is covered 100%. You can run all tests as normally you do it:
```
go test -test.v
```

## License
MIT licensed. See the LICENSE file for details.
