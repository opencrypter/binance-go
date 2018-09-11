package binance

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Sdk struct {
	client Client
	clock  Clock
}

type Clock interface {
	Now() *int64
}

type clock struct {
}

func (t clock) Now() *int64 {
	now := int64(time.Nanosecond) * time.Now().UnixNano() / int64(time.Millisecond)
	return &now
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
	signed     bool
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

func (r *request) StringParam(key string, value *string) *request {
	if value != nil {
		r.parameters.Set(key, *value)
	}
	return r
}

func (r *request) Float64Param(key string, value *float64) *request {
	if value != nil {
		r.parameters.Set(key, strconv.FormatFloat(*value, 'f', 8, 64))
	}
	return r
}

func (r *request) Int64Param(key string, value *int64) *request {
	if value != nil {
		r.parameters.Set(key, strconv.FormatInt(*value, 10))
	}
	return r
}

func (r *request) Sign() *request {
	r.signed = true
	return r
}

func (c *client) Do(request *request) ([]byte, error) {
	r, _ := c.createRequest(request)
	c.addSignature(request, r)
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

func (c *client) createRequest(request *request) (*http.Request, error) {
	if request.method == "GET" {
		return c.get(request)
	}
	return c.post(request)
}

func (c *client) sign(request *request) string {
	signature := hmac.New(sha256.New, []byte(c.apiSecret))
	signature.Write([]byte(request.parameters.Encode()))

	return hex.EncodeToString(signature.Sum(nil))
}

func (c *client) addSignature(request *request, httpRequest *http.Request) {
	if request.signed {
		query := httpRequest.URL.Query()
		httpRequest.URL.RawQuery = query.Encode() + "&signature=" + c.sign(request)
	}
}

func (c *client) get(request *request) (*http.Request, error) {
	parameters := request.parameters.Encode()
	return http.NewRequest(request.method, c.baseUrl+request.path + "?" + parameters, nil)
}

func (c *client) post(request *request) (*http.Request, error) {
	form := request.parameters.Encode()
	return http.NewRequest(request.method, c.baseUrl+request.path, strings.NewReader(form))
}

func New(apiKey string, apiSecret string) Sdk {
	return Sdk{
		client: &client{
			baseUrl:   "https://api.binance.com",
			apiKey:    apiKey,
			apiSecret: apiSecret,
		},
		clock: clock{},
	}
}
