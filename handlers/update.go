package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"main.go/data"
)

func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, er := strconv.Atoi(vars["id"])
	if er != nil {
		http.Error(rw, "Cannot convert id", http.StatusBadRequest)
	}
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err := data.UpdateProduct(id, &prod)

	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Something went wrong", http.StatusInternalServerError)
		return
	}
}