//go:generate swagger generate spec -o ./swagger.yml

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"main.go/handlers"
)

func main() {
	
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	ph := handlers.NewProducts(l)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	// sm.Use()
	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts)
	putRouter.Use(ph.MiddlewareProductValidation)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProducts)
	postRouter.Use(ph.MiddlewareProductValidation)

	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/{id:[0-9]+}", ph.DeleteProduct)
 
	ops := middleware.RedocOpts{SpecURL: "/swagger.json"}
	sh := middleware.Redoc(ops, nil) 
	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.json", http.FileServer(http.Dir("./")))

	bindAddress := ":9090"

	s := http.Server{
		Addr: bindAddress,
		Handler: sm,
		ErrorLog: l,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout: 5 * time.Second,
	}

	go func ()  {
		l.Println("Starting server on port 9090")
		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server : %s\n", err)
			os.Exit(1)
		}	
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan

	l.Println("Receive terminate, graceful shutdown", sig)

	ts := time.Second * 30
	tc, _ := context.WithTimeout(context.Background(), ts)

	s.Shutdown(tc)
}