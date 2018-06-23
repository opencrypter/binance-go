package binance

import (
	"errors"
	"io/ioutil"
	"net/http"
)

type Sdk struct {
	client Client
}

type Client interface {
	Get(url string) ([]byte, error)
}

type client struct {
	baseUrl   string
	apiKey    string
	apiSecret string
}

func (c *client) Get(path string) ([]byte, error) {
	request, _ := http.NewRequest("GET", c.baseUrl+path, nil)
	request.Header.Set("X-MBX-APIKEY", c.apiKey)

	client := http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	responseBody, _ := ioutil.ReadAll(response.Body)
	if response.StatusCode >= 300 {
		return nil, errors.New("Error " + string(response.StatusCode) + ": " + string(responseBody))
	}

	return responseBody, nil
}

func New(apiKey string, apiSecret string) Sdk {
	return Sdk{
		client: &client{
			baseUrl:   "https://api.binance.com",
			apiKey:    apiKey,
			apiSecret: apiSecret,
		},
	}
}
