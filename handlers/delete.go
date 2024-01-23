package handlers

import (
	"net/http"

	"main.go/data"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Delete a product
// responses:
// 	201: noContent

// DeleteProduct deletes a product from the database
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err := data.DeleteProduct(id, &prod)

	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Something went wrong", http.StatusInternalServerError)
		return
	}
}