package handlers

import (
	"net/http"

	"github.com/lubell16/productsApi/productApi/data"
)

//  swagger:route GET /products products ListProducts
//  Returns a list of products
//  responses:
//  200: productsResponse

func (p *Products) ListAll(rw http.ResponseWriter, r *http.Request) {
	p.l.Debug("Get all records")
	rw.Header().Add("Content-Type", "application/json")

	cur := r.URL.Query().Get("currency")

	prods, err := p.productDB.GetProducts(cur)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	err = data.ToJSON(prods, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Error("Unable to serializing product", "error", err)
	}
}

// ListSingle handles Get Requests
func (p *Products) ListSingle(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	id := getProductID(r)
	cur := r.URL.Query().Get("currency")
	p.l.Debug("[DEBUG] get record id", "id", id)

	prod, err := p.productDB.GetProductByID(id, cur)

	switch err {
	case nil:

	case data.ErrProductNotFound:
		p.l.Error("[ERROR] fetching product", err)

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return

	default:
		p.l.Error("[ERROR] fetching product", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	err = data.ToJSON(prod, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Error("[ERROR] serializing product", err)
	}
}
