package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Router function
func HandlerRouting() {
	r := mux.NewRouter()
	r.HandleFunc("/employee", GetEmployees).Methods("GET")
	r.HandleFunc("/employee/{id}", GetEmployeeById).Methods("GET")
	r.HandleFunc("/employee", CreateEmployee).Methods("POST")
	r.HandleFunc("/employee/{id}", UpdateEmployee).Methods("PUT")
	r.HandleFunc("/employee/{id}", DeleteEmployee).Methods("DELETE")

	r.HandleFunc("/products", GetProducts).Methods("GET")
	r.HandleFunc("/products/{id}", GetProductsID).Methods("GET")
	r.HandleFunc("/products", CreateProducts).Methods("POST")
	r.HandleFunc("/products/{id}", UpdateProducts).Methods("PUT")
	r.HandleFunc("/products/{id}", DeleteProducts).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8085", r))
}
