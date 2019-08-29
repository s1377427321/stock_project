package common

import (
	"fmt"
	"strconv"
	"sort"
			"net/http"
	"io/ioutil"
	"strings"
	"github.com/astaxie/beego/logs"
	"errors"
	)

func Round64(f float64, n int) float64 {
	floatStr := fmt.Sprintf("%."+strconv.Itoa(n)+"f", f)
	inst, _ := strconv.ParseFloat(floatStr, 64)
	return inst
}


func GetMinMax(a,b,c float64) (aret float64 , bret float64) {
	temp:=[]float64{a,b,c}

	sort.Float64s(temp)

	return temp[0],temp[len(temp)-1]
}



func GetPriceFromUrlSZSH(code string) (t float64, err error) {

	url :=fmt.Sprintf("http://hq.sinajs.cn/list=%s",code)
	response, err := http.Get(url)
	if err !=nil {
		return 0,errors.New(fmt.Sprintf("%s  , %s",url,err))
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	headS := strings.Split(string(body), `=`)
	bodyS := strings.Split(headS[1], `,`)

	currentPrice, err := strconv.ParseFloat(bodyS[3], 64)
	if err != nil {
		panic(" strconv.ParseFloat error " + err.Error())
	}

	logs.Info(" GetPriceFromUrl ",url,currentPrice)
	return currentPrice, nil
	//return TestData(),nil
}


func GetPriceFromUrl(url string) (t float64, err error) {

	response, err := http.Get(url)
	if err !=nil {
		return 0,errors.New(fmt.Sprintf("%s  , %s",url,err))
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	headS := strings.Split(string(body), `=`)
	bodyS := strings.Split(headS[1], `,`)

	currentPrice, err := strconv.ParseFloat(bodyS[3], 64)
	if err != nil {
		panic(" strconv.ParseFloat error " + err.Error())
	}

	logs.Info(" GetPriceFromUrl ",url,currentPrice)
	return currentPrice, nil
	//return TestData(),nil
}

//var tempData []float64
//var startTemp bool
//
//func init()  {
//	tempData = make([]float64,0)
//	startTemp = false
//}
//func GetPriceFromUrl(url string) (t float64, err error) {
//	var result float64
//	if startTemp == true {
//		if len(tempData) >0 {
//			result=tempData[0]
//			testData = append(testData,result)
//			tempData = tempData[1:]
//		}else {
//			startTemp = false
//			result=testData[0]
//			tempData = append(tempData,result)
//			testData = testData[1:]
//		}
//	}else {
//		if len(testData) >0 {
//			result=testData[0]
//			tempData = append(tempData,result)
//			testData = testData[1:]
//		}else {
//			startTemp = true
//			result=tempData[0]
//			testData = append(testData,result)
//			tempData = tempData[1:]
//		}
//
//	}
//
//	beego.Info(result)
//	return result,nil
//}

var testData = []float64{
	10,
	10.01,
	10.20,
	10.30,
	10.36,
	10,51,
	10.64,
	10.33,
	10.23,
	10.11,
	10.02,
	9.89,
	9.72,
	9.69,
	9.78,
	9.99,
	10.11,
}

func TestData() float64 {
	tt := testData[0]
	testData = testData[1:]
	return tt
}