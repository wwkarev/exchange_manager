package tinkoff

import (
	"parser/tinkoff/head_row"
	"time"
)

type currencyIndexInterval struct {
	startIndex int
	endIndex   int
}

type rawOperation struct {
	datetime      time.Time
	operationType string
	sum           float32
	description   string
}

type rawOperationParser struct {
	rows                   [][]string
	headRowManager         head_row.HeadRowManager
	currencyIndexIntervals map[string]currencyIndexInterval
}

func InitRawOperationParser(rows [][]string, headRowManager head_row.HeadRowManager) *rawOperationParser {
	parser := rawOperationParser{}
	parser.rows = rows
	parser.headRowManager = headRowManager
	parser.currencyIndexIntervals = parser.getCurrencyIndexIntervals()
	return &parser
}

func (parser rawOperationParser) GetCurrencyOperationsMap() map[string][]rawOperation {
	currencyOperationsMap := map[string][]rawOperation{}
	for currency, indexInterval := range parser.currencyIndexIntervals {
		currencyOperationsMap[currency] = parser.getRawOperations(parser.rows[indexInterval.startIndex:indexInterval.endIndex])
	}
	return currencyOperationsMap
}

func (parser rawOperationParser) getCurrencyIndexIntervals() map[string]currencyIndexInterval {
	headRow, exists := parser.headRowManager.GetHeadRowByValue(headNameValue.operationsWithCash)
	if !exists {
		panic("No head row with value: " + headNameValue.operationsWithCash)
	}

	currencyIndexIntervals := map[string]currencyIndexInterval{}

	currentCurrencyHeadRow, exists := parser.headRowManager.GetNextHeadRow(headRow)
	_, exists = currencySet[currentCurrencyHeadRow.Value]
	if exists {
		for {
			nextHeadRow, exists := parser.headRowManager.GetNextHeadRow(currentCurrencyHeadRow)
			if !exists {
				break
			}

			currencyIndexIntervals[currentCurrencyHeadRow.Value] = currencyIndexInterval{currentCurrencyHeadRow.Index, nextHeadRow.Index}

			_, exists = currencySet[nextHeadRow.Value]
			if exists {
				currentCurrencyHeadRow = nextHeadRow
			} else {
				break
			}
		}
	}
	return currencyIndexIntervals
}

func (parser rawOperationParser) getRawOperations(rows [][]string) []rawOperation {
	var operations []rawOperation
	for _, row := range rows {
		if len(row) > currencyOperationColumnIndex.operationType {
			rawOperationType := row[currencyOperationColumnIndex.operationType]
			sumColumnNumber, isOperationTypeExists := currencyOperationColumnIndex.getColumnNumberByOperationType(rawOperationType)
			if isOperationTypeExists {
				sumStr := row[sumColumnNumber]
				sum := convertStringToFloat32(sumStr)
				operationDate := row[currencyOperationColumnIndex.date]
				operationTime := row[currencyOperationColumnIndex.time]
				if operationTime == "" {
					operationTime = "00:00:00"
				}
				datetimeStr := operationDate + " " + operationTime
				datetime, _ := time.Parse(dateLayout, datetimeStr)
				description := ""
				if len(row) > currencyOperationColumnIndex.description {
					description = row[currencyOperationColumnIndex.description]
				}
				operations = append(operations, rawOperation{
					datetime:      datetime,
					operationType: rawOperationType,
					sum:           sum,
					description:   description,
				})
			}
		}
	}
	return operations
}
