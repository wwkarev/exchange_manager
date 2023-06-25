package tinkoff

import "github.com/xuri/excelize"

const KeyDateLayout = "06-01-02_15_04_05"

type currencyOperationTypeCls struct {
	refill                string
	withdrawal            string
	exchange              string
	tariffCommission      string
	exchangeCommission    string
	bondRedemption        string
	partialBondRedemption string
	interest              string
	interestTax           string
	dividends             string
	dividendsTax          string
}

var currencyOperationType = currencyOperationTypeCls{
	refill:                "Пополнение счета",
	withdrawal:            "Вывод средств",
	exchange:              "Покупка/продажа",
	tariffCommission:      "Комиссия по тарифу",
	exchangeCommission:    "Комиссия за сделки",
	bondRedemption:        "Погашение облигации",
	partialBondRedemption: "Частичное погашение облигации (амортизация номинала)",
	interest:              "Выплата купонов",
	interestTax:           "Налог (купонный доход)",
	dividends:             "Выплата дивидендов",
	dividendsTax:          "Налог (дивиденды)",
}

type currencyOperationColumnIndexCls struct {
	operationType int
	date          int
	time          int
	refill        int
	withdrawal    int
	description   int
}

func initCurrencyOperation() currencyOperationColumnIndexCls {
	operationType, _ := excelize.ColumnNameToNumber("BH")
	date, _ := excelize.ColumnNameToNumber("AO")
	time, _ := excelize.ColumnNameToNumber("R")
	refill, _ := excelize.ColumnNameToNumber("CH")
	withdrawal, _ := excelize.ColumnNameToNumber("DC")
	description, _ := excelize.ColumnNameToNumber("EF")

	return currencyOperationColumnIndexCls{
		operationType: operationType - 1,
		date:          date - 1,
		time:          time - 1,
		refill:        refill - 1,
		withdrawal:    withdrawal - 1,
		description:   description - 1,
	}
}

func (currencyOperationColumnIndex currencyOperationColumnIndexCls) getColumnNumberByOperationType(operationType string) (int, bool) {
	columnNumber := -1
	switch operationType {
	case currencyOperationType.refill:
		columnNumber = currencyOperationColumnIndex.refill
	case currencyOperationType.withdrawal:
		columnNumber = currencyOperationColumnIndex.withdrawal
	case currencyOperationType.exchange:
		columnNumber = currencyOperationColumnIndex.withdrawal
	case currencyOperationType.tariffCommission:
		columnNumber = currencyOperationColumnIndex.withdrawal
	case currencyOperationType.exchangeCommission:
		columnNumber = currencyOperationColumnIndex.withdrawal
	case currencyOperationType.bondRedemption:
		columnNumber = currencyOperationColumnIndex.refill
	case currencyOperationType.partialBondRedemption:
		columnNumber = currencyOperationColumnIndex.refill
	case currencyOperationType.interest:
		columnNumber = currencyOperationColumnIndex.refill
	case currencyOperationType.interestTax:
		columnNumber = currencyOperationColumnIndex.withdrawal
	case currencyOperationType.dividends:
		columnNumber = currencyOperationColumnIndex.refill
	case currencyOperationType.dividendsTax:
		columnNumber = currencyOperationColumnIndex.withdrawal
	}
	return columnNumber, columnNumber != -1
}

var currencyOperationColumnIndex = initCurrencyOperation()
