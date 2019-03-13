package common

import (
	"fmt"
	"strconv"
	"sort"
			"net/http"
	"io/ioutil"
	"strings"
	"github.com/astaxie/beego/logs"
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


func GetPriceFromUrl(url string) (t float64, err error) {

	response, _ := http.Get(url)
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

var testData = []float64{
	12.18,
	12.18,
	12.18,
	11.18,
	11.18,
	11.18,
	11.00,
	11.00,
	11.00,
	10.76,
	10.76,
	10.76,
	10.76,
	10.76,
	10.76,
	10.01,
	10.01,
	10.01,
	9.99,
	9.99,
	9.99,
	9.99,
	9.99,
	9.99,
	9.76,
	9.76,
	9.76,
	9.76,
	9.76,
	9.76,
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