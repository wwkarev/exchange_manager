package tinkoff

import "github.com/xuri/excelize"

type OperationType struct {
	buy  string
	sell string
}

var operationType = OperationType{
	buy:  "Покупка",
	sell: "Продажа",
}

type OperationColumnIndex struct {
	dealNumber          int
	date                int
	time                int
	operationType       int
	ticker              int
	sum                 int
	quantity            int
	currency            int
	accruedInterestSum  int
	brokerCommissionSum int
	commissionTicker    int
}

func initOperation() OperationColumnIndex {
	dealNumber, _ := excelize.ColumnNameToNumber("A")
	date, _ := excelize.ColumnNameToNumber("I")
	time, _ := excelize.ColumnNameToNumber("M")
	operationType, _ := excelize.ColumnNameToNumber("AD")
	ticker, _ := excelize.ColumnNameToNumber("AM")
	sum, _ := excelize.ColumnNameToNumber("BB")
	quantity, _ := excelize.ColumnNameToNumber("AY")
	currency, _ := excelize.ColumnNameToNumber("BR")
	accruedInterestSum, _ := excelize.ColumnNameToNumber("BH")
	brokerCommissionSum, _ := excelize.ColumnNameToNumber("BX")
	commissionTicker, _ := excelize.ColumnNameToNumber("CD")

	return OperationColumnIndex{
		dealNumber:          dealNumber - 1,
		date:                date - 1,
		time:                time - 1,
		operationType:       operationType - 1,
		ticker:              ticker - 1,
		sum:                 sum - 1,
		quantity:            quantity - 1,
		currency:            currency - 1,
		accruedInterestSum:  accruedInterestSum - 1,
		brokerCommissionSum: brokerCommissionSum - 1,
		commissionTicker:    commissionTicker - 1,
	}
}

var operationColumnIndex = initOperation()
