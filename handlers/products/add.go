package handlers

import (
	"net/http"

	"github.com/thatbeardo/go-play/data"
)

//	swagger:route POST /products products createProduct
//	Adds a new product to the database
//
// 	responses:
//		200: postOk
//  	400: badRequest

// AddProducts adds products to the underlying storage solution
func (products *Products) AddProducts(w http.ResponseWriter, r *http.Request) {
	products.logger.Println("Adding a product..")
	newProduct := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&newProduct)
}
