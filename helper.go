package main

import (
	"fmt"
	"helper/models"
	"helper/pkg"
	//"github.com/gorilla/mux"
)

var Price []models.Price_st

func main() {
	//price := make(map[int]models.Price_st)
	pkg.Loads2()
	pkg.Loads()

	fmt.Printf("price: %v\n", models.Price)
	/*s1 := mux.NewRouter()
	s1.HandleFunc("/", pkg.IndexHandler)
	http.Handle("/", s1)

	fmt.Println("Server is listening...8181")
	http.ListenAndServe(":8181", nil)*/
}
