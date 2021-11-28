package main

import "gorm.io/gorm"

//
type Employee struct {
	gorm.Model
	EmpName   string  `json:"empnames"`
	EmpSalary float64 `json:"salary"`
	Email     string  `json:"email"`
}

type Product struct {
	gorm.Model
	Description string  `json:"prdname"`
	Price       float64 `json:"price"`
	Unite       string  `json:"unite"`
	CostPrice   float64 `json:"costprice"`
}

type Users struct {
	gorm.Model
	Name     string `json:"usr_name"`
	LastName string `json:"usr_lastname"`
	Password string `json:"usr_password"`
	Email    string `json:"usr_email"`
}
