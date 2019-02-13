package common

import (
	"time"
	"errors"
	"net/http"
	"io/ioutil"
	"strings"
	"strconv"
)

/*
华宝添益
120.79.154.53:4661/add?code=sh511990&height=100.05&low=99.905
白云机场
120.79.154.53:4661/addstock?code=sh600004&money=130000&copies=20&price=14.844
伊利股份
120.79.154.53:4661/addstock?code=sh600887&money=130000&copies=10&price=25.18
招商银行
120.79.154.53:4661/addstock?code=sh600036&money=130000&copies=10&price=29.32
复星医药
120.79.154.53:4661/addstock?code=sh600196&money=130000&copies=10&price=25.76
贵州茅台
120.79.154.53:4661/addstock?code=sh600519&money=130000&copies=10&price=717.92
*/

type InitNoticeStocksS struct {
	Name   string
	Code   string
	Height float64
	Low    float64
	Money  float64
}

type InitBuyStocksS struct {
	Name   string
	Code   string
	Money  float64
	Copies int
	Price  float64
}

func GetPriceFromUrl(url string) (t float64, err error) {
	h := time.Now().Hour()
	if h < 9 || h >= 15 {
		return 0, errors.New("time not useful")
	}

	response, _ := http.Get(url)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	headS := strings.Split(string(body), `=`)
	bodyS := strings.Split(headS[1], `,`)

	currentPrice, err := strconv.ParseFloat(bodyS[3], 64)
	if err != nil {
		panic(" strconv.ParseFloat error " + err.Error())
	}
	return currentPrice, nil
	//return TestData()
}

var testData = []float64{
	12.18,
	12.18,
	12.18,
	//11.18,
	//11.18,
	//11.18,
	//11.00,
	//11.00,
	//11.00,
	//10.76,
	//10.76,
	//10.76,
	//10.01,
	//10.01,
	//10.01,
	//9.99,
	//9.99,
	//9.99,
	//9.76,
	//9.76,
	//9.76,
	9.66,
	17.56,
	17.56,
	17.56,
	17.56,
	17.56,
	17.56,
	17.56,
	17.56,
	17.56,
}

func TestData() float64 {
	tt := testData[0]
	testData = testData[1:]
	return tt
}
