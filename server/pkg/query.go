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
	firm = Firm(firm)
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
	defer rows.Close()
	for rows.Next() {

		if err := rows.Scan(&tmp.Firm); err != nil {
			log.Fatal("Ошибка выбора OID фирмы %v\n", err)
		}

	}
	rows, err = condb.Query("SELECT [Oid], [Caption]  FROM [basebasebase].[dbo].[Ware]  WHERE [Code]=? AND [Producer]=?", number, tmp.Firm) //выбираем Oid запчасти по фирме и номеру запчасти
	if err != nil {
		log.Fatal("Ошибка выбора OID запчасти %v\n", err)
	}
	defer rows.Close()
	for rows.Next() {

		if err := rows.Scan(&tmp.Oid, &tmp.Name); err != nil {
			log.Fatal(err)
		}

	}
	if tmp.Oid == "" {
		fmt.Printf("Not fount in base num=%v firm=%v\n", number, firm)
		tmp.PresencePrice = "----"
		tmp.SalesPrice = "----"
		return tmp

	}
	rows, err = condb.Query("SELECT  [DefaultWarehouse] FROM [basebasebase].[dbo].[WarehouseMinQuantity]  WHERE [ParentWare]=? AND [Warehouse] = 619", tmp.Oid) //ячейка Мытищи
	if err != nil {
		log.Fatal("Ошибка запроса при выборе ячейка Мытищи %v\n", err)
	}
	defer rows.Close()
	for rows.Next() {

		if err := rows.Scan(&tmp.Cellm); err != nil {
			log.Fatal("Ошибка при scan ячейка Мытищи %v\n", err)
		}

	}

	rows, err = condb.Query("SELECT  [Path] FROM [basebasebase].[dbo].[Warehouse]  WHERE [Oid]=? ", tmp.Cellm)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {

		if err := rows.Scan(&tmp.Cellm); err != nil {
			log.Fatal("Ошибка при scan конекретно ячейки Мытищи %v\n", err)
		}

	}
	rows, err = condb.Query("SELECT  [DefaultWarehouse] FROM [basebasebase].[dbo].[WarehouseMinQuantity]  WHERE [ParentWare]=? AND [Warehouse] = 168", tmp.Oid) //ячейка Титан
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {

		if err := rows.Scan(&tmp.Cellt); err != nil {
			log.Fatal(err)
		}

	}

	rows, err = condb.Query("SELECT  [Path] FROM [basebasebase].[dbo].[Warehouse]  WHERE [Oid]=? ", tmp.Cellt)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {

		if err := rows.Scan(&tmp.Cellt); err != nil {
			log.Fatal("Ошибка скан конкретно ячейки Титан %v\n", err)
		}

	}

	rows, err = condb.Query("SELECT  [PresencePrice], [SalesPrice] FROM [basebasebase].[dbo].[PricePresence] WHERE [Ware]=?", tmp.Oid)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {

		if err := rows.Scan(&tmp.PresencePrice, &tmp.SalesPrice); err != nil {
			log.Fatal("Ошибка scan прайсов %v\n", err)
		}

	}

	pp, _ := strconv.ParseFloat(tmp.PresencePrice, 2)
	sp, _ := strconv.ParseFloat(tmp.SalesPrice, 2)
	tmp.PresencePrice = fmt.Sprintf("%5.2f", pp)
	tmp.SalesPrice = fmt.Sprintf("%5.2f", sp)

	return tmp
}

var Flag bool

func Analogue(detail Analog) {

	condb, errdb := sql.Open("mssql", "server=192.168.1.40;user id=admin;password=12345;")
	if errdb != nil {
		fmt.Println(" Error open db:", errdb.Error())
	}
	defer condb.Close()
	//fmt.Printf("Scan number=%v  firm=%v\n", detail.Number, detail.Firm)
	Analogs = append(Analogs, detail)
	for Flag = true; Flag; {
		for _, a := range Analogs {

			recurA(a, condb)
			recurB(a, condb)
		}
		//Flag = false
	}
	Analogs = Analogs[1:]
}
func recurA(num Analog, condb *sql.DB) {
	var analog Analog
	Flag = false
	//fmt.Printf("Ищем для firm=%v Number=%v\n", num.Firm, num.Number)
	rows, err := condb.Query("SELECT [DetailA], [ProducerA] FROM [basebasebase].[dbo].[Analogue]  WHERE [DetailB]=? AND [ProducerB]=?", num.Number, num.Firm)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {

		if err := rows.Scan(&analog.Number, &analog.Firm); err != nil {
			log.Fatal("Ошибка выбора OID фирмы %v\n", err)
		}
		//fmt.Printf("RS2 Найден аналог:%v\n", analog)
		if contains(Analogs, analog) {
			return
		}
		Analogs = append(Analogs, analog)
		recurA(analog, condb)
		Flag = true

	}
	return

}

func recurB(num Analog, condb *sql.DB) {
	var analog Analog
	Flag = false

	rows, err := condb.Query("SELECT  [DetailB], [ProducerB] FROM [basebasebase].[dbo].[Analogue]  WHERE [DetailA]=? AND [ProducerA]=?", num.Number, num.Firm)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {

		if err := rows.Scan(&analog.Number, &analog.Firm); err != nil {
			log.Fatal("Ошибка выбора OID фирмы %v\n", err)
		}
		//fmt.Printf("Найден аналог:%v\n", analog)
		if contains(Analogs, analog) {
			return
		}
		Analogs = append(Analogs, analog)
		recurB(analog, condb)
		Flag = true
	}
	return

}

// Contains указывает, содержится ли x в a.
func contains(a []Analog, x Analog) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
