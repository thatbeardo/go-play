package handlers

import (
	"context"
	"net/http"

	"github.com/thatbeardo/go-play/data"
)

// KeyProduct is used in the middleware
type KeyProduct struct{}

// ProductValidationMiddleware to carry out product validation before PUT and POST reqs
func (products *Products) ProductValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		newProduct := &data.Product{}
		err := newProduct.FromJSON(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = newProduct.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, *newProduct)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
