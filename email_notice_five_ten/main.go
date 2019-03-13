package main

import (
	"fmt"
	"strings"
	"time"
	"net/http"
	"io/ioutil"
	"github.com/axgle/mahonia"
	"stock_statistics/model"
	"encoding/csv"
	"io"
	"strconv"
	"sort"
	"common"
)

var BeforeURL = "http://quotes.money.163.com/service/chddata.html?code=%s&start=%s&end=%s"
var mainUrl = "http://hq.sinajs.cn/list=%s"
var HowManyDay = 30

func main() {
	gg:=GetStockPriceData("sh600004", 30)
	fmt.Println(len(gg))
}

func GetStockPriceData(code string, dayNum int64) model.DayTradeSlice {
	dts:=GetBeforData(code,dayNum)
	u := fmt.Sprintf(mainUrl, code)
	currentPrice, err := common.GetPriceFromUrl(u)
	if err!=nil {
		return nil
	}
	tn:=time.Now().Format("20060102")
	temp:=&model.DayTrade{
		Date:   tn,
		Close:  float32(currentPrice),
	}
	dts = append(dts,temp)
	sort.Sort(dts)
	return dts
}

//获取多天的数据
func GetBeforData(code string, dayNum int64) model.DayTradeSlice {
	var url string
	//end := time.Now().Format("20060102")
	end := time.Unix(time.Now().Unix()-3600*int64(time.Now().Hour()+1), 0).Format("20060102")
	before := time.Unix(time.Now().Unix()-3600*24*dayNum, 0).Format("20060102")
	if strings.Contains(code, "sz") {
		s := strings.Split(code, "sz")
		url = fmt.Sprintf(BeforeURL, "1"+s[1], before, end)
	} else if strings.Contains(code, "sh") {
		s := strings.Split(code, "sh")
		url = fmt.Sprintf(BeforeURL, "0"+s[1], before, end)
	} else {
		return nil
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

	days := make(model.DayTradeSlice, 0)

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

		Close, _ := strconv.ParseFloat(cols[3], 32)
		Open, _ := strconv.ParseFloat(cols[6], 32)
		Low, _ := strconv.ParseFloat(cols[5], 32)
		High, _ := strconv.ParseFloat(cols[4], 32)
		Volume, _ := strconv.Atoi(cols[11])
		Money, _ := strconv.ParseFloat(cols[12], 32)
		Code, _ := strconv.Atoi(strings.TrimLeft(cols[1], "'"))

		if Open == 0 {
			continue
		}

		//获取本地location
		timeLayout := "2006-01-02" //转化所需模板
		targetLayout := "20060102"
		loc, _ := time.LoadLocation("Local")                         //重要：获取时区
		theTime, _ := time.ParseInLocation(timeLayout, cols[0], loc) //使用模板在对应时区转化为time.time类型
		timeN := theTime.Format(targetLayout)                        //转化为时间戳 类型是int64

		days = append(days, &model.DayTrade{
			Date:   timeN,
			Code:   Code,
			Close:  float32(Close),
			High:   float32(High),
			Low:    float32(Low),
			Open:   float32(Open),
			Volume: Volume,
			Money:  int(Money),
		})

	}

	sort.Sort(days)

	return days
}
