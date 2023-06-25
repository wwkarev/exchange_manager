package tinkoff

import (
	"parser/tinkoff/head_row"
)

type assetInfoParser struct {
	rows           [][]string
	headRowManager head_row.HeadRowManager
}

func initAssetInfoParser(rows [][]string, headRowManager head_row.HeadRowManager) *assetInfoParser {
	return &assetInfoParser{rows: rows, headRowManager: headRowManager}
}

func (parser assetInfoParser) GetIsinTickerMap() map[string]string {
	isinTickerMap := map[string]string{}
	headRow, exists := parser.headRowManager.GetHeadRowByValue(headNameValue.assetsInfo)
	if exists {
		nextHeadRow, exists := parser.headRowManager.GetNextHeadRow(headRow)
		if exists {
			for _, row := range parser.rows[headRow.Index:nextHeadRow.Index] {
				if len(row) > assetInfoColumnIndex.isin && len(row) > assetInfoColumnIndex.ticker {
					ticker := row[assetInfoColumnIndex.ticker]
					marketplace := row[assetInfoColumnIndex.marketplace]
					if ticker != "" && ticker != assetInfoColumnName.ticker && marketplace != "ВНБ" {
						isinTickerMap[row[assetInfoColumnIndex.isin]] = ticker
					}
				}
			}
		}
	}
	return isinTickerMap
}
