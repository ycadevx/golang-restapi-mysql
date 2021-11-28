package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

//***********************************************************//
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

//***********************************************************//
///  *** PRODUCTS ****
//Create Product
func CreateProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var prd Product
	json.NewDecoder(r.Body).Decode(&prd)
	Database.Create(&prd)
	json.NewEncoder(w).Encode(prd)
}

// Get All Products
func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var products []Product
	Database.Find(&products)
	json.NewEncoder(w).Encode(products)
}

//GET Single Product
func GetProductsID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var products Product
	Database.First(&products, mux.Vars(r)["id"])
	json.NewEncoder(w).Encode(products)
}

//Update Product
func UpdateProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var products Product
	Database.First(&products, mux.Vars(r)["id"])
	json.NewDecoder(r.Body).Decode(&products)
	Database.Save(&products)
	json.NewEncoder(w).Encode(products)

}

// Delete Product
func DeleteProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var prd Product
	Database.Delete(&prd, mux.Vars(r)["id"])
	json.NewEncoder(w).Encode("Product is deleted")
}

//***********************************************************//
// Get All Users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []Users
	Database.Find(&users)
	json.NewEncoder(w).Encode(users)

}

//Create Users
func CreateUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user Users
	json.NewDecoder(r.Body).Decode(&user)
	Database.Create(&user)
	json.NewEncoder(w).Encode(user)
}

//Update Users
func UpdateUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user Users
	Database.First(&user, mux.Vars(r)["id"])
	json.NewDecoder(r.Body).Decode(&user)
	Database.Save(&user)
	json.NewEncoder(w).Encode(user)

}

//GET Single Users
func GetUsersID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var usr Users
	Database.First(&usr, mux.Vars(r)["id"])
	json.NewEncoder(w).Encode(usr)
}

// Delete User
func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var usr Users
	Database.Delete(&usr, mux.Vars(r)["id"])
	json.NewEncoder(w).Encode("User is deleted")
}

//***********************************************************//
