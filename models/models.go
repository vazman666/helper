package models

//import "gitlab.skillbox.ru/timur_taitsenov/go_developer_pro/lesson4/sources/pkg/mod/golang.org/x/text@v0.3.6/number"

type price_st struct {
	number           string //номер детали
	name             string //название детали
	PresencePrice    string //старая закупка
	SalesPrice       string //старая продажная
	NewPresencePrice string //закупочная из прайса
	NewSalesPrice    string //новая продажная
}
