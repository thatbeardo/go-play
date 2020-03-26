package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/thatbeardo/go-play/data"
)

//	swagger:route DELETE /products/{id} products deleteProduct
//	Update a products details
//
//	responses:
//		201: noContentResponse
//  	404: errorResponse
//  	501: errorResponse

// DeleteProduct deletes a product from the productList
func (product *Products) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	err := data.DeleteProduct(id)
	if err != nil || err == data.ErrProductNotFound {
		http.Error(w, "The product does not exist", http.StatusNotFound)
	}
}
