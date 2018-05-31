package stock

import (
	. "constant"
	"fmt"
	"io/ioutil"
	//	"log"
	"model"
	"net/http"
	"regexp"
	"time"
)

type AllStock struct {
	Stocks     map[string]*model.Stock
	UpdateTime int64
}

func NewAllStock() *AllStock {
	return &AllStock{}
}

func (this *AllStock) UpdateFromApi() {
	stocks, err := GetAllStock()
	if err != nil {
		fmt.Errorf("UpdateFromApi", err)
	}

	this.Stocks = stocks
	this.UpdateTime = time.Now().Unix()
}

var reg = regexp.MustCompile("~(?P<code>\\d+)`(?P<name>.*?)`(?P<py>\\w+)")

func GetAllStock() (map[string]*model.Stock, error) {
	resp, err := http.Get(ALL_STOCK_API)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	matches := reg.FindAllStringSubmatch(string(body), -1)
	stocks := make(map[string]*model.Stock)
	for _, m := range matches {
		s := &model.Stock{
			Code:   m[1],
			CnName: m[2],
			PyName: m[3],
			Type:   model.Code2Type(m[1]),
		}
		stocks[s.Code] = s
	}
	return stocks, nil

}
