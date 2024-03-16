package viacep

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/brunoliveiradev/GoExpertPostGrad-Challenge/internal/address"
	"github.com/brunoliveiradev/GoExpertPostGrad-Challenge/pkg/model"
	"github.com/brunoliveiradev/GoExpertPostGrad-Challenge/util"
	"io"
	"log"
	"net/http"
)

type viaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type viaCepResponseErrorNotFound struct {
	Erro bool `json:"erro"`
}

type Client struct {
	HTTPClient *http.Client
	BaseURL    string
}

// NewClient creates a new viaCep client.
func NewClient(httpClient *http.Client) *Client {
	return &Client{
		HTTPClient: httpClient,
		BaseURL:    "https://viacep.com.br/ws/",
	}
}

// GetAddress make a request to viaCep  API and returns the address.
func (vc *Client) GetAddress(ctx context.Context, cep string) (*address.Response, error) {
	url := fmt.Sprintf("%s%s/json/", vc.BaseURL, cep)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	log.Println("[INFO] Request to viaCep API: ", url)
	resp, err := vc.HTTPClient.Do(req)
	if err != nil {
		log.Println("[ERROR] Request to viaCep API failed:", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("[ERROR] error on request to viaCep: status code", resp.StatusCode)
		return nil, &util.CustomError{Status: resp.StatusCode, Message: "error on request to viaCep API"}
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("[ERROR] Reading response body failed:", err)
		return nil, err
	}

	var errorResponse viaCepResponseErrorNotFound
	if err := json.Unmarshal(bodyBytes, &errorResponse); err == nil && errorResponse.Erro {
		return nil, &util.CustomError{Status: http.StatusNotFound, Message: "CEP not found on viaCep API!"}
	}

	var viaCEP viaCep
	if err := json.Unmarshal(bodyBytes, &viaCEP); err != nil {
		log.Println("[ERROR] Decoding viaCep response failed:", err)
		return nil, err
	}

	return &address.Response{
		Source:  "ViaCep API",
		Address: toAddressModel(&viaCEP),
	}, nil
}

func toAddressModel(viaCep *viaCep) *model.Address {
	return &model.Address{
		CEP:               viaCep.Cep,
		State:             viaCep.Uf,
		City:              viaCep.Localidade,
		Neighborhood:      viaCep.Bairro,
		Street:            viaCep.Logradouro,
		AdditionalDetails: &viaCep.Complemento,
		IBGE:              &viaCep.Ibge,
		GIA:               &viaCep.Gia,
		DDD:               &viaCep.Ddd,
		SIAFI:             &viaCep.Siafi,
	}
}
