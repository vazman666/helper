package models

import "sync"

//import "gitlab.skillbox.ru/timur_taitsenov/go_developer_pro/lesson4/sources/pkg/mod/golang.org/x/text@v0.3.6/number"
//const Providers=1\\количество обрабатываемых поставщиков
type Price_st struct {
	Remark           string //кому
	Number           string //номер детали
	Firm             string //Фирма запчасти
	Name             string //название детали
	PresencePrice    string //старая закупка
	SalesPrice       string //старая продажная
	NewPresencePrice string //закупочная из прайса
	Sales165         string
	Sales205         string
}

var M sync.Mutex
var Price []Price_st
var Provider [9]int

func InitProvider(provider string) {
	//Price = make([]Price_st, 1)
	switch provider {
	case "tiss":
		Provider = [9]int{6, 1, 0, 2, 3, 3, 3, 3, 3}
	default:
		Provider = [9]int{6, 0, 0, 0, 0, 0, 0, 0, 0}
	}
}
