package pkg

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

func Change(newprice, number, oid, newpresenceprice string) {
	condb, errdb := sql.Open("mssql", "server=192.168.1.40;user id=admin;password=12345;")
	if errdb != nil {
		log.Fatalf(" Error open db:%v\n", errdb.Error())
	}
	if newprice != "---" {
		if number == "" || oid == "" {
			return
		}
		fmt.Printf("Number=%v\n", number)
		Oid := ""
		rows, err := condb.Query("SELECT  [Oid] FROM [basebasebase].[dbo].[Ware]  WHERE [OriginalCode]=? AND [OID]=?", number, oid)
		//fmt.Println("--s--")
		if err != nil {
			log.Fatal(err)
		}
		for rows.Next() {

			err := rows.Scan(&Oid)
			if err != nil {
				log.Fatal("Error: ", err)
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
