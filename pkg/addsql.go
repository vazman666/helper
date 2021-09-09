package pkg

import (
	//"database/sql"
	"database/sql"
	"fmt"
	"helper/models"
	"log"
	"strconv"

	_ "github.com/denisenkom/go-mssqldb"

	_ "github.com/denisenkom/go-mssqldb"
)

func Addsql() {
	fmt.Printf("startting sql\n")
	condb, errdb := sql.Open("mssql", "server=192.168.1.40;user id=admin;password=12345;")
	if errdb != nil {
		fmt.Println(" Error open db:", errdb.Error())
	}
	var (
		Oid string
		//Ware          string
		//OriginalCode  string
		PresencePrice string
		SalesPrice    string
	)
	rows, err := condb.Query("use basebasebase") //используем базу данных basebasebase
	if err != nil {
		log.Fatal(err)
	}

	for i, value := range models.Price {
		fmt.Printf("\n %v   %v\n", i, value)
		Oid = ""
		rows, err = condb.Query("SELECT  [Oid] FROM [basebasebase].[dbo].[Ware]  WHERE [OriginalCode]=? ", value.Number)
		if err != nil {
			log.Fatal(err)
		}
		for rows.Next() {

			err := rows.Scan(&Oid)
			if err != nil {
				//log.Fatal("Error: ", err)
				log.Println(err)
				fmt.Printf("\n err=%v\n", err)
			}

		}
		fmt.Printf("\n\n\nOid=   %v\n", Oid)
		rows, err = condb.Query("SELECT TOP (1000) [PresencePrice], [SalesPrice] FROM [basebasebase].[dbo].[PricePresence] WHERE [Ware]=?", Oid) //используем базу данных tim
		if err != nil {
			log.Fatal(err)
		}

		for rows.Next() {

			err := rows.Scan(&PresencePrice, &SalesPrice)
			if err != nil {
				//log.Fatal("Error: ", err)
				log.Println(err)
			}
			log.Println(PresencePrice, SalesPrice)
		}
		if Oid == "" {
			PresencePrice = "---"
			SalesPrice = "---"
		} else {
			pp, _ := strconv.ParseFloat(PresencePrice, 2)
			sp, _ := strconv.ParseFloat(SalesPrice, 2)
			PresencePrice = fmt.Sprintf("%5.2f", pp)
			SalesPrice = fmt.Sprintf("%5.2f", sp)
		}

		models.Price[i].PresencePrice = PresencePrice
		models.Price[i].SalesPrice = SalesPrice
	}

}
