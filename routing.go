package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Router function
func HandlerRouting() {
	r := mux.NewRouter()
	//Employee
	r.HandleFunc("/employee", GetEmployees).Methods("GET")
	r.HandleFunc("/employee/{id}", GetEmployeeById).Methods("GET")
	r.HandleFunc("/employee", CreateEmployee).Methods("POST")
	r.HandleFunc("/employee/{id}", UpdateEmployee).Methods("PUT")
	r.HandleFunc("/employee/{id}", DeleteEmployee).Methods("DELETE")
	//Products
	r.HandleFunc("/products", GetProducts).Methods("GET")
	r.HandleFunc("/products/{id}", GetProductsID).Methods("GET")
	r.HandleFunc("/products", CreateProducts).Methods("POST")
	r.HandleFunc("/products/{id}", UpdateProducts).Methods("PUT")
	r.HandleFunc("/products/{id}", DeleteProducts).Methods("DELETE")
	//Users
	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", GetUsersID).Methods("GET")
	r.HandleFunc("/users", CreateUsers).Methods("POST")
	r.HandleFunc("/users/{id}", UpdateUsers).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUsers).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8085", r))
}
