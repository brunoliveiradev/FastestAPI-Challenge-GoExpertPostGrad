package brasilapi

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/brunoliveiradev/GoExpertPostGrad-Challenge/internal/address"
	"github.com/brunoliveiradev/GoExpertPostGrad-Challenge/pkg/model"
	"github.com/brunoliveiradev/GoExpertPostGrad-Challenge/util"
	"net/http"
)

type brasilApiCep struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

type Client struct {
	HTTPClient *http.Client
	BaseURL    string
}

// NewClient creates a new BrasilAPI client.
func NewClient(httpClient *http.Client) *Client {
	return &Client{
		HTTPClient: httpClient,
		BaseURL:    "https://brasilapi.com.br/api/cep/v1/",
	}
}

func (bc *Client) GetAddress(ctx context.Context, cep string) (*address.Response, error) {
	url := fmt.Sprintf("%s%s", bc.BaseURL, cep)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return nil, &util.CustomError{Status: http.StatusNotFound, Message: "CEP not found on BrasilAPI!"}
		}
		return nil, &util.CustomError{Status: resp.StatusCode, Message: "error on request to BrasilAPI"}
	}

	var brasilApiCep brasilApiCep
	if err := json.NewDecoder(resp.Body).Decode(&brasilApiCep); err != nil {
		return nil, err
	}

	return &address.Response{
		Source:  "BrasilAPI",
		Address: toAddressModel(&brasilApiCep),
	}, nil
}

func toAddressModel(brasilApiCep *brasilApiCep) *model.Address {
	return &model.Address{
		CEP:          brasilApiCep.Cep,
		State:        brasilApiCep.State,
		City:         brasilApiCep.City,
		Neighborhood: brasilApiCep.Neighborhood,
		Street:       brasilApiCep.Street,
		Service:      &brasilApiCep.Service,
	}
}
