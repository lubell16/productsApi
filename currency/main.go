package main

import (
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/lubell16/productsApi/currency/data"
	"github.com/lubell16/productsApi/currency/protos"
	"github.com/lubell16/productsApi/currency/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()

	rates, err := data.NewRates(log)
	if err != nil {
		log.Error("UNable to generate rates", "error", err)
		os.Exit(1)
	}
	// create a new gRPC server, use WithInsecure to allow http connections
	gs := grpc.NewServer()

	// creates an instance of the Currency server
	c := server.NewCurrency(rates, log)

	// register the currency server
	protos.RegisterCurrencyServer(gs, c)

	reflection.Register(gs)
	// create a TCP socket for inbound server connections
	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("Unable to create listener", "error", err)
		os.Exit(1)
	}
	log.Info("Starting Currency Server on port :9092")

	// Listen for requests
	gs.Serve(l)
}
