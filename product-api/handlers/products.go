package handlers

import (
	"context"
	"go-microservices/product-api/data"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Products ...
type Products struct {
	l *log.Logger
}

// NewProducts ...
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(res http.ResponseWriter, req *http.Request) {
	p.l.Println("Handle GET Products")
	productList := data.GetProducts()
	err := productList.ToJSON(res)
	// d, err := json.Marshal(productList)
	if err != nil {
		http.Error(res, "Unable to parse json", http.StatusInternalServerError)
	}
	// res.Write(d)
}

func (p *Products) AddProduct(res http.ResponseWriter, req *http.Request) {
	p.l.Println("Handle POST Products")

	prod := req.Context().Value(KeyProduct{}).(*data.Product)

	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(prod)
}

func (p *Products) UpdateProduct(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	idString := vars["id"]
	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(res, "Invalid id", http.StatusBadRequest)
	}
	p.l.Println("Handle PUT Products", id)

	prod := req.Context().Value(KeyProduct{}).(*data.Product)

	p.l.Printf("Prod: %#v", prod)
	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(res, "Product not found", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(res, "Product not found", http.StatusBadRequest)
	}
}

type KeyProduct struct{}

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
