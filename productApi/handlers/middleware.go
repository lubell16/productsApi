package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lubell16/productsApi/productApi/data"
)

// MiddlewareValidateProduct validates the product in the request and calls next if ok

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
		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)

		// Call thenext handler, which can be another midleware in the chain, or the final handler
		next.ServeHTTP(rw, req)
	})
}
