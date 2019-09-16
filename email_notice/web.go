package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"fmt"
	"strconv"
	"net/http"
	. "email_notice/common"
	"github.com/astaxie/beego/logs"
	"time"
	"strings"
)

func AddNoticeStock(code string, heightPrice, lowPrice, money float64) {
	mx.Lock()
	var s *Stock
	var isNew = false
	if oldStock, ok := NoticeStockS[code]; ok {
		s = oldStock
	} else {
		s = &Stock{
			BuyMoney:       money,
			Code:           code,
			NoticeCallBack: NoticeEmail,
			NoticeLimit:    NoticeLimit,
		}
		NoticeStockS[code] = s
		isNew = true
	}

	logs.Info("---AddNoticeStock  ", code, heightPrice, lowPrice)
	if strings.Index(code, "0") == 0 {
		s.Url = fmt.Sprintf(mainUrl, "sz"+code)
	} else if strings.Index(code, "6") == 0 || strings.Index(code, "5") == 0 {
		s.Url = fmt.Sprintf(mainUrl, "sh"+code)
	}

	//s.Url = fmt.Sprintf(mainUrl, code)
	s.LowPrice = lowPrice
	s.HightPrice = heightPrice
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
//http://swjswj.vip:4661/add?code=sz000895&height=26.00&low=23.70&money=160000
func addNoticeStock(c echo.Context) error {
	code := c.QueryParam("code")
	heightPrice, _ := strconv.ParseFloat(c.QueryParam("height"), 64)
	lowPrice, _ := strconv.ParseFloat(c.QueryParam("low"), 64)
	money, _ := strconv.ParseFloat(c.QueryParam("money"), 64)

	AddNoticeStock(code, heightPrice, lowPrice, money)

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

func addStockStopWinLose(c echo.Context) error {
	code := c.QueryParam("code")
	name := c.QueryParam("name")
	price, _ := strconv.ParseFloat(c.QueryParam("price"), 64)

	magnification, _ := strconv.Atoi(c.QueryParam("magnification"))

	InstanceStopWinLoseManage().Add(code, name, price, magnification)

	result := InstanceStopWinLoseManage().ShowItems(code)

	return c.String(http.StatusOK, result+"\n  add OK"+time.Now().Format("2006-01-02 15:04:05"))
	//return c.String(http.StatusOK, "add OK "+time.Now().Format("2006-01-02 15:04:05"))
}

func deleteStockStopWinLose(c echo.Context) error {
	code := c.QueryParam("code")

	result := InstanceStopWinLoseManage().ShowItems(code)

	InstanceStopWinLoseManage().DeleteStock(code)
	return c.String(http.StatusOK, result+"\n  deleteStockStopWinLose OK "+time.Now().Format("2006-01-02 15:04:05"))
}

func addLowSuctionStopWinLose(c echo.Context) error {
	code := c.QueryParam("code")
	price, _ := strconv.ParseFloat(c.QueryParam("price"), 64)
	buy_num, _ := strconv.Atoi(c.QueryParam("buy_num"))
	InstanceStopWinLoseManage().AddLowSuction(code, buy_num, price)

	result := InstanceStopWinLoseManage().ShowItems(code)
	return c.String(http.StatusOK, result+"\n addLowSuctionStopWinLose OK "+time.Now().Format("2006-01-02 15:04:05"))
}

func reduceHighThrowStopWinLose(c echo.Context) error {
	code := c.QueryParam("code")
	price, _ := strconv.ParseFloat(c.QueryParam("price"), 64)
	buy_num, _ := strconv.Atoi(c.QueryParam("buy_num"))
	InstanceStopWinLoseManage().ReduceHighThrow(code, buy_num, price)

	result := InstanceStopWinLoseManage().ShowItems(code)

	return c.String(http.StatusOK, result+"\n reduceHighThrowStopWinLose OK "+time.Now().Format("2006-01-02 15:04:05"))
}

func showStocksStopWinLose(c echo.Context) error {
	code := c.QueryParam("code")
	result := InstanceStopWinLoseManage().ShowItems(code)
	if result == "" {
		return c.String(http.StatusOK, "Nothing "+time.Now().Format("2006-01-02 15:04:05"))
	}

	return c.HTML(http.StatusOK, result)
}

func showStockGrid(c echo.Context) error {
	code := c.QueryParam("code")
	name := c.QueryParam("name")
	buyPrice, _ := strconv.ParseFloat(c.QueryParam("buy_price"), 64)
	bearLose, _ := strconv.ParseFloat(c.QueryParam("bear_lose"), 64)
	allMoney, _ := strconv.ParseFloat(c.QueryParam("all_money"), 64)
	divide, _ := strconv.ParseFloat(c.QueryParam("divide"), 64)

	sellPrice := buyPrice * (1 - bearLose)

	newgrid := NewStockGrid(code, name, buyPrice,sellPrice, bearLose, allMoney, divide)

	newgrid.Do()

	//return c.JSON(http.StatusOK, newgrid.Grids)
	return c.HTML(http.StatusOK, newgrid.ToString())
}

func showStockGrid2(c echo.Context) error {
	code := c.QueryParam("code")
	name := c.QueryParam("name")
	buyPrice, _ := strconv.ParseFloat(c.QueryParam("buy_price"), 64)
	sellPrice, _ := strconv.ParseFloat(c.QueryParam("sell_price"), 64)
	allMoney, _ := strconv.ParseFloat(c.QueryParam("all_money"), 64)
	divide, _ := strconv.ParseFloat(c.QueryParam("divide"), 64)

	bearLose := (buyPrice - sellPrice) / buyPrice

	newgrid := NewStockGrid(code, name, buyPrice, sellPrice, bearLose, allMoney, divide)

	newgrid.Do()

	//return c.JSON(http.StatusOK, newgrid.Grids)
	return c.HTML(http.StatusOK, newgrid.ToString())
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

	e.GET("/", showStocksStopWinLose)
	e.GET("/addswl", addStockStopWinLose)
	e.GET("/deleteswl", deleteStockStopWinLose)
	e.GET("/addlowsuctionswl", addLowSuctionStopWinLose)
	e.GET("/reducehighthrowswl", reduceHighThrowStopWinLose)
	e.GET("/showstocksswl", showStocksStopWinLose)
	//e.GET("/getstockbuyNum",getStockBuyNum)

	e.GET("/showgrid", showStockGrid)
	e.GET("/showgrid2", showStockGrid2)

	fmt.Println("RunHttpServer ----------------- ")
	err := e.Start(httpPort)
	if err != nil {
		panic(err.Error())
	}
}
