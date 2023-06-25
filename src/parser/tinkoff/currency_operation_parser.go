package tinkoff

import (
	prsr "parser"
)

type ownerCalculator interface {
	GetOwner(parser *currencyOperationParser, _rawOperation *rawOperation) string
}

type simpleOwnerCalculator struct {
	owner string
}

func (ownerCalculator simpleOwnerCalculator) GetOwner(parser *currencyOperationParser, _rawOperation *rawOperation) string {
	return ownerCalculator.owner
}

type descriptionOwnerCalculator struct {
	isinTickerMap map[string]string
}

func (ownerCalculator descriptionOwnerCalculator) GetOwner(parser *currencyOperationParser, _rawOperation *rawOperation) string {
	srcOwner := ""
	isin, isFound := getIsinFromDescription(_rawOperation.description)
	if isFound {
		srcOwner, _ = ownerCalculator.isinTickerMap[isin]
	}
	return srcOwner
}

type currencyOperationParser struct {
	currencyOperationsMap map[string][]rawOperation
	srcOwnerCalculator    ownerCalculator
	destOwnerCalculator   ownerCalculator
	desiredOperationType  string
}

func (parser currencyOperationParser) GetOperations() []prsr.Operation {
	var operations []prsr.Operation
	for currency, rawOperations := range parser.currencyOperationsMap {
		for _, _rawOperation := range rawOperations {
			if _rawOperation.operationType == parser.desiredOperationType {
				key := generateKey(_rawOperation.datetime, _rawOperation.sum)
				operations = append(operations, prsr.Operation{
					Key:                       key,
					Datetime:                  _rawOperation.datetime,
					SrcOwner:                  parser.srcOwnerCalculator.GetOwner(&parser, &_rawOperation),
					SrcTicker:                 currency,
					SrcSum:                    _rawOperation.sum,
					SrcAccruedInterestTicker:  currency,
					SrcAccruedInterestSum:     0,
					DestOwner:                 parser.destOwnerCalculator.GetOwner(&parser, &_rawOperation),
					DestTicker:                currency,
					DestSum:                   _rawOperation.sum,
					DestAccruedInterestTicker: currency,
					DestAccruedInterestSum:    0,
					CommissionTicker:          currency,
					CommissionSum:             0,
				})
			}
		}
	}

	return operations
}

type refillOperationParser struct {
	currencyOperationParser
}

func initRefillOperationParser(currencyOperationsMap map[string][]rawOperation) *refillOperationParser {
	return &refillOperationParser{currencyOperationParser{
		currencyOperationsMap: currencyOperationsMap,
		srcOwnerCalculator:    simpleOwnerCalculator{prsr.OwnerType.UserCache},
		destOwnerCalculator:   simpleOwnerCalculator{prsr.OwnerType.User},
		desiredOperationType:  currencyOperationType.refill,
	}}
}

type withdrawalOperationParser struct {
	currencyOperationParser
}

func initWithdrawalOperationParser(currencyOperationsMap map[string][]rawOperation) *withdrawalOperationParser {
	return &withdrawalOperationParser{currencyOperationParser{
		currencyOperationsMap: currencyOperationsMap,
		srcOwnerCalculator:    simpleOwnerCalculator{prsr.OwnerType.User},
		destOwnerCalculator:   simpleOwnerCalculator{prsr.OwnerType.UserCache},
		desiredOperationType:  currencyOperationType.withdrawal,
	}}
}

type tariffOperationParser struct {
	currencyOperationParser
}

func initTariffOperationParser(currencyOperationsMap map[string][]rawOperation) *tariffOperationParser {
	return &tariffOperationParser{currencyOperationParser{
		currencyOperationsMap: currencyOperationsMap,
		srcOwnerCalculator:    simpleOwnerCalculator{prsr.OwnerType.User},
		destOwnerCalculator:   simpleOwnerCalculator{prsr.OwnerType.Tariff},
		desiredOperationType:  currencyOperationType.tariffCommission,
	}}
}

type dividendOperationParser struct {
	currencyOperationParser
}

func initDividendOperationParser(currencyOperationsMap map[string][]rawOperation, isinTickerMap map[string]string) *dividendOperationParser {
	return &dividendOperationParser{
		currencyOperationParser{
			currencyOperationsMap: currencyOperationsMap,
			srcOwnerCalculator:    descriptionOwnerCalculator{isinTickerMap},
			destOwnerCalculator:   simpleOwnerCalculator{prsr.OwnerType.User},
			desiredOperationType:  currencyOperationType.dividends,
		},
	}
}

type dividendTaxOperationParser struct {
	currencyOperationParser
}

func initDividendTaxOperationParser(currencyOperationsMap map[string][]rawOperation, isinTickerMap map[string]string) *dividendTaxOperationParser {
	return &dividendTaxOperationParser{
		currencyOperationParser{
			currencyOperationsMap: currencyOperationsMap,
			srcOwnerCalculator:    simpleOwnerCalculator{prsr.OwnerType.User},
			destOwnerCalculator:   descriptionOwnerCalculator{isinTickerMap},
			desiredOperationType:  currencyOperationType.dividendsTax,
		},
	}
}

type bondInterestOperationParser struct {
	currencyOperationParser
}

func initBondInterestOperationParser(currencyOperationsMap map[string][]rawOperation, isinTickerMap map[string]string) *bondInterestOperationParser {
	return &bondInterestOperationParser{
		currencyOperationParser{
			currencyOperationsMap: currencyOperationsMap,
			srcOwnerCalculator:    descriptionOwnerCalculator{isinTickerMap},
			destOwnerCalculator:   simpleOwnerCalculator{prsr.OwnerType.User},
			desiredOperationType:  currencyOperationType.interest,
		},
	}
}

type bondInterestTaxOperationParser struct {
	currencyOperationParser
}

func initBondInterestTaxOperationParser(currencyOperationsMap map[string][]rawOperation, isinTickerMap map[string]string) *bondInterestTaxOperationParser {
	return &bondInterestTaxOperationParser{
		currencyOperationParser{
			currencyOperationsMap: currencyOperationsMap,
			srcOwnerCalculator:    simpleOwnerCalculator{prsr.OwnerType.User},
			destOwnerCalculator:   descriptionOwnerCalculator{isinTickerMap},
			desiredOperationType:  currencyOperationType.interestTax,
		},
	}
}

type bondRedemptionOperationParser struct {
	currencyOperationParser
}

func initBondRedemptionOperationParser(currencyOperationsMap map[string][]rawOperation, isinTickerMap map[string]string) *bondRedemptionOperationParser {
	return &bondRedemptionOperationParser{
		currencyOperationParser{
			currencyOperationsMap: currencyOperationsMap,
			srcOwnerCalculator:    descriptionOwnerCalculator{isinTickerMap},
			destOwnerCalculator:   simpleOwnerCalculator{prsr.OwnerType.User},
			desiredOperationType:  currencyOperationType.bondRedemption,
		},
	}
}

type partialBondRedemptionOperationParser struct {
	currencyOperationParser
}

func initPartialBondRedemptionOperationParser(currencyOperationsMap map[string][]rawOperation, isinTickerMap map[string]string) *partialBondRedemptionOperationParser {
	return &partialBondRedemptionOperationParser{
		currencyOperationParser{
			currencyOperationsMap: currencyOperationsMap,
			srcOwnerCalculator:    descriptionOwnerCalculator{isinTickerMap},
			destOwnerCalculator:   simpleOwnerCalculator{prsr.OwnerType.User},
			desiredOperationType:  currencyOperationType.partialBondRedemption,
		},
	}
}
