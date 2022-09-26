package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lubell16/working/currency/protos"
)

type Products struct {
	l  *log.Logger
	cc protos.CurrencyClient
}

func NewProducts(l *log.Logger, cc protos.CurrencyClient) *Products {
	return &Products{l, cc}
}

type KeyProduct struct{}

// ErrInvalidProductPath is an error message when the product path is not valid
var ErrInvalidProductPath = fmt.Errorf("Invalid Path, path should be /products/[id]")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

// getProductID returns the product ID from the URL
// Panics fi cannot convert the id into an integer
// this should never happen as the router ensures that
// this is a valid numbner

func getProductID(r *http.Request) int {
	// parse the product id from the url
	vars := mux.Vars(r)

	//convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// this should never happen
		panic(err)

	}
	return id
}
