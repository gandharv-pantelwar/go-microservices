package handlers

import (
	"go-microservices/product-api/data"
	"net/http"
)

// swagger:route GET /products Products GetProducts
// Returns a list of productss
// responses:
// 200: productsResponse

// GetProducts returns the products from the data store
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
