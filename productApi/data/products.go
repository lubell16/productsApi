package data

import (
	"fmt"
)

// Product defines the structure for an API product
// swagger:model

type Product struct {
	// the id for the product
	//
	// required:false
	// min:1
	ID int `json:"id"`

	// the name for the product
	//
	// required:true
	// max length: 255
	Name string `json:"name" validate:"required"`

	// the description for this poduct
	//
	// required: false
	// max length: 10000
	Description string `json:"description"`

	// the price for the product
	//
	// required: true
	// min: 0.01
	Price float32 `json:"price" validate:"gt=0"`

	// the SKU for the product
	//
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	SKU string `json:"sku" validate:"required,sku"`
}

//Products is a collection of Product
type Products []*Product

//returns a list of products
func GetProducts() Products {
	return productList
}

func GetProductByID(id int) (*Product, error) {
	pos := findProduct(id)
	if id == -1 {
		return nil, ErrProductNotFound
	}
	return productList[i], nil
}

func AddProduct(p *Product) {
	maxID := productList[len(productList)-1].ID
	p.ID = maxID + 1
	productList = append(productList, p)
}

func UpdateProduct(p Product) error {
	pos := findProduct(p.ID)
	if pos == -1 {
		return ErrProductNotFound
	}
	productList[pos] = &p
	return nil
}

func DeleteProduct(id int) error {
	i := findProduct(id)
	productList = append(productList[:i], productList[i+1])
	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) int {
	for i, p := range productList {
		if p.ID == id {
			return i
		}
	}
	return -1

}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy Mily Coffee",
		Price:       2.45,
		SKU:         "fda-gfd-amo",
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "vwr-ako-dde",
	},
}
