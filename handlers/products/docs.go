// Package handlers classification of Product API
//
// Documentation for Product APIS
//
//	Schemes: http
// 		host: localhost:9090
//		BasePath: /
//		Version: 1.0.0
//
//	Consumes:
//		-application/json
//
//	Produces:
//		-application/json
//	swagger:meta
package handlers

import "github.com/thatbeardo/go-play/data"

// Data structure representing a single product
// swagger:response createProduct
type productResponseWrapper struct {
	// Newly created product
	// in: body
	Body data.Product
}

// swagger:parameters updateProduct createProduct
type productParamsWrapper struct {
	// Product data structure to Update or Create.
	// Note: the id field is ignored by update and create operations
	// in: body
	// required: true
	Body data.Product
}

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// swagger:parameters deleteProduct updateProduct
type productIDParamsWrapper struct {
	// The id of the product for which the operation relates
	// in: path
	// required: true
	ID int `json:"id"`
}

// No content Response
// swagger:response noContent
type productsNoContent struct {
	//	No response returned when a product is deleted successfully
	//	in:body
}

// Product with given ID not found
// swagger:response productNotFound
type productNotFound struct {
	// A response denoting the error returned
	// in:body
	Body string
}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}
