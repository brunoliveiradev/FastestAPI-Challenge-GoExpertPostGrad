package viacep

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

// setupTestServer configures a new httptest.Server and returns a new Client using the server URL
func setupTestServer(response string, statusCode int) (*Client, func(), error) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(statusCode)
		rw.Write([]byte(response))
	}))

	client := NewClient(server.Client())
	client.BaseURL = server.URL + "/"

	teardown := func() { server.Close() }

	return client, teardown, nil
}

func TestViaCepGetAddressSuccess(t *testing.T) {
	response := `{"cep": "01001000", "logradouro": "Praça da Sé", "complemento": "", "bairro": "Sé", "localidade": "São Paulo", "uf": "SP", "ibge": "3550308", "gia": "1004", "ddd": "11", "siafi": "7107"}`
	client, teardown, err := setupTestServer(response, http.StatusOK)
	assert.NoError(t, err)
	defer teardown()

	resp, err := client.GetAddress(context.Background(), "01001-000")

	assert.NoError(t, err)
	assert.Equal(t, "ViaCep API", resp.Source)
	assert.Equal(t, "01001000", resp.Address.CEP)
}

func TestViaCepGetAddressNotFound(t *testing.T) {
	response := `{"erro": true}`
	client, teardown, err := setupTestServer(response, http.StatusOK) // ViaCep returns 200 OK with response body {"erro": true} for not found CEPs
	assert.NoError(t, err)
	defer teardown()

	_, err = client.GetAddress(context.Background(), "99999999")

	assert.Error(t, err)
	var customErr *util.CustomError
	assert.True(t, errors.As(err, &customErr))
	assert.Equal(t, http.StatusNotFound, customErr.Status)
	assert.Equal(t, "CEP not found on viaCep API!", customErr.Message)
}

func TestViaCepGetAddressServerError(t *testing.T) {
	client, teardown, err := setupTestServer("", http.StatusInternalServerError)
	assert.NoError(t, err)
	defer teardown()

	_, err = client.GetAddress(context.Background(), "01001-000")

	assert.Error(t, err)
	var customErr *util.CustomError
	assert.True(t, errors.As(err, &customErr))
	assert.Equal(t, http.StatusInternalServerError, customErr.Status)
	assert.Equal(t, "error on request to viaCep API", customErr.Message)
}

func TestRealViaCepGetAddressIntegration(t *testing.T) {
	httpClient := &http.Client{Timeout: 5 * time.Second}
	client := NewClient(httpClient)

	cep := "01001-000"

	resp, err := client.GetAddress(context.Background(), cep)

	// Verifique se não há erro na chamada da API
	assert.NoError(t, err)
	assert.Equal(t, "ViaCep API", resp.Source)
	assert.Equal(t, cep, resp.Address.CEP)
	assert.Equal(t, "SP", resp.Address.State)
	assert.Equal(t, "São Paulo", resp.Address.City)
	assert.Equal(t, "Sé", resp.Address.Neighborhood)
}

func TestRealViaCepGetAddressNotFoundIntegration(t *testing.T) {
	httpClient := &http.Client{Timeout: 7 * time.Second}
	client := NewClient(httpClient)

	cep := "99999999" // invalid CEP

	_, err := client.GetAddress(context.Background(), cep)

	assert.Error(t, err, "An error was expected when querying a nonexistent CEP")

	var customErr *util.CustomError
	if assert.True(t, errors.As(err, &customErr), "The error should be of type *util.CustomError") {
		assert.Equal(t, http.StatusNotFound, customErr.Status, "The HTTP status code for a nonexistent CEP should be 404")
		assert.Equal(t, "CEP not found on viaCep API!", customErr.Message, "The error message did not match the expected one")
	}
}
