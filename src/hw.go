package main

import (
	"fmt"
	"github.com/xuri/excelize"
	"os"
	"parser/tinkoff"
	hr "parser/tinkoff/head_row"
)

func main() {
	fileName := os.Args[1]
	reportSheetName := "broker_rep"

	file, err := excelize.OpenFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	rows, err := file.GetRows(reportSheetName)
	if err != nil {
		panic(err)
	}
	head_rows := hr.FindHeadRows(rows)
	headRowManager := *hr.InitHeadRowManager(head_rows)

	universalOperationParser := *tinkoff.InitUniversalOperationParser(rows, headRowManager)
	for _, operation := range universalOperationParser.GetOperations() {
		fmt.Println(operation)
	}
}
