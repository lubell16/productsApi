package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lubell16/productsApi/productApi/data"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Deletes a product from the list
// responses:
// 201: noContentResponse

//  DeleteProducts Deletes a product from the database
func (p *Products) DeleteProducts(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, "UNable to convert ID", http.StatusBadRequest)
		return
	}
	p.l.Debug("[DEBUG] deleting record id", id)

	err = data.DeleteProduct(id)

	if err == data.ErrProductNotFound {
		p.l.Error("[Error] deleting record id does not exist")

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)

		return
	}
	if err != nil {
		p.l.Error("[ERROR] deleting record", "Error", err)
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}
	rw.WriteHeader(http.StatusNoContent)
}
