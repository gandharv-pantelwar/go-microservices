package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		log.Println("Hello World")
		d, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(res, "OOPS", http.StatusBadRequest)
			// res.WriteHeader(http.StatusBadRequest)
			// res.Write([]byte("OOPs"))
			return
		}
		log.Printf("Data: %s\n", d)

		fmt.Fprintf(res, "Hello %s", d)
	})
	http.ListenAndServe(":9090", nil)
}
