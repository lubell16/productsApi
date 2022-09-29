package data

import (
	"testing"

	"github.com/hashicorp/go-hclog"
	"github.com/lubell16/productsApi/currency/protos"
)

func TestNewRates(t *testing.T) {
	l := hclog.Default()

	conn, err := grpc.Dial("localhost:9092", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// creates the client
	cc := protos.NewCurrencyClient(conn)

	// create products DB
	db := data.NewProductsDB(cc, l)

}
