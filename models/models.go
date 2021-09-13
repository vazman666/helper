package models

//import "gitlab.skillbox.ru/timur_taitsenov/go_developer_pro/lesson4/sources/pkg/mod/golang.org/x/text@v0.3.6/number"
//const Providers=1\\количество обрабатываемых поставщиков
type Price_st struct {
	Remark           string //кому
	Number           string //номер детали
	Firm             string //Фирма запчасти из прайса
	Name             string //название детали
	PresencePrice    string //старая закупка
	SalesPrice       string //старая продажная
	NewPresencePrice string //закупочная из прайса
	Sales165         string
	Sales205         string
	FirmSql          string //фирма из базы
	WareHouse        string //ячейка
	Oid              string
}

var Price []Price_st
var Xlsx []Price_st
