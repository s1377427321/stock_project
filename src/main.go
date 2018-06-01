package main

import (
	"flag"
	"fmt"
	"http_server"
	"log"
	"model"
	"stock"
	"time"
	"trade"
)

func update(helper *trade.DayTradeHelper, stock *model.Stock,sem chan int) {
	log.Println(stock.Code)
	helper.Update(stock)
	<-sem
}

func dailyTradeUpdate() {

	all := stock.NewAllStock()
	all.UpdateFromApi()
	helper := trade.NewDayTradeHelper()
	fmt.Println(helper)
	sem := make(chan int, 1)
	for _, stock := range all.Stocks {
		if stock.Type == model.HU_A ||
			stock.Type == model.SHEN_A ||
			stock.Type == model.CHUANGYE ||
			stock.Type == model.ZHONG_XIAO {

			sem <- 1
			go update(helper, stock, sem)
		}
	}

	//pool := commo.NewPool(20,10)
	//
	//for _,stock:=range all.Stocks{
	//	if stock.Type == model.HU_A ||
	//		stock.Type == model.SHEN_A ||
	//		stock.Type == model.CHUANGYE ||
	//		stock.Type == model.ZHONG_XIAO {
	//
	//		pool.JobQueue<- func() {
	//			update(helper, stock)
	//		}
	//
	//	}
	//}

	log.Println("----------dailyTradeUpdate  end----------------")

}

func main() {
	boolServer := flag.Bool("web", false, "start up a webserver to query stocks info")
	boolTrade := flag.Bool("dailyTrade", false, "update day trade info")
	boolStock := flag.Bool("dailyStock", false, "update stock code->name info")
	flag.Parse()
	if *boolServer {
		log.Println("aaaaaaaaaaaaa")
		//	httpServer()
	} else if *boolTrade {
		dailyTradeUpdate()
		log.Println("bbbbbbbbbbbbbbb")
	} else if *boolStock {
		log.Println("cccccccccccccccc")
		//	dailyAllStockUpdate()
	} else {
		fmt.Println("error input, use --help shows cmd")
	}

	http_server.RunHttpServer()
	for {
		time.Sleep(100 * time.Minute)
	}
}
