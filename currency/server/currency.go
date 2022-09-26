package server

import (
	"context"

	"github.com/lubell16/productsApi/currency/protos"

	"github.com/hashicorp/go-hclog"
)

// Currency is a gRPC server it implements the methods defined by the Currency Server interface
type Currency struct {
	log hclog.Logger
}

// NewCurrency creates a new Currency server
func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{l}
}

// GetRate implements the Currency Server GetRate method and returns the currency exchange rate
// for the two given currencies
func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle request for GetRate", "base", rr.GetBase(), "dest", rr.GetDestination())
	return &protos.RateResponse{Rate: 0.5}, nil
}
