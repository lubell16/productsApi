package handlers

import (
	"net/http"

	"github.com/lubell16/productsApi/productApi/data"
)

// swagger:route POST /products products createProduct
// Create a new product
//
// responses:
//	200: productResponse

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	p.l.Debug("Inserting product: %#v\n", prod)

	p.productDB.AddProduct(&prod)

}
