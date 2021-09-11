package pkg

import (
	"fmt"

	"helper/models"

	"github.com/tealeg/xlsx/v3"
)

func Loads2() {
	var mySlice [][][]string
	//var value string
	mySlice, _ = xlsx.FileToSlice("tiss08_09_2021.xlsx")
	//value = mySlice[0][0][0]
	for i, val := range mySlice[0] {
		for j, val2 := range val {
			switch j {
			case models.Provider[0]:
				models.Price[i].Remark = val2
			}

			fmt.Printf("i=%v j=%v val2=%v   ", i, j, val2)
		}
	}
	//fmt.Printf("%v\n", mySlice)
}
