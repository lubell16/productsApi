package main

import (
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	protos "github.com/lubell16/working/currency/protos"
	"github.com/lubell16/working/currency/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()

	// create a new gRPC server, use WithInsecure to allow http connections
	gs := grpc.NewServer()

	// creates an instance of the Currency server
	c := server.NewCurrency(log)

	// register the currency server
	protos.RegisterCurrencyServer(gs, c)

	reflection.Register(gs)
	// create a TCP socket for inbound server connections
	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("Unable to create listener", "error", err)
		os.Exit(1)
	}
	// Listen for requests
	gs.Serve(l)
}
