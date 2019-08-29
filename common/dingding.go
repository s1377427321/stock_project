package common

import (
	"fmt"
	"net/http"
	"bytes"
	"github.com/astaxie/beego"
)

var Dingdingurl1 string = "https://oapi.dingtalk.com/robot/send?access_token=f6d35e625ec61e88b58d77dc98d5ba7c1342349aa882892347b13fcba98c5530"

func DingDingNotify1(content string)  {
	NotifyDingDing(content,Dingdingurl1)
}

func NotifyDingDing(content string, url string) {
	formt := `
        {
            "msgtype": "text", 
            "text": {
        		"content": "%s"
			}, 
			 "at": {
        		"isAtAll": true
    		}	
        }`
	body := fmt.Sprintf(formt, content)

	jsonValue := []byte(body)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	beego.Info("notifyDingDingIllegal manager server %v  %v  %v", resp, err)
	if (err != nil) {
		beego.Error(nil, "notifyDingDingIllegal %v  %v  %v", err, resp, body)
		return
	}
}
