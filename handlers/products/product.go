package handlers

import "log"

// Products represents the handler for the products endpoint
type Products struct {
	logger *log.Logger
}

// NewProducts accepts a logger and returns an instance of the product handler
func NewProducts(logger *log.Logger) *Products {
	return &Products{logger}
}
