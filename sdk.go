package binanceGoSdk

import (
	"io/ioutil"
	"net/http"
)

const (
	apiBaseUrl = "https://api.binance.com"
)

type Sdk struct {
	client Client
}

type Client interface {
	Get(url string) ([]byte, error)
}

type client struct {
	http      *http.Client
	apiKey    string
	apiSecret string
}

func (_m *client) Get(path string) ([]byte, error) {
	response, err := http.Get(apiBaseUrl + path)
	if err != nil {
		return nil, err
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

func New(apiKey string, apiSecret string) Sdk {
	return Sdk{
		client: &client{
			http:      http.DefaultClient,
			apiKey:    apiKey,
			apiSecret: apiSecret,
		},
	}
}
