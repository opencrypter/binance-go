package binance

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Sdk struct {
	client Client
}

type Client interface {
	Do(request *request) ([]byte, error)
}

type client struct {
	baseUrl   string
	apiKey    string
	apiSecret string
}

type request struct {
	method     string
	path       string
	parameters url.Values
	signature  string
}

func newRequest(method string, path string) *request {
	return &request{
		method:     method,
		path:       path,
		parameters: url.Values{},
	}
}

func (r *request) Param(key string, value string) *request {
	r.parameters.Set(key, value)
	return r
}

func (c *client) Do(request *request) ([]byte, error) {
	parameters := request.parameters.Encode()

	r, _ := http.NewRequest("GET", c.baseUrl+request.path+parameters, nil)
	r.Header.Set("X-MBX-APIKEY", c.apiKey)

	client := http.Client{}
	response, err := client.Do(r)

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
