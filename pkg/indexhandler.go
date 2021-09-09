package pkg

import (
	//"helper/template"
	"fmt"
	"html/template"
	"net/http"

	"helper/models"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	models.M.Lock()
	defer models.M.Unlock()
	fmt.Printf("price: %v\n", models.Price)

	/*
		products := []models.Product{}

		for rows.Next() {

			p := models.Product{}
			err := rows.Scan(&p.Id, &p.Partnum, &p.Qty, &p.Price, &p.Provider, &p.Name, &p.Remark, &p.Date, &p.Color)
			if err != nil {
				fmt.Println(err)
				continue
			}
			products = append(products, p)

		}*/

	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, models.Price)

}
