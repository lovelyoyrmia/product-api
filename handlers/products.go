package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"main.go/data"
)


type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products  {
	return &Products{l}
}

type KeyProduct struct {}

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler  {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		err := prod.FromJson(r.Body)

		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(rw, "Unable to unmarshall json", http.StatusBadRequest)
			return
		}

		// validate
		err = prod.Validate()
		if err != nil {
			http.Error(rw, fmt.Sprintf("Error validating product %s", err), http.StatusBadRequest)
			return
		}
		
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)
		next.ServeHTTP(rw, req)
	})
}