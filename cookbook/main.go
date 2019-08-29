package main

import (
	"time"
	"cookbook/cook"
	"common/email"
	"common"
)

var emailUrls = []string{"1377427321@qq.com"}
var emailServer = &email.EmailServers{
	ServerEmail:    "1377427321@qq.com",
	ServerPassword: "atpncirernxrhchj",
	ServerPort:     465,
	ServerIP:       "smtp.qq.com",
}

func main() {

	cb := &cook.RandomCook{
		SelectCooks:    &cook.CookItem{},
		NotSelectCooks: &cook.CookItem{},
		CookBook:       &cook.CookItem{},
	}

	cb.SetCookBook("cookbook.config")

	loopEveryDay(cb)
	for range time.NewTicker(1 * time.Hour).C {
	}
}

func loopEveryDay(cb *cook.RandomCook) {
	for range time.NewTicker(50 * time.Minute).C {
		hour := time.Now().Hour()
		//day := time.Now().Day()
		if hour >= 18 && hour < 19 {
			//single := day % 2
			////随机菜谱
			//if single == 0 {
			//	cb.RandomCooks(4, 2, 1)
			//} else {
			//	cb.RandomCooks(4, 2, 0)
			//}
			res1 := cb.RandomCooks(4, 2, 1)

			res2 := cb.RandomCooks(4, 2, 1)

			endRes := res1 + "\n" + res2
			notice(endRes)
		}
	}

	//for range time.NewTicker(5 * time.Second).C {
	//	res := cb.RandomCooks(4, 2, 1)
	//	notice(res)
	//}
}

func notice(content string) {
	ec := &email.EmailContent{
		NickName:     "文健",
		Subject:      "CookEveryDay",
		BodyContent:  content,
		NoticeEmails: emailUrls,
	}
	email.SendEmailTo(emailServer, ec)
	common.DingDingNotify1(content)
}
