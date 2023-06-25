package tinkoff

const dateLayout = "02.01.2006 15:04:05"

var currencySet = map[string]bool{
	"RUB": true,
	"USD": true,
	"EUR": true,
	"CNY": true,
	"JPY": true,
}

type headNameValueCls struct {
	operations         string
	operationsWithCash string
	assetsInfo         string
	instrumentsInfo    string
}

var headNameValue = headNameValueCls{
	operations:         "1.1 Информация о совершенных и исполненных сделках на конец отчетного периода",
	operationsWithCash: "2. Операции с денежными средствами",
	assetsInfo:         "4.1 Информация о ценных бумагах",
	instrumentsInfo:    "4.2 Информация об инструментах, не квалифицированных в качестве ценной бумаги",
}
