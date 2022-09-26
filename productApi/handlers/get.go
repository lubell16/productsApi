package handlers

import (
	"context"
	"net/http"

	"github.com/lubell16/productsApi/currency/protos"
	"github.com/lubell16/productsApi/productApi/data"
)

//  swagger:route GET /products products ListProducts
//  Returns a list of products
//  responses:
//  200: productsResponse

//  GetProducts returns the products from the data store
func (p *Products) ListAll(rw http.ResponseWriter, r *http.Request) {

	p.l.Println("[DEBUG] get all records")
	rw.Header().Add("Content-Type", "application/json")
	// fetch products from the datastore
	lp := data.GetProducts()

	//serialize the list to JSON
	err := data.ToJSON(lp, rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)

	}
}

// ListSingle handles Get Requests
func (p *Products) ListSingle(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	id := getProductID(r)
	p.l.Println("[DEBUG] get record id", id)

	prod, err := data.GetProductByID(id)

	switch err {
	case nil:

	case data.ErrProductNotFound:
		p.l.Println("[ERROR] fetching product", err)

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return

	default:
		p.l.Println("[ERROR] fetching product", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	// get exchange rate
	rr := &protos.RateRequest{
		Base:        protos.Currencies(protos.Currencies_value["EUR"]),
		Destination: protos.Currencies(protos.Currencies_value["GBP"]),
	}
	resp, err := p.cc.GetRate(context.Background(), rr)
	if err != nil {
		p.l.Println("[Error] error getting new rate", err)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	p.l.Printf("Resp %#v", resp)
	prod.Price = prod.Price * resp.Rate

	err = data.ToJSON(prod, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing product", err)
	}
}
