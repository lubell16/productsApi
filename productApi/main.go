package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	gohandlers "github.com/gorilla/handlers"
	"github.com/hashicorp/go-hclog"
	protos "github.com/lubell16/productsApi/currency/protos"
	"google.golang.org/grpc"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/lubell16/productsApi/productApi/data"
	"github.com/lubell16/productsApi/productApi/handlers"
)

func main() {

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

	// create the handlers
	ph := handlers.NewProducts(l, db)

	// creates the mux
	sm := mux.NewRouter()

	// handlers for api
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", ph.ListAll).Queries("currency", "{[A-Z]{3}}")
	getRouter.HandleFunc("/products", ph.ListAll)

	getRouter.HandleFunc("/products/{id:[0-9]+}", ph.ListSingle).Queries("currency", "{[A-Z]{3}}")
	getRouter.HandleFunc("/products/{id:[0-9]+}", ph.ListSingle)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProduct)
	putRouter.Use(ph.MiddlewareProductValidation)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)
	postRouter.Use(ph.MiddlewareProductValidation)

	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/{id:[0-9]+}", ph.DeleteProducts)

	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)
	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))
	s := &http.Server{
		Addr:         ":9090",
		Handler:      ch(sm),
		ErrorLog:     l.StandardLogger(&hclog.StandardLoggerOptions{}),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		l.Info("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Error("Error starting server", "error", err)
			os.Exit(1)
		}
	}()
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Info("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
