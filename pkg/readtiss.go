package pkg

import (
	"fmt"
	"helper/models"
	"strconv"

	"github.com/tealeg/xlsx/v3"
)

func Readtiss() {
	wb, err := xlsx.OpenFile("origin.xlsx")
	if err != nil {
		panic(err)
	}

	sh, ok := wb.Sheet["Sheet"]
	if !ok {
		fmt.Println("Sheet does not exist")
		return
	}

	for i := 1; i < sh.MaxRow; i++ {

		theCell, err := sh.Cell(i, 3)
		if err != nil {
			panic(err)
		}
		newpresenceprice, err := theCell.FormattedValue()
		if err != nil {
			panic(err)
		}
		theCell, err = sh.Cell(i, 1)
		if err != nil {
			panic(err)
		}
		number, err := theCell.FormattedValue()
		if err != nil {
			panic(err)
		}
		theCell, err = sh.Cell(i, 2)
		if err != nil {
			panic(err)
		}
		name, err := theCell.FormattedValue()
		if err != nil {
			panic(err)
		}
		theCell, err = sh.Cell(i, 0)
		if err != nil {
			panic(err)
		}
		firm, err := theCell.FormattedValue()
		if err != nil {
			panic(err)
		}
		theCell, err = sh.Cell(i, 6)
		if err != nil {
			panic(err)
		}
		remark, err := theCell.FormattedValue()
		if err != nil {
			panic(err)
		}

		newpresencepriceInt, _ := strconv.ParseFloat(newpresenceprice, 2)

		sales165 := fmt.Sprintf("%5.2f", newpresencepriceInt*1.65)
		sales205 := fmt.Sprintf("%5.2f", newpresencepriceInt*2.05)

		tmp := models.Price_st{Number: number, Name: name, NewPresencePrice: newpresenceprice, Firm: firm, Remark: remark, Sales165: sales165, Sales205: sales205}
		models.Price = append(models.Price, tmp)
	}
}
