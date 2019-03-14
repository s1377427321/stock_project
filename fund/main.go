package main

import (
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
	"fund/common"
)

var Engine *xorm.Engine

func init() {

}

func main() {
	common.GetDealBuyStaticsInfos()
	//common.GetDealBuyStaticsInfos()
	//var err error
	//Engine, err = xorm.NewEngine("mysql", structs.SQLParams)
	//if err != nil {
	//	panic(err)
	//}
	//defer Engine.Close()
	//
	//fmt.Println("start  ",time.Now().Format("2006-01-02 15:04:05.000"))
	////
	//////sb := &structs.StockBasicInfo{
	//////	TsCode:     "000001.SZ",
	//////	Symbol:     "000001",
	//////	Name:       "平安银行",
	//////	Area:       "深圳",
	//////	Industry:   "银行",
	//////	FullName:   "平安银行股份有限公司",
	//////	EnName:     "Ping An Bank Co., Ltd.",
	//////	Marker:     "主板",
	//////	Exchange:   "SZSE",
	//////	CurrType:   "CNY",
	//////	ListStatus: "L",
	//////	ListDate:   "19910403",
	//////	DelistDate: "19910403",
	//////	IsHs:       "S",
	//////}
	//////
	//////sql := fmt.Sprintf("replace into stock_basic(`ts_code`,`symbol`,`name`,`area`,`industry`,`fullname`,`enname`,`market`,`exchange`,`curr_type`,`list_status`,`list_date`,`delist_date`,`is_hs`) VALUES(\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\");", sb.TsCode, sb.Symbol, sb.Name, sb.Area, sb.Industry, sb.FullName, sb.EnName, sb.Marker, sb.Exchange, sb.CurrType, sb.ListStatus, sb.ListDate, sb.DelistDate, sb.IsHs)
	//////
	//////res, err := Engine.Exec(sql)
	//////if err != nil {
	//////	panic(err)
	//////}
	//////fmt.Println(res.RowsAffected())
	////
	//_, contect, _ := common.GetAllStocksCode()
	////fmt.Println(fileds)
	////fmt.Println(contect)
	////
	//allStocks := make([]*structs.StockBasicInfo, 0)
	//sql := "replace into stock_basic(`ts_code`,`symbol`,`name`,`area`,`industry`,`fullname`,`enname`,`market`,`exchange`,`curr_type`,`list_status`,`list_date`,`delist_date`,`is_hs`) VALUES"
	//insertSql:=make([]string,0)
	//for _, v := range contect {
	//	if v[0] == "" {
	//		logs.Error("Error ")
	//		continue
	//	}
	//
	//	sb := &structs.StockBasicInfo{
	//		TsCode:     v[0],
	//		Symbol:     v[1],
	//		Name:       v[2],
	//		Area:       v[3],
	//		Industry:   v[4],
	//		FullName:   v[5],
	//		EnName:     v[6],
	//		Marker:     v[7],
	//		Exchange:   v[8],
	//		CurrType:   v[9],
	//		ListStatus: v[10],
	//		ListDate:   v[11],
	//		DelistDate: v[12],
	//		IsHs:       v[13],
	//	}
	//	allStocks = append(allStocks, sb)
	//	sqlChild := fmt.Sprintf("(\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\")", sb.TsCode, sb.Symbol, sb.Name, sb.Area, sb.Industry, sb.FullName, sb.EnName, sb.Marker, sb.Exchange, sb.CurrType, sb.ListStatus, sb.ListDate, sb.DelistDate, sb.IsHs)
	//	insertSql = append(insertSql,sqlChild)
	//	//fmt.Print(res.RowsAffected())
	//}
	//
	//for i:=0;i<len(insertSql)-1;i++ {
	//	sql+=insertSql[i]+","
	//}
	//
	//sql+=insertSql[len(insertSql)-1]+";"
	//
	//effct,err:=Engine.Exec(sql)
	//if err!=nil {
	//	panic(err)
	//}
	//
	//en,_:=effct.RowsAffected()
	//
	//fmt.Println("SQL RowsAffected",en)
	//
	////go func() {
	////	for range time.NewTicker(10 * time.Second).C {
	////		err := Engine.Ping()
	////		if err != nil {
	////			panic(err)
	////		}
	////	}
	////}()
	//
	////common.GetShiYingLv()
	////sqls := fmt.Sprintf("select * from %s where code=%s", "buy_statics", "510880")
	////result, err := Engine.Query(sqls)
	////if err != nil {
	////	fmt.Println(err)
	////}
	//
	////CodeInfos := make([]*structs.BuyStaticsInfo, 0)
	////
	////for _, v := range result {
	////	newCf := &structs.BuyStaticsInfo{}
	////	newCf.Code = string(v["code"])
	////	newCf.CodeName = string(v["code_name"])
	////	fim, err := strconv.ParseFloat(string(v["fixed_investment_money"]), 64)
	////	if err != nil {
	////		//panic(err)
	////		newCf.FixedInvestMentMoney = 0
	////	}else {
	////		newCf.FixedInvestMentMoney = fim
	////	}
	////
	////
	////	fpr, err := strconv.ParseFloat(string(v["first_profitability_rate"]), 64)
	////	if err != nil {
	////		//panic(err)
	////		newCf.FirstProfitabilityRate = 0
	////	}else {
	////		newCf.FirstProfitabilityRate = fpr
	////	}
	////
	////
	////	cpr, err := strconv.ParseFloat(string(v["current_profitability_rate"]), 64)
	////	if err != nil {
	////		//panic(err)
	////		newCf.CurretProfitabilityRate = 0
	////	}else {
	////		newCf.CurretProfitabilityRate = cpr
	////	}
	////
	////
	////	pe, err := strconv.ParseFloat(string(v["pe"]), 64)
	////	if err != nil {
	////		//panic(err)
	////		newCf.Pe = 0
	////	}else {
	////		newCf.Pe = pe
	////	}
	////
	////
	////	pb, err := strconv.ParseFloat(string(v["pb"]), 64)
	////	if err != nil {
	////		//panic(err)
	////		newCf.Pb = 0
	////	}else {
	////		newCf.Pb = pb
	////	}
	////
	////
	////	dy, err := strconv.ParseFloat(string(v["dividend_yield"]), 64)
	////	if err != nil {
	////		//panic(err)
	////		newCf.DividentYield = 0
	////	}else {
	////		newCf.DividentYield = dy
	////	}
	////
	////
	////	roe, err := strconv.ParseFloat(string(v["roe"]), 64)
	////	if err != nil {
	////		//panic(err)
	////		newCf.Roe = 0
	////	}else {
	////		newCf.Roe = roe
	////	}
	////
	////
	////	//
	////	//
	////	//for kk, vv := range v {
	////	//	fmt.Println(kk)
	////	//	fmt.Println(string(vv))
	////	//}
	////	CodeInfos = append(CodeInfos, newCf)
	////}
	////
	////
	////for _, v := range CodeInfos {
	////	common.DealProfitabilityRate(v)
	////}
	//
	////fmt.Println(result)
	////fileds,contect,_:=common.GetStockBasic()
	////fmt.Println(fileds)
	////fmt.Println(contect)
	//
	////time.Sleep(1 * time.Hour)
	//
	////perIncome:= common.GetLatestYearPerIncome("601066.SH")
	//
	////logs.Info(perIncome)
	//
	////v:=common.GetDateClosePrice("603288.SH","20190311")
	////logs.Info(v)
	////common.GetLatestYearPerIncome("603288.SH")
	////var allIncome float64
	////var allMarketValue float64
	////
	////now := time.Now().Format("20060102")
	////nowInt,_:= strconv.Atoi(now)
	////nowInt -=1
	////nowS:=strconv.Itoa(nowInt)
	////logs.Info(now)
	////
	////for _,v:=range stocks.Hongli_etf{
	////	//if k>5 {
	////	//	break
	////	//}
	////	codeByte :=[]byte(v.Code)
	////	first:=codeByte[0]
	////	if first  == []byte("6")[0] {
	////		v.Code +=".SH"
	////	}else if first  == []byte("0")[0] {
	////		v.Code +=".SZ"
	////	}
	////
	////	perIncome:= common.GetLatestYearPerIncome(v.Code)
	////	allIncome += perIncome*v.BuyNums
	////
	////	currentV:=common.GetDateClosePrice(v.Code,nowS)
	////	allMarketValue+=v.BuyNums*currentV
	////}
	////
	////logs.Info(allIncome)
	////logs.Info(allMarketValue)
	////logs.Info(allIncome/allMarketValue)
}
