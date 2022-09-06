package handlers

import (
	"net/http"

	"github.com/lubell16/working/productApi/data"
)

//  swagger:route GET /products products ListProducts
//  Returns a list of products
//  responses:
//  200: productsResponse

//  GetProducts returns the products from the data store
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// fetch products from the datastore
	lp := data.GetProducts()

	//serialize the list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)

	}
}