package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	gohandlers "github.com/gorilla/handlers"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/lubell16/productsApi/product-images/files"
	"github.com/lubell16/productsApi/product-images/handlers"

	// @ts-ignore
	"github.com/nicholasjackson/env"
)

var basePath = env.String("BASE_PATH", false, "./imagestore", "Base path to save images")
var bindAddress = env.String("BIND_ADDRESS", false, ":9091", "Bind address for the server")
var logLevel = env.String("LOG_LEVEL", false, "debug", "Log output level for the server [debug, info, trace]")

//var basePath = env.String("BASE_PATH", false, "./imagestore", "Base path to save images")

func main() {
	env.Parse()

	l := hclog.New(
		&hclog.LoggerOptions{
			Name:  "product-images",
			Level: hclog.LevelFromString(*logLevel),
		},
	)

	sl := l.StandardLogger(&hclog.StandardLoggerOptions{InferLevels: true})
	//create a logger for the server from the default logger

	//create the storage class, use local storage
	stor, err := files.NewLocal(*basePath, 1024*1000*5)
	if err != nil {
		l.Error("Unable to create storage", "error", err)
		os.Exit(1)
	}
	//create the handlers
	fh := handlers.NewFiles(stor, l)
	mw := handlers.GzipHandler{}

	//create a new serve mux and register handlers
	sm := mux.NewRouter()
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	ph := sm.Methods(http.MethodPost).Subrouter()
	ph.HandleFunc("/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}", fh.UploadREST)
	ph.HandleFunc("/", fh.UploadMultipart)

	// get files
	gh := sm.Methods(http.MethodGet).Subrouter()
	gh.Handle(
		"/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}",
		http.StripPrefix("/images/", http.FileServer(http.Dir(*basePath))))
	gh.Use(mw.GzipMiddleware)
	//create a new server
	s := &http.Server{
		Addr:         *bindAddress,
		Handler:      ch(sm),
		ErrorLog:     sl,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		l.Info("Starting server", "bind_address", *bindAddress)

		err := s.ListenAndServe()
		if err != nil {
			l.Error("Unable to start the server", "error", err)
			os.Exit(1)
		}
	}()

	//trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	//Block until a signal is received
	sig := <-c
	l.Info("Shutting down server with", "signal", sig)

	//gracefully shut down server, waiting 30 seconds for current operations to complete
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
