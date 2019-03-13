package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"fmt"
	"strconv"
	"net/http"
	. "email_notice/common"
	"github.com/astaxie/beego/logs"
)

func AddNoticeStock(code string, heightPrice, lowPrice, money float64) {
	mx.Lock()
	var s *Stock
	var isNew = false
	if oldStock, ok := NoticeStockS[code]; ok {
		s = oldStock
	} else {
		s = &Stock{
			BuyMoney:money,
			Code:           code,
			NoticeCallBack: NoticeEmail,
			NoticeLimit:    NoticeLimit,
		}
		NoticeStockS[code] = s
		isNew = true
	}

	logs.Info("---AddNoticeStock  ",code,heightPrice,lowPrice)

	s.Url = fmt.Sprintf(mainUrl, code)
	s.LowPrice = lowPrice
	s.HeightPrice = heightPrice
	s.Count = 0
	mx.Unlock()
	if isNew {
		s.Start()
	}
}

func AddBuyStock(code string, price, money float64, copies int) {
	bmx.Lock()
	var s *BuyStock
	var isNew = false
	if oldStock, ok := BuyStocks[code]; ok {
		s = oldStock
	} else {
		s = &BuyStock{
			StockName:                  code,
			BuyPrice:                   price,
			AllMoney:                   money,
			NumberOfCopies:             copies,
			NumberOfCopiesPice:         make(map[float64]*BuyLimit, 0),
			OrderNumberOfCopiesPiceKey: make([]float64, 0),
			AddNoticeFunc:              AddNoticeStock,
			DeleteNoticeFunc:           DeleteNoticeStock,
		}
		BuyStocks[code] = s
		isNew = true
	}

	s.StockUrl = fmt.Sprintf(mainUrl, code)

	bmx.Unlock()
	if isNew {
		s.Start()
	}
}

func DeleteBuyStock(code string) bool {
	if old, ok := BuyStocks[code]; ok {
		mx.Lock()
		old.Close()
		delete(BuyStocks, code)
		mx.Unlock()
		return true
	} else {
		return false
	}
}

func DeleteNoticeStock(code string) bool {
	if old, ok := NoticeStockS[code]; ok {
		mx.Lock()
		old.Close()
		delete(NoticeStockS, code)
		mx.Unlock()
		return true
	} else {
		return false
	}
}

//localhost:4661/delete?code=sh511990
func deleteNoticeStock(c echo.Context) error {
	code := c.QueryParam("code")

	isDelete := DeleteNoticeStock(code)

	if isDelete {
		return c.String(http.StatusOK, "Delete OK")
	} else {
		return c.String(http.StatusOK, "Delete Object Not Exist")
	}
}

//120.79.154.53:4661/add?code=sh511990&height=100.05&low=99.905&money=130000
//netstat -aon|findstr "40051"
//localhost:4661/add?code=sh511990&height=100.05&low=99&money=130000
//localhost:4661/add?code=sh511990&height=100.05&low=99.905&money=130000
func addNoticeStock(c echo.Context) error {
	code := c.QueryParam("code")
	heightPrice, _ := strconv.ParseFloat(c.QueryParam("height"), 64)
	lowPrice, _ := strconv.ParseFloat(c.QueryParam("low"), 64)
	money, _ := strconv.ParseFloat(c.QueryParam("money"), 64)

	AddNoticeStock(code, heightPrice, lowPrice,money)

	return c.String(http.StatusOK, "Add OK")
}

//localhost:4661/addstock?code=sh600004&money=130000&copies=20&price=14.844
func addStock(c echo.Context) error {
	code := c.QueryParam("code")
	price, _ := strconv.ParseFloat(c.QueryParam("price"), 64)
	money, _ := strconv.ParseFloat(c.QueryParam("money"), 64)
	copies, _ := strconv.Atoi(c.QueryParam("copies"))

	AddBuyStock(code, price, money, copies)

	return c.String(http.StatusOK, "Add Stock OK")
}

//localhost:4661/deletestock?code=sh600004
func deleteStock(c echo.Context) error {
	code := c.QueryParam("code")

	isSuccess := DeleteBuyStock(code)

	if isSuccess {
		return c.String(http.StatusOK, "Delete OK")
	} else {
		return c.String(http.StatusOK, "Delete Object Not Exist")
	}

}

func RunHttpServer() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	e.Use(middleware.Secure())
	e.Use(middleware.BodyLimit("1M"))

	e.Static("/", "./assets")

	e.GET("/add", addNoticeStock)
	e.GET("/delete", deleteNoticeStock)
	e.GET("/addstock", addStock)
	e.GET("/deletestock", deleteStock)
	fmt.Println("RunHttpServer ----------------- ")
	err := e.Start(httpPort)
	fmt.Println(err)
	panic(err.Error())
}
