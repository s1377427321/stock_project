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
)

func GetFields(fields []interface{}) []string {
	res := make([]string, 0)
	for _, v := range fields {
		//logs.Info(k,v)
		res = append(res, v.(string))
	}

	return res
}

func GetContent(con []interface{}) [][]string {
	res := make([][]string, 0)
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
			}
			items = append(items, value)

		}
		res = append(res, items)
		//fmt.Println(reflect.TypeOf(v))
		//logs.Info(k,v)
	}
	return res
}

func PostToUrl(apiName string, fields string, ps *Params, ) ([]string, [][]string, error) {
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

	logs.Info("----------", string(jsb))

	req := bytes.NewBuffer(jsb)

	body_type := "application/json;charset=utf-8"
	resp, err := http.Post(HTTP_URL, body_type, req)
	body, _ := ioutil.ReadAll(resp.Body)
	logs.Info("----------", string(body))

	var dat map[string]interface{}
	json.Unmarshal(body, &dat)
	if v, ok := dat["code"]; ok {
		if v.(float64) != 0 {
			return nil, nil, errors.New("request error :" + err.Error())
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
				items := GetContent(v.(map[string]interface{})["items"].([]interface{}))
				//logs.Info(f)
				//logs.Info(items)
				return f, items, nil
				//return  v.(map[string]interface{})["fields"].(interface{}),v.(map[string]interface{})["items"].([]interface{}),nil
			}
		}
	}
	return nil, nil, errors.New("NOT FIND")
}

//获取最新四个季度的每股收入
func GetLatestYearPerIncome(code string) float64 {
	now := time.Now().Format("20060102")
	logs.Info(now)
	fields, items, _ := GetIncome(code, "20170101", now)
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
		jdata, err := json.Marshal(v)
		if err != nil {
			panic("NO NO NO")
		}

		logs.Info(string(jdata))

		temp := &StockBasicInfo{}
		err = json.Unmarshal(jdata, temp)
		if err != nil {
			panic(err)
		}

		logs.Info(temp)

		resultS = append(resultS, temp)

	}

	return resultS, nil
}
