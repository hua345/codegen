package excel

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func ParseExcelApi(excelPath string) {
	excelFile, err := excelize.OpenFile(excelPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, value := range excelFile.GetSheetMap() {
		rows, err := excelFile.GetRows(value)
		if err != nil {
			panic(nil)
		}
		for _, row := range rows {
			for _, colCell := range row {
				fmt.Print(colCell, "\t")
			}
			fmt.Println()
		}
	}
}
