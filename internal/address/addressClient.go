package address

import (
	"context"
	"github.com/brunoliveiradev/GoExpertPostGrad-Challenge/pkg/model"
	"net/http"
)

// Response is the data structure that represents an address response.
// @Description Data structure representing an address response.
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
