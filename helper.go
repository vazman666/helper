package main

import (
	"fmt"
	"helper/models"
	"helper/pkg"
	//"github.com/gorilla/mux"
)

var Price []models.Price_st

func main() {
	models.Price = make([]models.Price_st, 1)
	models.Xlsx = make([]models.Price_st, 1)

	pkg.Readtiss()
	pkg.Addsql()
	pkg.CreateXLSX()

	fmt.Printf("Xls: %v\n", models.Xlsx)

}
