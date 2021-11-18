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
