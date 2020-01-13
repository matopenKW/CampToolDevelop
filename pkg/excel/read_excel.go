package main

import (
	"github.com/loadoff/excl"
)

const EXCEL_PATH = "./work/"

func main() {
	// Excelファイルを読み込み
	w, _ := excl.Open(EXCEL_PATH + "test1.xlsx")
	// シートを開く
	s, _ := w.OpenSheet("Sheet1")
	// 一行目を取得
	r := s.GetRow(1)
	// 1列目のセルを取得
	c := r.GetCell(1)
	// セルに10を出力
	c.SetNumber("10")
	// 2列目のセルにABCDEという文字列を出力
	c = r.SetString("ABCDE", 2)
	// シートを閉じる
	s.Close()
	// 保存
	w.Save(EXCEL_PATH + "new.xlsx")

}
