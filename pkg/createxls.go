package pkg

import (
	"helper/models"

	"github.com/tealeg/xlsx/v3"
)

func CreateXLSX() {
	handle := models.Price_st{
		Remark:           "Кому",
		Number:           "Номер запчатсти",
		Firm:             "производитель",
		Name:             "Наименование",
		PresencePrice:    "Закупка из SQL",
		SalesPrice:       "Продажная из SQL",
		NewPresencePrice: "Новаая закупка",
		Sales165:         "*1,65",
		Sales205:         "*2,05"}

	wb := xlsx.NewFile() //создаём новый экскиз экселя

	sheetTest, err := wb.AddSheet("test") //добавляем страничку
	if err != nil {
		panic(err)
	}
	sheetTest.SetColWidth(0, 5, 12.5)

	row1 := sheetTest.AddRow()
	_ = row1.WriteStruct(&handle, -1)
	row1.SetHeight(15)
	for i := 0; i < 6; i++ {
		_ = sheetTest.SetColAutoWidth(i, xlsx.DefaultAutoWidth)
	}
	for _, value := range models.Price {

		row1 = sheetTest.AddRow()        //добавляем строку
		_ = row1.WriteStruct(&value, -1) //и вставляе в эту строку строку из прайс
		row1.SetHeight(15)

	}
	err = wb.Save("a.xlsx")
	if err != nil {
		panic(err)
	}

}
