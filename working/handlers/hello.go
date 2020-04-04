package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	h.l.Println("Hello World")
	d, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "OOPS", http.StatusBadRequest)
		// res.WriteHeader(http.StatusBadRequest)
		// res.Write([]byte("OOPs"))
		return
	}
	h.l.Printf("Data: %s\n", d)

	fmt.Fprintf(res, "Hello %s", d)
}
