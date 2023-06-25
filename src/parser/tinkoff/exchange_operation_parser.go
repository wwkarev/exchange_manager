package tinkoff

import (
	"parser"
	"parser/tinkoff/head_row"
	"time"
)

type exchangeOperationParser struct {
	rows           [][]string
	headRowManager head_row.HeadRowManager
	startIndex     int
	endIndex       int
}

func initExchangeOperationParser(rows [][]string, headRowManager head_row.HeadRowManager) *exchangeOperationParser {
	operationParser := exchangeOperationParser{}
	operationParser.rows = rows
	operationParser.headRowManager = headRowManager
	operationParser.initOperationsInterval()
	return &operationParser
}

func (operationParser *exchangeOperationParser) initOperationsInterval() {
	headRow, exists := operationParser.headRowManager.GetHeadRowByValue(headNameValue.operations)
	if !exists {
		panic("No head row with value: " + headNameValue.operations)
	}
	startIndex := headRow.Index
	endIndex := len(operationParser.rows)
	nextHeadRow, exists := operationParser.headRowManager.GetNextHeadRow(headRow)
	if exists {
		endIndex = nextHeadRow.Index
	}
	operationParser.startIndex = startIndex
	operationParser.endIndex = endIndex
}

func (operationParser exchangeOperationParser) GetOperations() []parser.Operation {
	var operations []parser.Operation
	firstColumnValue := "Номер сделки"
	columnOffset := 0
	for i, row := range operationParser.rows[operationParser.startIndex+2 : operationParser.endIndex] {
		if row[0] == "" && row[1] == "" || row[0] == firstColumnValue {
			columnOffset = 0
			continue
		} else if row[0] == "" && row[1] == firstColumnValue {
			columnOffset = 1
			continue
		}

		key := string(i) + row[operationColumnIndex.dealNumber+columnOffset]
		datetimeStr := row[operationColumnIndex.date+columnOffset] + " " + row[operationColumnIndex.time+columnOffset]
		datetime, _ := time.Parse(dateLayout, datetimeStr)
		commissionTicker := row[operationColumnIndex.commissionTicker+columnOffset]
		brokerCommissionSum := convertStringToFloat32(row[operationColumnIndex.brokerCommissionSum+columnOffset])
		srcTickerIndex := operationColumnIndex.currency + columnOffset
		srcSumIndex := operationColumnIndex.sum + columnOffset
		destTickerIndex := operationColumnIndex.ticker + columnOffset
		destSumIndex := operationColumnIndex.quantity + columnOffset
		srcAccruedInterestTicker := row[srcTickerIndex]
		destAccruedInterestTicker := row[srcTickerIndex]
		srcAccruedIntersetSum := convertStringToFloat32(row[operationColumnIndex.accruedInterestSum+columnOffset])
		var destAccruedIntersetSum float32 = 0
		if row[operationColumnIndex.operationType+columnOffset] == operationType.sell {
			srcTickerIndex = operationColumnIndex.ticker + columnOffset
			srcSumIndex = operationColumnIndex.quantity + columnOffset
			destTickerIndex = operationColumnIndex.currency + columnOffset
			destSumIndex = operationColumnIndex.sum + columnOffset
			srcAccruedInterestTicker = row[destTickerIndex]
			destAccruedInterestTicker = row[destTickerIndex]
			srcAccruedIntersetSum = 0
			destAccruedIntersetSum = convertStringToFloat32(row[operationColumnIndex.accruedInterestSum+columnOffset])
		}
		srcTicker := row[srcTickerIndex]
		srcSum := convertStringToFloat32(row[srcSumIndex])
		destTicker := row[destTickerIndex]
		destSum := convertStringToFloat32(row[destSumIndex])
		operations = append(operations, parser.Operation{
			Key:                       key,
			Datetime:                  datetime,
			SrcOwner:                  parser.OwnerType.User,
			SrcTicker:                 srcTicker,
			SrcSum:                    srcSum,
			SrcAccruedInterestTicker:  srcAccruedInterestTicker,
			SrcAccruedInterestSum:     srcAccruedIntersetSum,
			DestOwner:                 parser.OwnerType.User,
			DestTicker:                destTicker,
			DestSum:                   destSum,
			DestAccruedInterestTicker: destAccruedInterestTicker,
			DestAccruedInterestSum:    destAccruedIntersetSum,
			CommissionTicker:          commissionTicker,
			CommissionSum:             brokerCommissionSum,
		})
	}
	return operations
}
