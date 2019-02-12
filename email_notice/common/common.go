package common

import (
	"time"
	"errors"
	"net/http"
	"io/ioutil"
	"strings"
	"strconv"
)

func GetPriceFromUrl(url string) (t float64, err error) {
	h := time.Now().Hour()
	if h<9 || h>=15 {
		return 0,errors.New("time not useful")
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
	17.56, 17.56,
	17.56,
	17.56,
}

func TestData() float64 {
	tt := testData[0]
	testData = testData[1:]
	return tt
}