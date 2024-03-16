package model

type Address struct {
	CEP               string  `json:"cep"`
	State             string  `json:"state"`
	City              string  `json:"city"`
	Neighborhood      string  `json:"neighborhood"`
	Street            string  `json:"street"`
	AdditionalDetails *string `json:"additional_details,omitempty"`
	IBGE              *string `json:"ibge,omitempty"`
	GIA               *string `json:"gia,omitempty"`
	Service           *string `json:"service,omitempty"`
	DDD               *string `json:"ddd,omitempty"`
	SIAFI             *string `json:"siafi,omitempty"`
}
