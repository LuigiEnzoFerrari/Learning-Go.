package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"io/ioutil"
	// "encoding/json"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	response, _ := http.Get("https://fakestoreapi.com/products?limit=2")
	responseData, _ := ioutil.ReadAll(response.Body)
	// res, _ := json.Marshal(responseData)
	// w.Write(responseData)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
	// []byte("Gorilla!\n")
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", YourHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
