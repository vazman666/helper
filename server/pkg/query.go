package pkg

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/denisenkom/go-mssqldb"
)

func QuerySQL(number, firm string) OurDataStruct { //функкция возвращает по номеру и фирме соответствующую структуру
	var tmp OurDataStruct
	//var rezult OurData
	condb, errdb := sql.Open("mssql", "server=192.168.1.40;user id=admin;password=12345;")
	if errdb != nil {
		fmt.Println(" Error open db:", errdb.Error())
	}
	defer condb.Close()
	fmt.Printf("Scan number=%v  firm=%v\n", number, firm)
	if number == "" {
		return tmp
	}
	rows, err := condb.Query("SELECT  [Oid] FROM [basebasebase].[dbo].[Producer]  WHERE [Caption]=? ", firm) //выбираем OID нужной нам фирмы из базы
	if err != nil {
		log.Fatal(err)
	}
	err = rows.Scan(&tmp.Firm)
	if err != nil {
		log.Fatal(err)

	}

	rows, err = condb.Query("SELECT TOP (1000) [Oid] FROM [basebasebase].[dbo].[Ware]  WHERE [Code]=? AND [Producer]=?", number, tmp.Firm) //выбираем Oid запчасти по фирме и номеру запчасти

	if err != nil {
		log.Fatal(err)
	}

	err = rows.Scan(&tmp.Oid)
	if err != nil {
		log.Fatal(err)

	}
	rows, err = condb.Query("SELECT  [DefaultWarehouse] FROM [basebasebase].[dbo].[WarehouseMinQuantity]  WHERE [ParentWare]=? AND [Warehouse] = 619", tmp.Oid)
	if err != nil {
		log.Fatal(err)
	}
	err = rows.Scan(&tmp.Cell) //код ячейки
	if err != nil {
		log.Fatal(err)
	}
	rows, err = condb.Query("SELECT  [Path] FROM [basebasebase].[dbo].[Warehouse]  WHERE [Oid]=? ", tmp.Cell)
	if err != nil {
		log.Fatal(err)
	}
	err = rows.Scan(&tmp.Cell) //путь ячейки
	if err != nil {
		log.Fatal(err)
	}
	rows, err = condb.Query("SELECT TOP (1000) [PresencePrice], [SalesPrice] FROM [basebasebase].[dbo].[PricePresence] WHERE [Ware]=?", tmp.Oid)
	if err != nil {
		log.Fatal(err)
	}

	err = rows.Scan(&tmp.PresencePrice, &tmp.SalesPrice)
	if err != nil {
		log.Fatal(err)
	}

	pp, _ := strconv.ParseFloat(tmp.PresencePrice, 2)
	sp, _ := strconv.ParseFloat(tmp.SalesPrice, 2)
	tmp.PresencePrice = fmt.Sprintf("%5.2f", pp)
	tmp.SalesPrice = fmt.Sprintf("%5.2f", sp)
	//rezult = append(rezult, tmp)

	return tmp
}

/*func OpenSql() {
	condb, errdb := sql.Open("mssql", "server=192.168.1.40;user id=admin;password=12345;")
	if errdb != nil {
		fmt.Println(" Error open db:", errdb.Error())
	}
	defer condb.Close()
}*/
