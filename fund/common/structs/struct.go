package structs

import "fmt"

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
	Code                    string
	CodeName                string
	FixedInvestMentMoney    float64
	FirstProfitabilityRate  float64
	CurretProfitabilityRate float64
	Pe                      float64
	Pb                      float64
	DividentYield           float64
	Roe                     float64

}

type StockBasicInfo struct {
	TsCode     string `json:"ts_code"`     //TS代码
	Symbol     string `json:"symbol"`      //股票代码
	Name       string `json:"name"`        //股票名称
	Area       string `json:"area"`        //所在地域
	Industry   string `json:"industry"`    //str	所属行业
	Fullname   string `json:"fullname"`   //str	股票全称
	Enname     string `json:"enname"`    //英文全称
	Market     string `json:"market"`      //市场类型 （主板/中小板/创业板）
	Exchange   string `json:"exchange"`    //str	交易所代码
	CurrType   string `json:"curr_type"`   //str	交易货币
	ListStatus string `json:"list_status"` //str	上市状态： L上市 D退市 P暂停上市
	ListDate   string `json:"list_date"`   //str	上市日期
	DelistDate string `json:"delist_date"` //str	退市日期
	IsHs       string `json:"is_hs"`       //是否沪深港通标的，N否 H沪股通 S深股通
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
