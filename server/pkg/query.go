package pkg

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/denisenkom/go-mssqldb"
)

func QuerySQL(number string) OurData {
	var tmp OurDataStruct
	var rezult OurData
	condb, errdb := sql.Open("mssql", "server=192.168.1.40;user id=admin;password=12345;")
	if errdb != nil {
		fmt.Println(" Error open db:", errdb.Error())
	}
	defer condb.Close()

	rows, err := condb.Query("SELECT TOP (1000) [Oid],[Producer] FROM [basebasebase].[dbo].[Ware]  WHERE [OriginalCode]=? ", number)

	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&tmp.Oid, &tmp.Caption)
		if err != nil {
			log.Fatal(err)

		}
		rows2, err := condb.Query("SELECT  [Caption] FROM [basebasebase].[dbo].[Producer]  WHERE [OID]=? ", tmp.Caption)
		if err != nil {
			log.Fatal(err)
		}
		for rows2.Next() {

			err := rows2.Scan(&tmp.Firm) //фирма из базы
			if err != nil {
				log.Fatal(err)

			}
			rows3, err := condb.Query("SELECT TOP (1000) [PresencePrice], [SalesPrice] FROM [basebasebase].[dbo].[PricePresence] WHERE [Ware]=?", tmp.Oid)
			if err != nil {
				log.Fatal(err)
			}

			for rows3.Next() {

				err := rows3.Scan(&tmp.PresencePrice, &tmp.SalesPrice)
				if err != nil {
					log.Fatal(err)
				}
			}
			pp, _ := strconv.ParseFloat(tmp.PresencePrice, 2)
			sp, _ := strconv.ParseFloat(tmp.SalesPrice, 2)
			tmp.PresencePrice = fmt.Sprintf("%5.2f", pp)
			tmp.SalesPrice = fmt.Sprintf("%5.2f", sp)
			rezult = append(rezult, tmp)

		}
	}
	return rezult
}

/*func OpenSql() {
	condb, errdb := sql.Open("mssql", "server=192.168.1.40;user id=admin;password=12345;")
	if errdb != nil {
		fmt.Println(" Error open db:", errdb.Error())
	}
	defer condb.Close()
}*/
