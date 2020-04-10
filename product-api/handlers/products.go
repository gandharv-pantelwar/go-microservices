package handlers

import (
	"context"
	"go-microservices/product-api/data"
	"log"
	"net/http"
)

// Products ...
type Products struct {
	l *log.Logger
}

// NewProducts ...
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

type KeyProduct struct{}

// MiddlewareProductValidation ...
func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		prod := &data.Product{}
		err := prod.FromJSON(req.Body)
		if err != nil {
			http.Error(res, "Unable to marshal json", http.StatusBadRequest)
			return
		}

		err = prod.Validate()
		if err != nil {
			http.Error(res, "Invalid params", http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(req.Context(), KeyProduct{}, prod)
		req = req.WithContext(ctx)
		next.ServeHTTP(res, req)
	})
}
