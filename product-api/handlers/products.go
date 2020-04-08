package handlers

import (
	"go-microservices/product-api/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
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
	} else if req.Method == http.MethodPost {
		p.addProduct(res, req)
		return
	} else if req.Method == http.MethodPut {
		r := regexp.MustCompile(`/([0-9]+)`)
		g := r.FindAllStringSubmatch(req.URL.Path, -1)
		p.l.Println("g:", g)
		if len(g) != 1 {
			http.Error(res, "Invalid URI", http.StatusBadRequest)
			return
		}
		if len(g[0]) != 2 {
			http.Error(res, "Invalid URI", http.StatusBadRequest)
			return
		}
		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(res, "Invalid URI", http.StatusBadRequest)
			return
		}
		// path := req.URL.Path
		p.l.Println("Got id:", id)
		p.updateProduct(id, res, req)
		return
	}
	res.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(res http.ResponseWriter, req *http.Request) {
	p.l.Println("Handle GET Products")
	productList := data.GetProducts()
	err := productList.ToJSON(res)
	// d, err := json.Marshal(productList)
	if err != nil {
		http.Error(res, "Unable to parse json", http.StatusInternalServerError)
	}
	// res.Write(d)
}

func (p *Products) addProduct(res http.ResponseWriter, req *http.Request) {
	p.l.Println("Handle POST Products")
	prod := &data.Product{}
	err := prod.FromJSON(req.Body)
	if err != nil {
		http.Error(res, "Unable to marshal json", http.StatusBadRequest)
	}

	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(prod)
}

func (p *Products) updateProduct(id int, res http.ResponseWriter, req *http.Request) {
	p.l.Println("Handle PUT Products")
	prod := &data.Product{}
	err := prod.FromJSON(req.Body)
	if err != nil {
		http.Error(res, "Unable to marshal json", http.StatusBadRequest)
	}

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
