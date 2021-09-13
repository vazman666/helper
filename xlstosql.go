package main

import (
	"database/sql"
	"fmt"
	"helper/models"
	"log"

	_ "github.com/denisenkom/go-mssqldb"

	"github.com/tealeg/xlsx/v3"
)

var Price []models.Price_st

func main() {

	condb, errdb := sql.Open("mssql", "server=192.168.1.40;user id=admin;password=12345;")
	if errdb != nil {
		fmt.Println(" Error open db:", errdb.Error())
	}

	wb, err := xlsx.OpenFile("a.xlsx")
	if err != nil {
		panic(err)
	}

	sh, ok := wb.Sheet["test"]
	if !ok {
		fmt.Println("Sheet does not exist")
		return
	}

	for i := 1; i < sh.MaxRow; i++ {

		theCell, err := sh.Cell(i, 6)
		if err != nil {
			panic(err)
		}
		newpresenceprice, err := theCell.FormattedValue()
		if err != nil {
			panic(err)
		}

		theCell, err = sh.Cell(i, 1)
		if err != nil {
			panic(err)
		}
		number, err := theCell.FormattedValue()
		if err != nil {
			panic(err)
		}

		theCell, err = sh.Cell(i, 5)
		if err != nil {
			panic(err)
		}
		newprice, err := theCell.FormattedValue()
		if err != nil {
			panic(err)
		}

		theCell, err = sh.Cell(i, 11)
		if err != nil {
			panic(err)
		}
		oid, err := theCell.FormattedValue()
		if err != nil {
			panic(err)
		}

		if newprice != "---" {
			if number == "" || oid == "" {
				continue
			}
			fmt.Printf("Number=%v\n", number)
			Oid := ""
			rows, err := condb.Query("SELECT  [Oid] FROM [basebasebase].[dbo].[Ware]  WHERE [OriginalCode]=? AND [OID]=?", number, oid)
			fmt.Println("--s--")
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
			fmt.Printf("Number %v закупка будет %v Продажа будет %v. OID=%v\n", number, newpresenceprice, newprice, Oid)
			rows, err = condb.Query("UPDATE [basebasebase].[dbo].[PricePresence] SET [SalesPrice]=?  WHERE [Ware]=? AND [fPricelistPresence]=1", newprice, Oid)
			if err != nil {
				log.Fatal(err)
			}
			rows, err = condb.Query("UPDATE [basebasebase].[dbo].[PricePresence] SET [SalesPrice]=?  WHERE [Ware]=? AND [fPricelistPresence]=2", newprice, Oid)
			if err != nil {
				log.Fatal(err)
			}
			rows, err = condb.Query("UPDATE [basebasebase].[dbo].[PricePresence] SET [PresencePrice]=?  WHERE [Ware]=? AND [fPricelistPresence]=1", newpresenceprice, Oid)
			if err != nil {
				log.Fatal(err)
			}
			rows, err = condb.Query("UPDATE [basebasebase].[dbo].[PricePresence] SET [PresencePrice]=?  WHERE [Ware]=? AND [fPricelistPresence]=2", newpresenceprice, Oid)
			if err != nil {
				log.Fatal(err)
			}

		}
	}
}
