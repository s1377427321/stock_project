package email

import (
	"log"
	"math/rand"
	"github.com/astaxie/beego"
	"fmt"
)

type EmailContent struct {
	NoticeEmails []string
	NickName string
	Subject string
	BodyContent string
}


type EmailServers struct {
	ServerEmail string
	ServerPassword string
	ServerPort int
	ServerIP string
}

func SendEmailTo(server *EmailServers, ec *EmailContent)  {
	client := New(server.ServerEmail,server.ServerPassword, ec.NickName, server.ServerIP, server.ServerPort, true)
	if err := client.SendEmail(ec.NoticeEmails,ec.Subject,ec.BodyContent);err !=nil{
		beego.Error(err.Error() +" ...."+server.ServerEmail)
	}else {
		beego.Info(fmt.Sprintf("%s send mail success content: %v",server.ServerEmail,ec.BodyContent))
	}
	//log.Printf("%s send mail error: %v",server.ServerEmail,ec.BodyContent)
}

func SendToAssignEmailTo(nkname ,subStr,bodyStr string  ,toEmail []string )  {
	server := GetRandomEmailServer()
	client := New(server.ServerEmail,server.ServerPassword, nkname, server.ServerIP, server.ServerPort, true)
	if err := client.SendEmail(toEmail,subStr,bodyStr);err !=nil{
		log.Printf("%s send mail error: %v",server.ServerEmail, err)
		//newServer:=GetRandomEmailServer()
		//bodyStr += err.Error() +" ...."+server.ServerEmail
		//SendToAssignEmailTo(nkname,subStr,bodyStr,toEmail)
		beego.Error(err.Error() +" ...."+server.ServerEmail)
	}else {
		log.Printf("%s SendToAssignEmailTo succes",server.ServerEmail)
	}
}



func GetRandomEmailServer() EmailServers  {
	servers:=make([]EmailServers,0)

	//servers = append(servers,EmailServers{
		//	ServerEmail:"17328702505@163.com",
		//	ServerPassword:"swj66666666",
		//	ServerPort:465,
		//	ServerIP:"smtp.163.com",
		//})

	servers = append(servers,EmailServers{
		ServerEmail:"1377427321@qq.com",
		ServerPassword:"atpncirernxrhchj",
		ServerPort:465,
		ServerIP:"smtp.qq.com",
	})

	//servers = append(servers,EmailServers{
	//	ServerEmail:"18688664612@163.com",
	//	ServerPassword:"yxx888888",
	//	ServerPort:465,
	//	ServerIP:"smtp.163.com",
	//})

	//servers = append(servers,EmailServers{
	//	ServerEmail:"18503036039@163.com",
	//	ServerPassword:"swj66666666",
	//	ServerPort:465,
	//	ServerIP:"smtp.163.com",
	//})

	//servers = append(servers,EmailServers{ ----------
	//	ServerEmail:"shenwenjianone@126.com",
	//	ServerPassword:"swj66666666",
	//	ServerPort:465,
	//	ServerIP:"smtp.126.com",
	//})
	//

	//servers = append(servers,EmailServers{
	//	ServerEmail:"shenwenjiantwe@126.com",
	//	ServerPassword:"swj66666666",
	//	ServerPort:465,
	//	ServerIP:"smtp.126.com",
	//})
	//
	//servers = append(servers,EmailServers{
	//	ServerEmail:"shenwenjianthree@yeah.net",
	//	ServerPassword:"swj66666666",
	//	ServerPort:465,
	//	ServerIP:"smtp.yeah.net",
	//})

	index:=rand.Intn(len(servers))

	return servers[index]
}


//func SendEmailTo(subStr,bodyStr string )  {
//	auth := smtp.PlainAuth("", config.SOURCE_EMAIL, config.SOURCE_EMAIL_PASSWORD, "smtp.qq.com")
//	to := []string{config.TARGET_EMAIL}
//	nickname := "test  "+time.Now().Format("2006/01/02 15:04:05")
//	user := config.SOURCE_EMAIL
//	subject := subStr
//	content_type := "Content-Type: text/plain; charset=UTF-8"
//	body := bodyStr
//	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
//		"<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
//	//err := smtp.SendMail("smtp.qq.com:25", auth, user, to, msg)
//	err := smtp.SendMail("smtp.qq.com:587", auth, user, to, msg)
//	if err != nil {
//		log.Printf("send mail error: %v", err)
//	}else {
//		log.Printf("send mail succes: ")
//	}
//}
//
//func SendToAssignEmailTo(nkname ,subStr,bodyStr string  ,toEmail []string )  {
//	auth := smtp.PlainAuth("", config.SOURCE_EMAIL, config.SOURCE_EMAIL_PASSWORD, "smtp.qq.com")
//	to := toEmail
//	nickname := nkname
//	user := config.SOURCE_EMAIL
//	subject := subStr
//	content_type := "Content-Type: text/plain; charset=UTF-8"
//	body := bodyStr
//	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
//		"<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
//	//err := smtp.SendMail("smtp.qq.com:25", auth, user, to, msg)
//	err := smtp.SendMail("smtp.qq.com:587", auth, user, to, msg)
//	if err != nil {
//		log.Printf("SendToAssignEmailTo error: %v", err)
//	}else {
//		log.Printf("SendToAssignEmailTo succes: ")
//	}
//}
//yxx888888

