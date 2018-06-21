package binance

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func runLocalServer() (server *httptest.Server, mux *http.ServeMux) {
	mux = http.NewServeMux()
	srv := httptest.NewServer(mux)

	return srv, mux
}

func TestClient_Get(t *testing.T) {
	server, mux := runLocalServer()
	defer server.Close()

	sdk := client{
		baseUrl:   server.URL,
		apiKey:    "key",
		apiSecret: "secret",
	}

	t.Run("It should return the server response", func(t *testing.T) {
		expectedSuccessResponse := []byte(`{"json":true}`)
		mux.HandleFunc("/testing", func(w http.ResponseWriter, r *http.Request) {
			w.Write(expectedSuccessResponse)
			w.WriteHeader(200)
		})

		response, _ := sdk.Get("/testing")
		assert.Equal(t, expectedSuccessResponse, response)
	})

	t.Run("It should return an error on receive an invalid path", func(t *testing.T) {
		_, err := sdk.Get("error")
		assert.Error(t, err)
	})

	t.Run("It should return an error on server error", func(t *testing.T) {
		_, err := sdk.Get("/missing-path")
		assert.Error(t, err)
	})
}

func TestNew(t *testing.T) {
	sdk := New("api-key", "api-secret")
	assert.Implements(t, (*Client)(nil), sdk.client)
}
