package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/thatbeardo/go-play/data"
)

//	swagger:route PUT /products/{id} products updateProduct
//	Updates a product by id
//
// 	responses:
//		201: noContentResponse
//  	404: errorResponse
//  	422: errorValidation

// UpdateProducts updates products in the underlying storage solution
func (products *Products) UpdateProducts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	newProduct := r.Context().Value(KeyProduct{}).(data.Product)

	err := data.UpdateProduct(id, &newProduct)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
	}

	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
	}
}
