package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

///  *** EMPLOYEE ****
// create employee handler
func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp Employee
	json.NewDecoder(r.Body).Decode(&emp)
	Database.Create(&emp)
	json.NewEncoder(w).Encode(emp)
}

// Get GetEmployees handler
func GetEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employees []Employee
	Database.Find(&employees)
	json.NewEncoder(w).Encode(employees)
}

// create Employess handler
func GetEmployeeById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee Employee
	Database.First(&employee, mux.Vars(r)["id"])
	json.NewEncoder(w).Encode(employee)
}

// update employees
func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee Employee
	Database.First(&employee, mux.Vars(r)["id"])
	json.NewDecoder(r.Body).Decode(&employee)
	Database.Save(&employee)
	json.NewEncoder(w).Encode(employee)

}

// delete employee
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp Employee
	Database.Delete(&emp, mux.Vars(r)["id"])
	json.NewEncoder(w).Encode("Employee is deleted ||")
}

///  *** PRODUCTS ****
//Create
func CreateProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var prd Product
	json.NewDecoder(r.Body).Decode(&prd)
	Database.Create(&prd)
	json.NewEncoder(w).Encode(prd)
}

// create GetEmployees handler
func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var products []Product
	Database.Find(&products)
	json.NewEncoder(w).Encode(products)
}

//GET products
func GetProductsID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var products Product
	Database.First(&products, mux.Vars(r)["id"])
	json.NewEncoder(w).Encode(products)
}

//update
func UpdateProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var products Product
	Database.First(&products, mux.Vars(r)["id"])
	json.NewDecoder(r.Body).Decode(&products)
	Database.Save(&products)
	json.NewEncoder(w).Encode(products)

}

// delete
func DeleteProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var prd Product
	Database.Delete(&prd, mux.Vars(r)["id"])
	json.NewEncoder(w).Encode("Product is deleted")
}
