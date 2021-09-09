package pkg

import (
	"fmt"

	"github.com/tealeg/xlsx/v3"
)

func Loads2() {
	var mySlice [][][]string
	//var value string
	mySlice, _ = xlsx.FileToSlice("tiss08_09_2021.xlsx")
	//value = mySlice[0][0][0]
	fmt.Printf("%v\n", mySlice)
}
