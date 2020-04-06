package handlers

import (
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

func (p *Products) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		p.getProducts(res, req)
		return
	}
	res.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(res http.ResponseWriter, req *http.Request) {
	productList := data.GetProducts()
	err := productList.ToJSON(res)
	// d, err := json.Marshal(productList)
	if err != nil {
		http.Error(res, "Unable to parse json", http.StatusInternalServerError)
	}
	// res.Write(d)
}
