package tinkoff

import (
	prsr "parser"
	"parser/tinkoff/head_row"
	"sort"
)

type UniversalOperationParser struct {
	operationParsers []prsr.OperationParser
}

func InitUniversalOperationParser(rows [][]string, headRowManager head_row.HeadRowManager) *UniversalOperationParser {
	currencyOperationsMap := (*InitRawOperationParser(rows, headRowManager)).GetCurrencyOperationsMap()
	isinTickerMap := (*initAssetInfoParser(rows, headRowManager)).GetIsinTickerMap()

	parsers := []prsr.OperationParser{
		initRefillOperationParser(currencyOperationsMap),
		initWithdrawalOperationParser(currencyOperationsMap),
		initTariffOperationParser(currencyOperationsMap),
		initDividendOperationParser(currencyOperationsMap, isinTickerMap),
		initDividendTaxOperationParser(currencyOperationsMap, isinTickerMap),
		initBondInterestOperationParser(currencyOperationsMap, isinTickerMap),
		initBondInterestTaxOperationParser(currencyOperationsMap, isinTickerMap),
		initBondRedemptionOperationParser(currencyOperationsMap, isinTickerMap),
		initPartialBondRedemptionOperationParser(currencyOperationsMap, isinTickerMap),
		initExchangeOperationParser(rows, headRowManager),
	}
	return &UniversalOperationParser{parsers}
}

func (parser UniversalOperationParser) GetOperations() []prsr.Operation {
	var operations []prsr.Operation
	for _, p := range parser.operationParsers {
		operations = append(operations, p.GetOperations()...)
	}
	sort.Slice(operations, func(i, j int) bool {
		return operations[i].Datetime.Before(operations[j].Datetime)
	})
	return operations
}
