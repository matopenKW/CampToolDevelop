package main

import (
	"CampToolDevelop/internal/apps"
	"fmt"
	"time"

	"github.com/loadoff/excl"
)

const EXCEL_PATH = "./work/"

func main() {

	wookbook, _ := excl.Open(EXCEL_PATH + "test1.xlsx")
	defer wookbook.Close()

	sheet, _ := wookbook.OpenSheet("Sheet1")
	defer sheet.Close()

	// 1列目のセルを取得
	// c := r.GetCell(1)
	// // セルに10を出力
	// c.SetNumber("10")
	// // 2列目のセルにABCDEという文字列を出力
	// c = r.SetString("ABCDE", 2)

	startIndex := 5
	for i, v := range selectDataList() {

		row := sheet.GetRow(startIndex + i)

		cell1 := row.GetCell(2)
		cell1.SetDate(cnvDate(v.Date))

		fmt.Println(v, i)

	}

	wookbook.Save(EXCEL_PATH + "new.xlsx")
}

func selectDataList() []*apps.Carfare {
	list := make([]*apps.Carfare, 0, 10)
	carfare := &apps.Carfare{Date: "20200113", Start: "町田", End: "新宿", Price: "2000", Bikou: "テスト"}
	list = append(list, carfare)
	list = append(list, carfare)
	list = append(list, carfare)
	list = append(list, carfare)
	return list

}

func cnvDate(date string) time.Time {

	t, _ := time.Parse("20060102", date)
	return t
}
