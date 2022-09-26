// Package classification of Product API
//
// Documentation for Product API
//
// Scemes:http
// BasePath:/
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta
package handlers

import (
	"github.com/lubell16/productsApi/productApi/data"
)

//  A list of products returns in response
//  swagger:response productsResponse
type productsResponseWrapper struct {
	//  All products in the system
	//  in: body
	Body []data.Product
}

// Data structure representing a single product
// swagger:response productResponse
type productResponseWrapper struct {
	// Newly created product
	// in: body
	Body data.Product
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// swagger:parameters deleteProduct
type productIDParameterWrapper struct {
	// The id of the product for which the operation relates
	// in: path
	// required: true
	ID int `json:"id"`
}
