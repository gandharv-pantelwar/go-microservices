package handlers

import (
	"go-microservices/product-api/data"
	"net/http"
)

func (p *Products) AddProduct(res http.ResponseWriter, req *http.Request) {
	p.l.Println("Handle POST Products")

	prod := req.Context().Value(KeyProduct{}).(*data.Product)

	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(prod)
}
