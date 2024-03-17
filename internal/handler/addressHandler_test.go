package handler

import (
	"encoding/json"
	"github.com/brunoliveiradev/GoExpertPostGrad-Challenge/internal/address"
	"github.com/brunoliveiradev/GoExpertPostGrad-Challenge/internal/api/brasilapi"
	"github.com/brunoliveiradev/GoExpertPostGrad-Challenge/internal/api/viacep"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupClient() (*brasilapi.Client, *viacep.Client) {
	return brasilapi.NewClient(http.DefaultClient), viacep.NewClient(http.DefaultClient)
}

func TestGetCepHandler_Success(t *testing.T) {
	bc, vc := setupClient()
	handler := NewAddressHandler(bc, vc)

	req, err := http.NewRequest("GET", "/api/cep/01001000", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	r := mux.NewRouter()
	r.HandleFunc("/api/cep/{cep}", handler.GetCepHandler)
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var resp address.Response
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.NotNil(t, resp.Source)
	assert.NotNil(t, resp.Address)
}

func TestGetCepHandler_InvalidCEP(t *testing.T) {
	bc, vc := setupClient()
	handler := NewAddressHandler(bc, vc)

	req, err := http.NewRequest("GET", "/api/cep/invalidcep", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/cep/{cep}", handler.GetCepHandler)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestGetCepHandler_NotFound(t *testing.T) {
	bc, vc := setupClient()
	handler := NewAddressHandler(bc, vc)

	req, err := http.NewRequest("GET", "/api/cep/89999999", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/cep/{cep}", handler.GetCepHandler)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestGetCepHandler_Timeout(t *testing.T) {
	bc, vc := setupClient()
	handler := NewAddressHandler(bc, vc)

	req, err := http.NewRequest("GET", "/api/cep/12345678", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/cep/{cep}", handler.GetCepHandler)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusRequestTimeout, rr.Code)
}
