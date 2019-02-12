package spider_struct

import (
	"time"
	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego/logs"
	"fmt"
)

type Spider struct {
	Url string
	isSendEmail bool
}

func (this *Spider)Start(times time.Duration)  {
	loopTime:= time.NewTicker(times)
	go func() {
		for range loopTime.C{
			this.SpiderUrl(this.Url)
		}
	}()
}

func (this *Spider)GetIsSendEmail() bool {
	return this.isSendEmail
}

func (this *Spider)SpiderUrl(url string) {
	doc, err :=goquery.NewDocument(url)

	if err!= nil {
		logs.Error(err)
		return
	}

	//doc.Find("strong[class=xp1 wryh zxj green]").Each(func(i int, selection *goquery.Selection) {
	//	fmt.Println(selection.Text())
	//})
	logs.Info("start ",this.Url)
	//doc.Find("div[class=\"fl xt1 data-left\"]").Each(func(i int, selection *goquery.Selection) {
	//	selection.Find("div[id=\"arrowud\"]").Each(func(i int, selection1 *goquery.Selection) {
	//
	//		fmt.Println(selection1.Text())
	//	})
	//})
	doc.Find("tr[id=512880]").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})

	logs.Info("end ",this.Url)

}

