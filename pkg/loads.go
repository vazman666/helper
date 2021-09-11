package pkg

import (
	"fmt"
	"helper/models"

	"github.com/tealeg/xlsx/v3"
)

func rowVisitor(r *xlsx.Row) error {
	return r.ForEachCell(cellVisitor)
}

func cellVisitor(c *xlsx.Cell) error {
	_, err := c.FormattedValue()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		//fmt.Println("Cell value:", value)
	}
	return err
}

func Loads() {
	models.Price = make([]models.Price_st, 1)

	Readtiss()
	Addsql()
	CreateXLSX()
}
