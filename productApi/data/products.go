package data

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-hclog"
	"github.com/lubell16/productsApi/currency/protos"
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
	Price float64 `json:"price" validate:"gt=0"`

	// the SKU for the product
	//
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	SKU string `json:"sku" validate:"required,sku"`
}

//Products is a collection of Product
type Products []*Product

type ProductsDB struct {
	currency protos.CurrencyClient
	log      hclog.Logger
}

func NewProductsDB(c protos.CurrencyClient, l hclog.Logger) *ProductsDB {
	return &ProductsDB{c, l}
}

//returns a list of products
func (p *ProductsDB) GetProducts(currency string) (Products, error) {
	if currency == "" {
		return productList, nil
	}
	rate, err := p.getRate(currency)
	if err != nil {
		p.log.Error("Unable to get rate", "currency", err)
		return nil, err
	}
	pr := Products{}
	for _, p := range productList {
		np := *p
		np.Price = np.Price * rate
		pr = append(pr, &np)
	}
	return pr, nil
}

func (p *ProductsDB) GetProductByID(id int, currency string) (*Product, error) {
	i := findIndexByProductID(id)
	if id == -1 {
		return nil, ErrProductNotFound
	}
	if currency == "" {
		return productList[i], nil
	}
	rate, err := p.getRate(currency)
	if err != nil {
		p.log.Error("Unable to get rate", "currency", err)
		return nil, err
	}
	np := *productList[i]
	np.Price = np.Price * rate

	return &np, nil
}

func AddProduct(p *Product) {
	maxID := productList[len(productList)-1].ID
	p.ID = maxID + 1
	productList = append(productList, p)
}

func (p *ProductsDB) UpdateProduct(pr Product) error {
	i := findIndexByProductID(pr.ID)
	if i == -1 {
		return ErrProductNotFound
	}
	productList[i] = &pr
	return nil
}

func DeleteProduct(id int) error {
	i := findIndexByProductID(id)
	productList = append(productList[:i], productList[i+1])
	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findIndexByProductID(id int) int {
	for i, p := range productList {
		if p.ID == id {
			return i
		}
	}
	return -1

}

func (p *ProductsDB) getRate(destination string) (float64, error) {
	rr := &protos.RateRequest{
		Base:        protos.Currencies(protos.Currencies_value["EUR"]),
		Destination: protos.Currencies(protos.Currencies_value["GBP"]),
	}
	resp, err := p.currency.GetRate(context.Background(), rr)
	return resp.Rate, err
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
