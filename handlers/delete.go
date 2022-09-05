package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lubell16/working/data"
)

func (p *Products) DeleteProducts(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, "UNable to convert ID", http.StatusBadRequest)
		return
	}
	p.l.Println("[DEBUG] deleting record id", id)

	err = data.DeleteProduct(id)

	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
