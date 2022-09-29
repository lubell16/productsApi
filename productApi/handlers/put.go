package handlers

import (
	"net/http"

	"github.com/lubell16/productsApi/productApi/data"
)

// swagger:route PUT /products products updateProduct
// Update a products details
//
// responses:
//	201: noContentResponse

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	p.l.Debug("[DEBUG] updating record id", prod.ID)

	err := p.productDB.UpdateProduct(prod)

	if err == data.ErrProductNotFound {
		p.l.Error("[ERROR] product not found", err)

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: "Product not found in database"}, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
