package models

import "sync"

//import "gitlab.skillbox.ru/timur_taitsenov/go_developer_pro/lesson4/sources/pkg/mod/golang.org/x/text@v0.3.6/number"
type provider struct {
	tiss int
}
type Price_st struct {
	Remark           provider //кому
	Number           string   //номер детали
	Firm             string   //Фирма запчасти
	Name             string   //название детали
	PresencePrice    string   //старая закупка
	SalesPrice       string   //старая продажная
	NewPresencePrice string   //закупочная из прайса
	Sales165         string
	Sales205         string
}

var M sync.Mutex
var Price []Price_st
