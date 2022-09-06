package handlers

import (
	"net/http"

	"github.com/lubell16/working/productApi/data"
)

// swagger:route POST /products products createProduct
// Create a new product
//
// responses:
//	200: productResponse

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)

}
