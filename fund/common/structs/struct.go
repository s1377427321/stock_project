package structs

import "fmt"

//163获取股票历史记录
var DAY_TRADE_API = "http://quotes.money.163.com/service/chddata.html?code=%s&start=%s&end=%s"

var HTTP_URL = "http://api.tushare.pro"

var TOKEN = "6ad2bcae0a39b5feab53acb555a149080db53e4f0640492485cbb8ce"

var SQLParams = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", "root", "123456", "120.79.154.53:3306", "stock")

type ReqParams struct {
	ApiName string `json:"api_name"`
	Token   string `json:"token"`
	Params  Params `json:"params"`
	Fields  string `json:"fields"`
}

type Params struct {
	Ts_Code     string `json:"ts_code"`
	Trade_Date  string `json:"trade_date"`
	Start_Date  string `json:"start_date"`
	End_Date    string `json:"end_date"`
	Ann_Date    string `json:"ann_date"`
	Period      string `json:"period"`
	Report_Type string `json:"report_type"`
	Comp_Type   string `json:"comp_type"`
	List_Status string `json:"list_status"`
	Exchange    string `json:"exchange"`
	Test        string `json:"test"`
}

type Respons struct {
	Request_Id string `json:"request_id"`
	Code       int    `json:"code"`
	Msg        bool   `json:"msg"`
	Data       string `json:"data"`
}

type StockItem struct {
	Code    string
	Name    string
	BuyNums float64
}

//buy_statics数据库处理字段
type BuyStaticsInfo struct {
	Code                     string  `json:"code"`
	CodeName                 string  `json:"code_name"`
	FixedInvestmentMoney     float64 `json:"fixed_investment_money"`
	FirstProfitabilityRate   float64 `json:"first_profitability_rate"`
	CurrentProfitabilityRate float64 `json:"current_profitability_rate"`
	Pe                       float64 `json:"pe"`
	Pb                       float64 `json:"pb"`
	DividentYield            float64 `json:"divident_yield"`
	Roe                      float64 `json:"roe"`
	SecondPower              float64 `json:"second_power"`
}

type StockBasicInfo struct {
	Id         int    `json:"id"`
	TsCode     string `json:"ts_code"`     //TS代码
	Symbol     string `json:"symbol"`      //股票代码
	Name       string `json:"name"`        //股票名称
	Area       string `json:"area"`        //所在地域
	Industry   string `json:"industry"`    //str	所属行业
	Fullname   string `json:"fullname"`    //str	股票全称
	Enname     string `json:"enname"`      //英文全称
	Market     string `json:"market"`      //市场类型 （主板/中小板/创业板）
	Exchange   string `json:"exchange"`    //str	交易所代码
	CurrType   string `json:"curr_type"`   //str	交易货币
	ListStatus string `json:"list_status"` //str	上市状态： L上市 D退市 P暂停上市
	ListDate   string `json:"list_date"`   //str	上市日期
	DelistDate string `json:"delist_date"` //str	退市日期
	IsHs       string `json:"is_hs"`       //是否沪深港通标的，N否 H沪股通 S深股通
}

//每日数据
type DialyStockInfo struct {
	Id     int     `json:"id"`
	Code   int     `json:"code"`
	Date   string  `json:"date"`
	Open   float64 `json:"open"`
	Close  float64 `json:"close"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Volume int     `json:"volume"`
	Money  int     `json:"money"`
}

type IncomeS struct {
	Data      []string
	Time      int
	PerInCome float64
}

type IncomeSS []*IncomeS

func (s IncomeSS) Len() int {
	return len(s)
}

func (s IncomeSS) Less(i, j int) bool {
	return s[i].Time > s[j].Time
}

//Swap()
func (s IncomeSS) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}