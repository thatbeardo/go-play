package handlers

import (
	"net/http"

	"github.com/thatbeardo/go-play/data"
)

//	swagger:route GET /products products listProducts
//	Returns a list of products
//	responses:
//		200: productsResponseWrapper

// GetProducts returns a list of products present in the underlying storage
func (products *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	productList := data.GetProducts()
	err := productList.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}
