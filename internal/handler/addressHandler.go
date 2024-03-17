package handler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/brunoliveiradev/GoExpertPostGrad-Challenge/internal/address"
	"github.com/brunoliveiradev/GoExpertPostGrad-Challenge/internal/api/brasilapi"
	"github.com/brunoliveiradev/GoExpertPostGrad-Challenge/internal/api/viacep"
	"github.com/brunoliveiradev/GoExpertPostGrad-Challenge/util"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"regexp"
	"time"
)

type AddressHandler struct {
	brasilAPIClient *brasilapi.Client
	viaCepClient    *viacep.Client
}

func NewAddressHandler(bc *brasilapi.Client, vc *viacep.Client) *AddressHandler {
	return &AddressHandler{
		brasilAPIClient: bc,
		viaCepClient:    vc,
	}
}

func (h *AddressHandler) GetCepHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cep := vars["cep"] // Get the CEP from the URL path

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if !IsValidCEP(cep) {
		http.Error(w, "A valid CEP is required", http.StatusBadRequest)
		return
	}

	log.Println("[INFO] Received request for CEP:", cep)

	responseChan := make(chan *address.Response, 2) // channel to receive the response from the services
	errorChan := make(chan *util.CustomError, 2)    // channel to receive the errors from the services

	// Call the both services API using Go routines and channels and retuning the fastest response
	go h.callGetAddressAndSendResponse(h.brasilAPIClient, cep, ctx, responseChan, errorChan)
	go h.callGetAddressAndSendResponse(h.viaCepClient, cep, ctx, responseChan, errorChan)

	// Wait for the responses and errors from the services
	resp, err := h.waitForResponses(ctx, responseChan, errorChan)
	if err != nil {
		http.Error(w, err.Message, err.Status)
		return
	}

	respJSON, marshalErr := json.Marshal(resp)
	if marshalErr != nil {
		log.Printf("[DEBUG] Failed to serialize response: %v\n", marshalErr)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	log.Printf("[INFO] Response: %s\n", respJSON)

	w.Header().Set("Content-Type", "application/json")
	w.Write(respJSON)
}

func (h *AddressHandler) callGetAddressAndSendResponse(s address.Service, cep string, ctx context.Context, responseChan chan<- *address.Response, errorChan chan<- *util.CustomError) {
	resp, err := s.GetAddress(ctx, cep) // Call the service API
	if err != nil {
		var customErr *util.CustomError
		ok := errors.As(err, &customErr)
		if !ok {
			customErr = &util.CustomError{Message: "Unexpected error", Status: http.StatusInternalServerError}
		}
		errorChan <- customErr
		return
	}
	responseChan <- resp
}

func (h *AddressHandler) waitForResponses(ctx context.Context, responseChan <-chan *address.Response, errorChan <-chan *util.CustomError) (*address.Response, *util.CustomError) {
	var notFoundErrors int
	for i := 0; i < 2; i++ {
		select {
		case resp := <-responseChan:
			if resp.Error == nil {
				return resp, nil // Return the first valid response without errors
			}
		case err := <-errorChan:
			if err.Status == http.StatusNotFound {
				notFoundErrors++
			} else {
				return nil, err // Returns the first error found (not a 404)
			}
		case <-ctx.Done():
			return nil, &util.CustomError{Status: http.StatusRequestTimeout, Message: "Timeout reached!"}
		}
	}

	if notFoundErrors == 2 {
		log.Println("[INFO] CEP not found!")
		return nil, &util.CustomError{Status: http.StatusNotFound, Message: "CEP not found!"}
	}

	log.Println("[DEBUG] Unexpected error!")
	return nil, &util.CustomError{Status: http.StatusInternalServerError, Message: "Unexpected error!"}
}

// IsValidCEP checks if a CEP is valid by matching it against a regular expression
func IsValidCEP(cep string) bool {
	re := regexp.MustCompile(`^\d{8}$`)
	return re.MatchString(cep)
}
