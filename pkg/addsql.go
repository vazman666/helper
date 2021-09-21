package pkg

import (
	//"database/sql"
	"database/sql"
	"fmt"
	"helper/models"
	"log"
	"strconv"
	"strings"

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
		Oid           string
		Producer      string
		Caption       string
		PresencePrice string
		SalesPrice    string
	)
	rows, err := condb.Query("use basebasebase") //используем базу данных basebasebase
	if err != nil {
		log.Fatal(err)
	}

	for i, value := range models.Price {

		Oid = ""
		rows, err = condb.Query("SELECT TOP (1000) [Oid],[Producer] FROM [basebasebase].[dbo].[Ware]  WHERE [OriginalCode]=? ", value.Number)
		if err != nil {
			log.Fatal(err)
		}
		models.Price[i].PresencePrice = "нет в sql"
		models.Price[i].SalesPrice = "нет в sql"
		models.Price[i].FirmSql = "нет в sql"
		models.Price[i].Oid = "нет в sql"
		for rows.Next() {

			err := rows.Scan(&Oid, &Producer)
			if err != nil {
				log.Fatal(err)

			}
			rows2, err := condb.Query("SELECT  [Caption] FROM [basebasebase].[dbo].[Producer]  WHERE [OID]=? ", Producer)
			if err != nil {
				log.Fatal(err)
			}
			for rows2.Next() {

				err := rows2.Scan(&Caption) //фирма из базы
				if err != nil {
					log.Fatal(err)

				}

			}
			if value.Firm == "JD" {
				value.Firm = "Just Drive"
			}
			if strings.ToUpper(Caption) != strings.ToUpper(value.Firm) { //Если фирма из накладной не соответствует фирме из базы,
				continue
			}
			rows3, err := condb.Query("SELECT TOP (1000) [PresencePrice], [SalesPrice] FROM [basebasebase].[dbo].[PricePresence] WHERE [Ware]=?", Oid) //используем базу данных tim
			if err != nil {
				log.Fatal(err)
			}

			for rows3.Next() {

				err := rows3.Scan(&PresencePrice, &SalesPrice)
				if err != nil {
					log.Fatal(err)
				}

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
			models.Price[i].FirmSql = Caption
			models.Price[i].Oid = Oid

		}
		if models.Price[i].Number != "" {
			models.Xlsx = append(models.Xlsx, models.Price[i])
		}
	}

}
