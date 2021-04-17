package structs

type Response struct {
	Quotesummary struct {
		Result []struct {
			Summarydetail struct {
				Maxage    int `json:"maxAge"`
				Pricehint struct {
					Raw     int    `json:"raw"`
					Fmt     string `json:"fmt"`
					Longfmt string `json:"longFmt"`
				} `json:"priceHint"`
				Previousclose struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"previousClose"`
				Open struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"open"`
				Daylow struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"dayLow"`
				Dayhigh struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"dayHigh"`
				Regularmarketpreviousclose struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"regularMarketPreviousClose"`
				Regularmarketopen struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"regularMarketOpen"`
				Regularmarketdaylow struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"regularMarketDayLow"`
				Regularmarketdayhigh struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"regularMarketDayHigh"`
				Dividendrate struct {
				} `json:"dividendRate"`
				Dividendyield struct {
				} `json:"dividendYield"`
				Exdividenddate struct {
				} `json:"exDividendDate"`
				Payoutratio struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"payoutRatio"`
				Fiveyearavgdividendyield struct {
				} `json:"fiveYearAvgDividendYield"`
				Beta struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"beta"`
				Trailingpe struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"trailingPE"`
				Forwardpe struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"forwardPE"`
				Volume struct {
					Raw     int    `json:"raw"`
					Fmt     string `json:"fmt"`
					Longfmt string `json:"longFmt"`
				} `json:"volume"`
				Regularmarketvolume struct {
					Raw     int    `json:"raw"`
					Fmt     string `json:"fmt"`
					Longfmt string `json:"longFmt"`
				} `json:"regularMarketVolume"`
				Averagevolume struct {
					Raw     int    `json:"raw"`
					Fmt     string `json:"fmt"`
					Longfmt string `json:"longFmt"`
				} `json:"averageVolume"`
				Averagevolume10Days struct {
					Raw     int    `json:"raw"`
					Fmt     string `json:"fmt"`
					Longfmt string `json:"longFmt"`
				} `json:"averageVolume10days"`
				Averagedailyvolume10Day struct {
					Raw     int    `json:"raw"`
					Fmt     string `json:"fmt"`
					Longfmt string `json:"longFmt"`
				} `json:"averageDailyVolume10Day"`
				Bid struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"bid"`
				Ask struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"ask"`
				Bidsize struct {
					Raw     int    `json:"raw"`
					Fmt     string `json:"fmt"`
					Longfmt string `json:"longFmt"`
				} `json:"bidSize"`
				Asksize struct {
					Raw     int    `json:"raw"`
					Fmt     string `json:"fmt"`
					Longfmt string `json:"longFmt"`
				} `json:"askSize"`
				Marketcap struct {
					Raw     int64  `json:"raw"`
					Fmt     string `json:"fmt"`
					Longfmt string `json:"longFmt"`
				} `json:"marketCap"`
				Yield struct {
				} `json:"yield"`
				Ytdreturn struct {
				} `json:"ytdReturn"`
				Totalassets struct {
				} `json:"totalAssets"`
				Expiredate struct {
				} `json:"expireDate"`
				Strikeprice struct {
				} `json:"strikePrice"`
				Openinterest struct {
				} `json:"openInterest"`
				Fiftytwoweeklow struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"fiftyTwoWeekLow"`
				Fiftytwoweekhigh struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"fiftyTwoWeekHigh"`
				Pricetosalestrailing12Months struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"priceToSalesTrailing12Months"`
				Fiftydayaverage struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"fiftyDayAverage"`
				Twohundreddayaverage struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"twoHundredDayAverage"`
				Trailingannualdividendrate struct {
				} `json:"trailingAnnualDividendRate"`
				Trailingannualdividendyield struct {
				} `json:"trailingAnnualDividendYield"`
				Navprice struct {
				} `json:"navPrice"`
				Currency     string      `json:"currency"`
				Fromcurrency interface{} `json:"fromCurrency"`
				Tocurrency   interface{} `json:"toCurrency"`
				Lastmarket   interface{} `json:"lastMarket"`
				Volume24Hr   struct {
				} `json:"volume24Hr"`
				Volumeallcurrencies struct {
				} `json:"volumeAllCurrencies"`
				Circulatingsupply struct {
				} `json:"circulatingSupply"`
				Algorithm interface{} `json:"algorithm"`
				Maxsupply struct {
				} `json:"maxSupply"`
				Startdate struct {
				} `json:"startDate"`
				Tradeable bool `json:"tradeable"`
			} `json:"summaryDetail"`
			Price struct {
				Maxage                 int `json:"maxAge"`
				Premarketchangepercent struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"preMarketChangePercent"`
				Premarketchange struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"preMarketChange"`
				Premarkettime  int `json:"preMarketTime"`
				Premarketprice struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"preMarketPrice"`
				Premarketsource  string `json:"preMarketSource"`
				Postmarketchange struct {
				} `json:"postMarketChange"`
				Postmarketprice struct {
				} `json:"postMarketPrice"`
				Regularmarketchangepercent struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"regularMarketChangePercent"`
				Regularmarketchange struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"regularMarketChange"`
				Regularmarkettime int `json:"regularMarketTime"`
				Pricehint         struct {
					Raw     int    `json:"raw"`
					Fmt     string `json:"fmt"`
					Longfmt string `json:"longFmt"`
				} `json:"priceHint"`
				Regularmarketprice struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"regularMarketPrice"`
				Regularmarketdayhigh struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"regularMarketDayHigh"`
				Regularmarketdaylow struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"regularMarketDayLow"`
				Regularmarketvolume struct {
					Raw     int    `json:"raw"`
					Fmt     string `json:"fmt"`
					Longfmt string `json:"longFmt"`
				} `json:"regularMarketVolume"`
				Averagedailyvolume10Day struct {
					Raw     int    `json:"raw"`
					Fmt     string `json:"fmt"`
					Longfmt string `json:"longFmt"`
				} `json:"averageDailyVolume10Day"`
				Averagedailyvolume3Month struct {
					Raw     int    `json:"raw"`
					Fmt     string `json:"fmt"`
					Longfmt string `json:"longFmt"`
				} `json:"averageDailyVolume3Month"`
				Regularmarketpreviousclose struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"regularMarketPreviousClose"`
				Regularmarketsource string `json:"regularMarketSource"`
				Regularmarketopen   struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"regularMarketOpen"`
				Strikeprice struct {
				} `json:"strikePrice"`
				Openinterest struct {
				} `json:"openInterest"`
				Exchange              string      `json:"exchange"`
				Exchangename          string      `json:"exchangeName"`
				Exchangedatadelayedby int         `json:"exchangeDataDelayedBy"`
				Marketstate           string      `json:"marketState"`
				Quotetype             string      `json:"quoteType"`
				Symbol                string      `json:"symbol"`
				Underlyingsymbol      interface{} `json:"underlyingSymbol"`
				Shortname             string      `json:"shortName"`
				Longname              string      `json:"longName"`
				Currency              string      `json:"currency"`
				Quotesourcename       string      `json:"quoteSourceName"`
				Currencysymbol        string      `json:"currencySymbol"`
				Fromcurrency          interface{} `json:"fromCurrency"`
				Tocurrency            interface{} `json:"toCurrency"`
				Lastmarket            interface{} `json:"lastMarket"`
				Volume24Hr            struct {
				} `json:"volume24Hr"`
				Volumeallcurrencies struct {
				} `json:"volumeAllCurrencies"`
				Circulatingsupply struct {
				} `json:"circulatingSupply"`
				Marketcap struct {
					Raw     int64  `json:"raw"`
					Fmt     string `json:"fmt"`
					Longfmt string `json:"longFmt"`
				} `json:"marketCap"`
			} `json:"price"`
		} `json:"result"`
		Error interface{} `json:"error"`
	} `json:"quoteSummary"`
}
