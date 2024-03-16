package brasilapi

import (
	"context"
	"errors"
	"github.com/brunoliveiradev/GoExpertPostGrad-Challenge/util"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func setupTestServer(response string, statusCode int) (*Client, func(), error) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(statusCode)
		rw.Write([]byte(response))
	}))

	// Create a new Client using the server URL
	client := NewClient(server.Client())
	client.BaseURL = server.URL + "/"

	// Teardown function to close the server
	teardown := func() { server.Close() }

	return client, teardown, nil
}

func TestGetAddressReturnsAddressOnSuccess(t *testing.T) {
	response := `{"cep": "01001-000", "state": "SP", "city": "São Paulo", "neighborhood": "Sé", "street": "Praça da Sé", "service": "Correios"}`
	client, teardown, err := setupTestServer(response, http.StatusOK)
	assert.NoError(t, err)
	defer teardown()

	resp, err := client.GetAddress(context.Background(), "01001000")

	assert.NoError(t, err)
	assert.Equal(t, "BrasilAPI", resp.Source)
	assert.Equal(t, "01001-000", resp.Address.CEP)
	assert.Equal(t, "01001-000", resp.Address.CEP)
}

func TestGetAddressReturnsErrorOnNotFound(t *testing.T) {
	client, teardown, err := setupTestServer("", http.StatusNotFound)
	assert.NoError(t, err)
	defer teardown()

	_, err = client.GetAddress(context.Background(), "99999999")

	assert.Error(t, err)
	var customErr *util.CustomError
	assert.True(t, errors.As(err, &customErr))
	assert.Equal(t, http.StatusNotFound, customErr.Status)
	assert.Equal(t, "CEP not found on BrasilAPI!", customErr.Message)
}

func TestGetAddressReturnsErrorOnServerError(t *testing.T) {
	client, teardown, err := setupTestServer("", http.StatusInternalServerError)
	assert.NoError(t, err)
	defer teardown()

	_, err = client.GetAddress(context.Background(), "01001000")

	assert.Error(t, err)
	var customErr *util.CustomError
	assert.True(t, errors.As(err, &customErr))
	assert.Equal(t, http.StatusInternalServerError, customErr.Status)
	assert.Equal(t, "error on request to BrasilAPI", customErr.Message)
}

func TestRealGetAddressIntegration(t *testing.T) {
	httpClient := &http.Client{Timeout: 4 * time.Second}
	client := NewClient(httpClient)

	cep := "89010025"

	resp, err := client.GetAddress(context.Background(), cep)

	assert.NoError(t, err)
	assert.Equal(t, cep, resp.Address.CEP)
}

func TestGetAddressWithNonexistentCEP(t *testing.T) {
	httpClient := &http.Client{Timeout: 4 * time.Second}
	client := NewClient(httpClient)

	cep := "89999999"

	_, err := client.GetAddress(context.Background(), cep)

	assert.Error(t, err)

	var customErr *util.CustomError
	assert.True(t, errors.As(err, &customErr))
	assert.Equal(t, http.StatusNotFound, customErr.Status)
	assert.Equal(t, "CEP not found on BrasilAPI!", customErr.Message)
}
