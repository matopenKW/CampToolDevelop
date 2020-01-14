package main

import (
	"CampToolDevelop/internal/apps"
	_ "github.com/loadoff/excl"
	"github.com/tealeg/xlsx"
	"log"
	"strings"
	"time"

	"fmt"
)

const EXCEL_PATH = "./work/"

const TAG_1 = "$"
const TAG_2 = "_"

var positionMap = make(map[string]CellPosition)

func main() {

	read1()
	write()

}

type CellPosition struct {
	key       string
	rowIndex  int
	cellIndex int
}

func (cellp *CellPosition) getPosition() (rowIdx int, cellIdx int) {
	return cellp.rowIndex, cellp.cellIndex
}

func read1() {
	file, err := xlsx.OpenFile(EXCEL_PATH + "test1.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	for _, sheet := range file.Sheets {
		for rowIdx, row := range sheet.Rows {
			for cellIdx, cell := range row.Cells {
				v := cell.String()

				if strings.Index(v, TAG_1) > -1 {
					cellp := CellPosition{
						key:       v,
						rowIndex:  rowIdx,
						cellIndex: cellIdx,
					}

					positionMap[v] = cellp

				}

			}
		}
	}
}

func write() {

	file, err := xlsx.OpenFile(EXCEL_PATH + "test1.xlsx")
	if err != nil {
		log.Fatal(err)
	}
	sheet := file.Sheets[0]

	print := getDataMap()

	log.Println(positionMap)

	for k, v := range print {

		cellp := positionMap["${"+k+"}"]

		sheet.Cell(cellp.getPosition()).Value = v
	}

	err = file.Save(EXCEL_PATH + "new.xlsx")
	if err != nil {
		log.Printf(err.Error())
	}

	log.Println("success")
}

// func read2() {

// 	wookbook, _ := excl.Open(EXCEL_PATH + "test1.xlsx")
// 	defer wookbook.Close()

// 	sheet, _ := wookbook.OpenSheet("Sheet1")
// 	defer sheet.Close()

// 	i := 0
// 	for i < 5 {
// 		row := sheet.GetRow(i)

// 		cells := row.CreateCells(1, 5)

// 		for _, v := range cells {

// 			log.Println(v)

// 		}

// 		// cell1 := row.GetCell(2)

// 		i++
// 	}

// }

// func print() {

// 	wookbook, _ := excl.Open(EXCEL_PATH + "test1.xlsx")
// 	defer wookbook.Close()

// 	sheet, _ := wookbook.OpenSheet("Sheet1")
// 	defer sheet.Close()

// 	startIndex := 5
// 	for i, v := range selectDataList() {

// 		row := sheet.GetRow(startIndex + i)

// 		cell1 := row.GetCell(2)
// 		cell1.SetDate(cnvDate(v.Date))

// 		log.Println(v, i)

// 	}

// 	wookbook.Save(EXCEL_PATH + "new.xlsx")
// }

func getDataMap() map[string]string {

	ret := make(map[string]string)

	ret["name"] = "Kazurnoti Ono"

	list := selectDataList()

	for i, v := range list {
		idx := i + 1
		ret[fmt.Sprintf("IKI_%d", idx)] = v.Start
		ret[fmt.Sprintf("kaeri_%d", idx)] = v.End
		ret[fmt.Sprintf("bikou_%d", idx)] = v.Bikou
	}

	log.Println(ret)

	return ret
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
