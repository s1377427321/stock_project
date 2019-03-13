package common
//
//var HTTP_URL = "http://api.tushare.pro"
//
//var TOKEN = "6ad2bcae0a39b5feab53acb555a149080db53e4f0640492485cbb8ce"
//
//type ReqParams struct {
//	ApiName string `json:"api_name"`
//	Token   string `json:"token"`
//	Params  Params `json:"params"`
//	Fields  string `json:"fields"`
//}
//
//type Params struct {
//	Ts_Code     string `json:"ts_code"`
//	Trade_Date  string `json:"trade_date"`
//	Start_Date  string `json:"start_date"`
//	End_Date    string `json:"end_date"`
//	Ann_Date    string `json:"ann_date"`
//	Period      string `json:"period"`
//	Report_Type string `json:"report_type"`
//	Comp_Type   string `json:"comp_type"`
//	Test        string `json:"test"`
//}
//
//type Respons struct {
//	Request_Id string `json:"request_id"`
//	Code       int    `json:"code"`
//	Msg        bool   `json:"msg"`
//	Data       string `json:"data"`
//}
//
//type StockItem struct {
//	Code    string
//	Name    string
//	BuyNums float64
//}
//
//
//type IncomeS struct {
//	Data []string
//	Time int
//	PerInCome float64
//
//}
//
//type IncomeSS []*IncomeS
//
//func (s IncomeSS) Len() int {
//	return len(s)
//}
//
//func (s IncomeSS) Less(i, j int) bool {
//	return s[i].Time > s[j].Time
//}
//
////Swap()
//func (s IncomeSS) Swap(i, j int) {
//	s[i], s[j] = s[j], s[i]
//}