# Binance-Go-SDK
[![Build Status](https://travis-ci.org/isd4n/binance-go-sdk.svg?branch=master)](https://travis-ci.org/isd4n/binance-go-sdk)
[![codecov](https://codecov.io/gh/isd4n/binance-go-sdk/branch/master/graph/badge.svg)](https://codecov.io/gh/isd4n/binance-go-sdk)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

An open source sdk for [Binance exchange](https://www.binance.com) written in Golang. There are endpoints/streams that
are not implemented yet. Feel free to collaborate with me if you need them now :)

## Installation
After you've configured your local Go package:
```bash
go get github.com/isd4n/binance-go-sdk
```

## Usage
This SDK is based on the official [binance api docs](https://github.com/binance-exchange/binance-official-api-docs)

You only have to call the constructor function in order to use it:

```go
import "github.com/isd4n/binance-go-sdk"

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

## Available web socket streams:
Not available yet.

## Tests
All is covered 100%. You can run all tests as normally you do it:
```
go test -test.v
```

## License
MIT licensed. See the LICENSE file for details.
