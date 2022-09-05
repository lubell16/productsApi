// Package classification of Product API
//
// Documentation for Product API
//
// Scemes:http
// BasePath:/
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/lubell16/working/data"
)

//  A list of products returns in response
//  swagger:response productsResponse
type productsResponse struct {
	// All products in the system
	// in: body
	Body []data.Product
}

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

type KeyProduct struct{}

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}
		//validating json
		err := prod.FromJSON(r.Body)

		if err != nil {
			p.l.Println("[ERROR] deserializing product")
			http.Error(rw, "Error reading product", http.StatusBadRequest)
			return
		}
		//validating product
		err = prod.Validate()

		if err != nil {
			p.l.Println("[ERROR] validating product")
			http.Error(
				rw,
				fmt.Sprintf("Error reading product: %s", err),
				http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)
		next.ServeHTTP(rw, req)
	})
}
