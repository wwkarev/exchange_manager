package head_row

import "errors"

func FindHeadRows(rows [][]string) []HeadRow {
	var headRows []HeadRow
	index := 0
	for {
		headRow, err := findNextHeadRowInfo(rows, index)
		if err != nil {
			break
		}
		headRows = append(headRows, headRow)
		index = headRow.Index + 1
	}
	return headRows
}

func findNextHeadRowInfo(rows [][]string, startIndex int) (HeadRow, error) {
	var index int
	var result string
	err := errors.New("Header row not found")
	for i, row := range rows[startIndex:] {
		var isHeadRow bool
		result, isHeadRow = checkOnHeadRow(row)
		if isHeadRow {
			index = i + startIndex
			err = nil
			break

		}
	}
	return HeadRow{index, result}, err
}

func checkOnHeadRow(row []string) (string, bool) {
	result := ""
	isHeadRow := true
	if len(row) == 1 {
		result = row[0]
	} else if len(row) > 1 {
		result = row[1]
		for i, cellValue := range row {
			if i != 1 && cellValue != "" {
				isHeadRow = false
				break
			}
		}
	} else {
		isHeadRow = false
	}
	return result, isHeadRow
}
