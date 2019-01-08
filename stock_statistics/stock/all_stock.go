package stock

import (
	. "stock_statistics/constant"

	"net/http"
	"regexp"
	"time"
	"stock_statistics/storage"
	"strconv"
	"fmt"
	"stock_statistics/model"
	"io/ioutil"
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
		id,_:=strconv.Atoi(m[1])
		s := &model.Stock{
			Id:id,
			Code:   m[1],
			CnName: m[2],
			PyName: m[3],
			Type:   model.Code2Type(m[1]),
		}
		stocks[s.Code] = s
	}
	return stocks, nil

}

func (this *AllStock) UpdateStorage() {
	s := make([]*model.Stock, len(this.Stocks))
	i := 0
	for _, v := range this.Stocks {
		s[i] = v
		i += 1
	}
	storage.UpSertStockInfo(s)
}
