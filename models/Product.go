package main

import "database/sql"

//Product - A struct for an object from the Products table - Product
type Product struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
}

func GetByID() error {
	sql.Open()
}

