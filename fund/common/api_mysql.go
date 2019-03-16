package common

import (
	"errors"
	"reflect"
	"fmt"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"bytes"
	"net/http"
	"io/ioutil"
	"time"
	"strconv"
	"sort"
	. "fund/common/structs"
	"fund/stocks"
	"github.com/go-xorm/xorm"
	"common"
	"log"
	"stock_statistics/constant"
	"github.com/axgle/mahonia"
	"encoding/csv"
	"strings"
	"io"

	commons "common"
)

var Engine *xorm.Engine

func init() {
	logs.Info("Start MYSQL")
	var err error
	Engine, err = xorm.NewEngine("mysql", SQLParams)
	if err != nil {
		panic(err)
	}
}

func GetFields(fields []interface{}) []string {
	res := make([]string, 0)
	for _, v := range fields {
		//logs.Info(k,v)
		res = append(res, v.(string))
	}

	return res
}

func GetContent(fields []interface{},con []interface{}) ([][]string,[]map[string]string) {
	res := make([][]string, 0)
	resMaps:= make([]map[string]string,0)
	for _, v := range con {
		items := make([]string, 0)
		for _, vv := range v.([]interface{}) {
			//logs.Info("+++++++++++++++++")
			//logs.Info(kk,vv,reflect.TypeOf(vv))
			if vv == nil {
				items = append(items, "null")
				continue
			}
			kind := reflect.TypeOf(vv).Kind()
			var value string
			switch kind {
			case reflect.Float64:
				convv := vv.(float64)
				value = fmt.Sprintf("%f", convv)
				//value = strconv.FormatFloat(convv,'E',-1,64)
				//float,_ := strconv.ParseFloat(value,64)
				//logs.Info("***** ",value)
			case reflect.String:
				value = vv.(string)
				if value == "" {
					panic("AAAAAAAAAAAAAA")
				}
				//logs.Info("***** ",value)
			case reflect.Int:
				convv := vv.(int)
				value = fmt.Sprintf("%d", convv)
				if value == "" {
					panic("AAAAAAAAAAAAAA")
				}
			default:
				panic("GetContent GetContent Error")
			}
			items = append(items, value)

		}

		resMap:=make(map[string]string,0)
		for i:=0;i<len(items);i++ {
			key:=fields[i].(string)
			resMap[key] = items[i]
		}

		res = append(res, items)
		resMaps = append(resMaps,resMap)

		//fmt.Println(reflect.TypeOf(v))
		//logs.Info(k,v)
	}
	return res,resMaps
}

func PostToUrl(apiName string, fields string, ps *Params, ) ([]string, [][]string,[]map[string]string, error) {
	//ps := common.Params{
	//	Ts_Code:    "002594.SZ",
	//	Trade_Date: "20180726",
	//}

	//rp := common.ReqParams{
	//	ApiName: "daily_basic",
	//	Token:   common.TOKEN,
	//	Params:  *ps,
	//	Fields:  "ts_code,pe,pb,close,total_share,float_share,free_share,total_mv",
	//}

	rp := ReqParams{
		ApiName: apiName,
		Token:   TOKEN,
		Params:  *ps,
		Fields:  fields,
	}

	jsb, err := json.Marshal(rp)
	if err != nil {
		panic("json error " + err.Error())
	}

	//logs.Info("----------", string(jsb))

	req := bytes.NewBuffer(jsb)

	body_type := "application/json;charset=utf-8"
	resp, err := http.Post(HTTP_URL, body_type, req)
	body, _ := ioutil.ReadAll(resp.Body)
	//logs.Info("----------", string(body))

	var dat map[string]interface{}
	json.Unmarshal(body, &dat)
	if v, ok := dat["code"]; ok {
		if v.(float64) != 0 {
			return nil, nil,nil, errors.New("request error :" + err.Error())
		}
	}

	for _, v := range dat {
		//logs.Info("----------", reflect.TypeOf(v))
		//logs.Info(k, v)
		if v != nil {
			//logs.Info(reflect.TypeOf(v).Kind())
			if reflect.TypeOf(v).Kind() == reflect.Map {
				//fmt.Println("AAAAAAAAAAAA")
				//
				//for vv,kk:=range v.(map[string]interface{})["items"].([]interface{}){
				//	fmt.Println("GGG  ",vv)
				//	fmt.Println("GGG  ",kk)
				//}
				//
				//
				//fmt.Println(v.(map[string]interface{})["items"].([]interface{}))
				//fmt.Println(reflect.TypeOf(v.(map[string]interface{})["items"]))
				//fmt.Println(reflect.TypeOf(v.(map[string]interface{})["items"]))
				//fmt.Println(v.(map[string]interface{})["fields"].([]interface{}))
				f := GetFields(v.(map[string]interface{})["fields"].([]interface{}))
				items,itmsMap := GetContent(v.(map[string]interface{})["fields"].([]interface{}),v.(map[string]interface{})["items"].([]interface{}))
				//logs.Info(f)
				//logs.Info(items)
				return f, items,itmsMap, nil
				//return  v.(map[string]interface{})["fields"].(interface{}),v.(map[string]interface{})["items"].([]interface{}),nil
			}
		}
	}
	return nil, nil,nil, errors.New("NOT FIND")
}

//获取最新四个季度的每股收入
func GetLatestYearPerIncome(code string) float64 {
	now := time.Now().Format("20060102")
	logs.Info(now)
	fields, items, _ ,_:= GetIncome(code, "20170101", now)
	ann_date_index := 0
	basic_eps := 0
	for k, v := range fields {
		if v == "ann_date" {
			ann_date_index = k

		}
		if v == "basic_eps" {
			basic_eps = k
		}
	}

	saveSlice := make(IncomeSS, 0)
	tempSave := make(map[int]int, 0)
	for _, v := range items {
		t, err := strconv.Atoi(v[ann_date_index])
		if err != nil {
			panic("time.Parse " + err.Error())
		}

		if _, ok := tempSave[t]; ok {
			continue
		}

		tempSave[t] = t
		newS := &IncomeS{
			Data: v,
			Time: t,
		}

		saveSlice = append(saveSlice, newS)
	}

	sort.Sort(saveSlice)

	for i := 0; i < len(saveSlice)-1; i++ {
		beforMoth := (saveSlice[i].Time / 100) % 100
		afterMoth := (saveSlice[i+1].Time / 100) % 100

		pc, err := strconv.ParseFloat(saveSlice[i].Data[basic_eps], 64)
		if err != nil {
			panic("ParseFloat  " + err.Error())
		}

		pc2, err := strconv.ParseFloat(saveSlice[i+1].Data[basic_eps], 64)
		if err != nil {
			panic("ParseFloat  " + err.Error())
		}

		if beforMoth >= 5 && afterMoth >= 4 {
			saveSlice[i].PerInCome = pc - pc2
		} else if beforMoth == 4 && afterMoth <= 4 {
			saveSlice[i].PerInCome = pc
		} else if beforMoth <= 4 && afterMoth >= 9 {
			saveSlice[i].PerInCome = pc - pc2
		} else {
			panic(fmt.Sprintf("beformaoth = %v ,afterMoth = %v", beforMoth, afterMoth))
		}

		logs.Info("+++ ", saveSlice[i].Data[0], saveSlice[i].Time, saveSlice[i].PerInCome)
	}

	var totolYearIncome float64 = 0
	for i := 0; i < 4; i++ {
		totolYearIncome += saveSlice[i].PerInCome
	}

	logs.Info("totolYearIncome =", totolYearIncome)
	logs.Info(fields)
	logs.Info(items)

	return totolYearIncome
}

//获取收盘价格
func GetDateClosePrice(code string, date string) float64 {
	fields, content, err := GetDailyBasic(code, date)
	if err != nil {
		panic("GetDateClosePrice  " + err.Error())
	}

	var closePriceIndex int
	for i := 0; i < len(fields); i++ {
		if fields[i] == "close" {
			closePriceIndex = i
		}
	}

	v, err := strconv.ParseFloat(content[0][closePriceIndex], 64)

	return v

}

//income  计算四个季度的动态市盈率
func GetShiYingLv() {
	var allIncome float64
	var allMarketValue float64

	now := time.Now().Format("20060102")
	nowInt, _ := strconv.Atoi(now)
	nowInt -= 1
	nowS := strconv.Itoa(nowInt)
	logs.Info(now)

	for _, v := range stocks.Hongli_etf {
		//if k>5 {
		//	break
		//}
		codeByte := []byte(v.Code)
		first := codeByte[0]
		if first == []byte("6")[0] {
			v.Code += ".SH"
		} else if first == []byte("0")[0] {
			v.Code += ".SZ"
		}

		perIncome := GetLatestYearPerIncome(v.Code)
		allIncome += perIncome * v.BuyNums

		currentV := GetDateClosePrice(v.Code, nowS)
		allMarketValue += v.BuyNums * currentV
	}

	logs.Info(fmt.Sprintf("%0.2f", allIncome))
	logs.Info(fmt.Sprintf("%0.2f", allMarketValue))
	logs.Info(fmt.Sprintf("%0.2f", allIncome/allMarketValue))
}

//插入所有股票信息到stock_basic数据库
func UpdateAllStocksAndInsertToSQL() {
	Engine, err := xorm.NewEngine("mysql", SQLParams)
	if err != nil {
		panic(err)
	}
	defer Engine.Close()

	fmt.Println("start  ", time.Now().Format("2006-01-02 15:04:05.000"))

	_, contect, _ := GetAllStocksCode()

	allStocks := make([]*StockBasicInfo, 0)
	sql := "replace into stock_basic(`ts_code`,`symbol`,`name`,`area`,`industry`,`fullname`,`enname`,`market`,`exchange`,`curr_type`,`list_status`,`list_date`,`delist_date`,`is_hs`) VALUES"
	insertSql := make([]string, 0)
	for _, v := range contect {
		if v[0] == "" {
			logs.Error("Error ")
			continue
		}

		sb := &StockBasicInfo{
			TsCode:     v[0],
			Symbol:     v[1],
			Name:       v[2],
			Area:       v[3],
			Industry:   v[4],
			Fullname:   v[5],
			Enname:     v[6],
			Market:     v[7],
			Exchange:   v[8],
			CurrType:   v[9],
			ListStatus: v[10],
			ListDate:   v[11],
			DelistDate: v[12],
			IsHs:       v[13],
		}
		allStocks = append(allStocks, sb)
		sqlChild := fmt.Sprintf("(\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\")", sb.TsCode, sb.Symbol, sb.Name, sb.Area, sb.Industry, sb.Fullname, sb.Enname, sb.Market, sb.Exchange, sb.CurrType, sb.ListStatus, sb.ListDate, sb.DelistDate, sb.IsHs)
		insertSql = append(insertSql, sqlChild)
		//fmt.Print(res.RowsAffected())
	}

	for i := 0; i < len(insertSql)-1; i++ {
		sql += insertSql[i] + ","
	}

	sql += insertSql[len(insertSql)-1] + ";"

	effct, err := Engine.Exec(sql)
	if err != nil {
		panic(err)
	}

	en, _ := effct.RowsAffected()

	fmt.Println("SQL RowsAffected", en)
}

//从mysql数据库中获取所有股票
func GetAllStocksInfo() ([]*StockBasicInfo, error) {
	resultS := make([]*StockBasicInfo, 0)

	Engine, err := xorm.NewEngine("mysql", SQLParams)
	if err != nil {
		panic(err)
	}
	defer Engine.Close()

	sqlS := fmt.Sprintf("SELECT * from stock_basic;")
	queryR, err := Engine.QueryString(sqlS)
	if err != nil {
		logs.Error(fmt.Sprintf("QueryString %s error :%s", sqlS, err.Error()))
		return nil, err
	}

	for _, v := range queryR {

		temp := &StockBasicInfo{}
		common.DataToStruct(v, temp)
		//logs.Info(temp)
		resultS = append(resultS, temp)
	}

	return resultS, nil
}

//trade_his  1  从网络获取个股股价等信息，插入表trade_his中
func UpdateDayTradeData() {
	stocks, err := GetAllStocksInfo()
	if err != nil {
		panic(err)
	}

	sem := make(chan int, 15)
	for _, stock := range stocks {
		go func(s *StockBasicInfo) {
			UpdateDayTradeDataNext1(s, sem)
		}(stock)
		select {
		case <-time.After(10 * time.Second):
			fmt.Println("========do time.After(1) ======")
			//delete(sem,chan)
		case <-sem:
			fmt.Println("========do self <-sem======")
		}

		//UpdateDayTradeDataNext1(stock, sem)
	}

}

//trade_his  2 插入操作
func UpdateDayTradeDataNext1(stock *StockBasicInfo, sem chan int) {
	defer commons.RecoverPanic()
	lastestStocksInfo := GetLatestDayStock(stock)

	now := time.Now().Format("20060102")
	var added []*DialyStockInfo
	if lastestStocksInfo == nil || lastestStocksInfo.Date == "" {
		before := time.Unix(time.Now().Unix()-3600*24*360*5, 0).Format("20060102")
		added = GetDailTradeFromCSV(stock, before, now)
	} else {
		last, _ := time.Parse("2006-01-02", lastestStocksInfo.Date)
		// init today
		from := time.Unix(last.Unix()+3600*24*1, 0).Format("20060102")
		if from == now {
			log.Println("not data need insert ", lastestStocksInfo.Code)
			sem <- 1
			return
		}

		added = GetDailTradeFromCSV(stock, from, now)
	}

	logs.Info("InsertTradeHis  ", stock.TsCode)

	if len(added) != 0 {
		InsertTradeHis(added)
	}

	logs.Info("SUCESS ", stock.TsCode)

	sem <- 1
}

//trade_his  3 获取最近一天记录到数据库的数据，为了更新数据用
func GetLatestDayStock(stock *StockBasicInfo) *DialyStockInfo {
	sqlStr := `
	SELECT * FROM trade_his WHERE code=%d ORDER BY date DESC LIMIT 1;
	`
	code, err := strconv.Atoi(stock.Symbol)
	if err != nil {
		panic(fmt.Sprintf("%v strconv.Atoi error %v", stock.Symbol, err))
	}

	sql := fmt.Sprintf(sqlStr, code)
	sqlResult, err := Engine.QueryString(sql)
	if err != nil || sqlResult == nil {
		return nil
	}

	ret := &DialyStockInfo{}

	commons.DataToStruct(sqlResult[0], ret)

	return ret
}

//trade_his  4 从163网上下载csv格式数据，然后保存到数据库中
func GetDailTradeFromCSV(stock *StockBasicInfo, begin, end string) []*DialyStockInfo {
	typeStock := []byte(stock.Symbol)[0]
	var url string
	if typeStock == []byte("0")[0] || typeStock == []byte("3")[0] {
		url = fmt.Sprintf(constant.DAY_TRADE_API, "1"+stock.Symbol, begin, end)
	} else if typeStock == []byte("6")[0] {
		url = fmt.Sprintf(constant.DAY_TRADE_API, "0"+stock.Symbol, begin, end)
	} else {
		panic(fmt.Sprintf("%v  error code id %v", "GetDailTradeFromCSV", stock.Symbol, ))
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil
	}
	enc := mahonia.NewDecoder("gbk")
	_, utf8Body, _ := enc.Translate(body, true)

	days := make([]*DialyStockInfo, 0)

	r := csv.NewReader(strings.NewReader(string(utf8Body)))
	r.Read()
	for {
		cols, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil
		}

		Close, _ := strconv.ParseFloat(cols[3], 64)
		Open, _ := strconv.ParseFloat(cols[6], 64)
		Low, _ := strconv.ParseFloat(cols[5], 64)
		High, _ := strconv.ParseFloat(cols[4], 64)
		Volume, _ := strconv.Atoi(cols[11])
		Money, _ := strconv.ParseFloat(cols[12], 64)
		Code, _ := strconv.Atoi(strings.TrimLeft(cols[1], "'"))

		if Open == 0 {
			continue
		}

		days = append(days, &DialyStockInfo{
			Date:   cols[0],
			Code:   Code,
			Close:  Close,
			High:   High,
			Low:    Low,
			Open:   Open,
			Volume: Volume,
			Money:  int(Money),
		})

	}

	return days
}

//trade_his  5 把数据插入到数据库中
func InsertTradeHis(stockInfos []*DialyStockInfo) {
	sql := "replace into trade_his(`code`,`date`,`open`,`close`,`high`,`low`,`volume`,`money`)VALUES"

	for i := 0; i < len(stockInfos); i++ {
		v := stockInfos[i]
		var tempStr string
		if i == len(stockInfos)-1 {
			tempStr = fmt.Sprintf("(%d,\"%s\",%f,%f,%f,%f,%d,%d);", v.Code, v.Date, v.Open, v.Close, v.High, v.Low, v.Volume, v.Money)
		} else {
			tempStr = fmt.Sprintf("(%d,\"%s\",%f,%f,%f,%f,%d,%d),", v.Code, v.Date, v.Open, v.Close, v.High, v.Low, v.Volume, v.Money)
		}
		sql += tempStr
	}

	_, err := Engine.Exec(sql)

	if err != nil {
		logs.Error(sql)
		panic(fmt.Sprintf("%v ERROR %v", "Engine.Exec", err))
	}
}
