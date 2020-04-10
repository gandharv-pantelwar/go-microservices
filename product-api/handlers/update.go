package handlers

import (
	"go-microservices/product-api/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// swagger:route POST /products/{id} Products UpdateProduct
// Returns a list of products
// responses:
//  201: noContentResponse

// UpdateProduct updates a product in database
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
