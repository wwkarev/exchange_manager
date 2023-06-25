package tinkoff

import "github.com/xuri/excelize"

type assetInfoColumnIndexCls struct {
	marketplace int
	ticker      int
	isin        int
}

func initAssetInfoColumnIndex() assetInfoColumnIndexCls {
	marketplace, _ := excelize.ColumnNameToNumber("U")
	ticker, _ := excelize.ColumnNameToNumber("AG")
	isin, _ := excelize.ColumnNameToNumber("AT")

	return assetInfoColumnIndexCls{
		marketplace: marketplace - 1,
		ticker:      ticker - 1,
		isin:        isin - 1,
	}
}

var assetInfoColumnIndex = initAssetInfoColumnIndex()

type assetInfoColumnNameCls struct {
	ticker string
	isin   string
}

var assetInfoColumnName = assetInfoColumnNameCls{
	ticker: "Код актива",
	isin:   "ISIN",
}
