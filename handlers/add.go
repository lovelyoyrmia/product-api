package handlers

import (
	"net/http"

	"main.go/data"
)

func (p *Products) AddProducts(rw http.ResponseWriter, r *http.Request) {
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	data.AddProduct(&prod)
}