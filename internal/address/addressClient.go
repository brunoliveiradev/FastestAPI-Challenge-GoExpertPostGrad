package address

import (
	"context"
	"github.com/brunoliveiradev/GoExpertPostGrad-Challenge/pkg/model"
	"net/http"
)

type Response struct {
	Source  string         `json:"source"`
	Address *model.Address `json:"address"`
	Error   *string        `json:"error,omitempty"`
}

type Service interface {
	GetAddress(ctx context.Context, cep string) (*Response, error)
}

type Client struct {
	HTTPClient *http.Client
	BaseURL    string
}
