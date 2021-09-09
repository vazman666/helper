package pkg

import (
	"fmt"
	"helper/models"
	"strconv"

	"github.com/tealeg/xlsx/v3"
)

//func Loads(price map[int]models.Price_st) {

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
	/*models.Price[0] = models.Price_st{Number: "8975248110",
		Name:             "Корпус ключа",
		PresencePrice:    "699",
		SalesPrice:       "1150",
		NewPresencePrice: "710",
		NewSalesPrice:    "1150"}
	models.Price[1] = models.Price_st{Number: "8975248110",
		Name:             "Корпус ключа",
		PresencePrice:    "699",
		SalesPrice:       "1150",
		NewPresencePrice: "710",
		NewSalesPrice:    "1150"}*/
	// open an existing file
	wb, err := xlsx.OpenFile("tiss08_09_2021.xlsx")
	if err != nil {
		panic(err)
	}
	// wb now contains a reference to the workbook
	// show all the sheets in the workbook
	fmt.Println("Sheets in this file:")
	for i, sh := range wb.Sheets {
		fmt.Println(i, sh.Name)
	}
	fmt.Println("----")
	// wb contains a reference to an opened workbook
	fmt.Println("Workbook contains", len(wb.Sheets), "sheets.")
	sheetName := "Sheet"
	sh, ok := wb.Sheet[sheetName]
	if !ok {
		fmt.Println("Sheet does not exist")
		return
	}

	fmt.Println("Max row in sheet:", sh.MaxRow)
	/* sh is a reference to a sheet, see above
	row, err := sh.Row(2)
	if err != nil {
		panic(err)
	}*/

	sh.ForEachRow(rowVisitor)
	/*theCell, err := sh.Cell(0, 3)
	if err != nil {
		panic(err)
	}
	 we got a cell, but what's in it?
	fv, err := theCell.FormattedValue()
	if err != nil {
		panic(err)
	}
	fmt.Println("Numeric cell?:", theCell.Type() == xlsx.CellTypeNumeric)
	fmt.Println("String:", theCell.String())
	fmt.Println("Formatted:", fv)
	fmt.Println("Formula:", theCell.Formula())*/
	for i := 1; i <= 5; /*sh.MaxRow;*/ i++ {
		//for i, val := range sh {
		fmt.Printf("%v\n", i)
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
		//fmt.Printf("строка %v, значение %v", i, theCellVal)
		newpresencepriceInt, _ := strconv.ParseFloat(newpresenceprice, 2)
		//sales205, _ := strconv.ParseFloat(newpresenceprice, 2)
		sales165 := fmt.Sprintf("%5.2f", newpresencepriceInt*1.65)
		sales205 := fmt.Sprintf("%5.2f", newpresencepriceInt*2.05)
		//sales165Str := strconv.Itoa(sales165 * 1.65)		//sales205Str := strconv.Itoa(sales205 * 2.05)
		tmp := models.Price_st{Number: number, Name: name, NewPresencePrice: newpresenceprice, Firm: firm, Remark: remark, Sales165: sales165, Sales205: sales205}
		models.Price = append(models.Price, tmp)

	}
	Addsql()
	CreateXLSX()
}
