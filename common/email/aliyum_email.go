package email

import (
	"net/smtp"
	"fmt"
	"bytes"
	"strings"
	"time"
	"mime"
	"log"
	"io/ioutil"
	"encoding/base64"
)

// define email interface, and implemented auth and send method
type Mail interface {
	Auth()
	Send(message Message) error
}
type SendMail struct {
	user     string
	password string
	host     string
	port     string
	auth     smtp.Auth
}
type Attachment struct {
	name        string
	contentType string
	withFile    bool
}
type Message struct {
	from        string
	to          []string
	cc          []string
	bcc         []string
	subject     string
	body        string
	contentType string
	attachment  Attachment
}
func main() {
	user := "XXX@XXXXX.top"
	password := "TESXXXXXX"
	host := "smtpdm.aliyun.com"
	port := "25"
	var mail Mail
	mail = &SendMail{user: user, password: password, host: host, port: port}
	message := Message{from: user,
		to:          []string{"XXXXX@qq.com", "XX@qq.com", "XXX@163.com"},
		cc:          []string{},
		bcc:         []string{},
		subject:     "HELLO WORLD",
		body:        "哈哈哈哈哈哈哈",
		contentType: "text/plain;charset=utf-8",
		// attachment: Attachment{
		//     name:        "test.jpg",
		//     contentType: "image/jpg",
		//     withFile:    true,
		// },
		attachment: Attachment{
			name:        "D:\\goProjects\\src\\测试pdf.pdf",
			contentType: "application/octet-stream",
			withFile:    true,
		},
	}
	err := mail.Send(message)
	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("Send mail success!")
	}
}
func (mail *SendMail) Auth() {
	// mail.auth = smtp.PlainAuth("", mail.user, mail.password, mail.host)
	mail.auth = LoginAuth(mail.user, mail.password)
}
func (mail SendMail) Send(message Message) error {
	mail.Auth()
	buffer := bytes.NewBuffer(nil)
	boundary := "GoBoundary"
	Header := make(map[string]string)
	Header["From"] = message.from
	Header["To"] = strings.Join(message.to, ";")
	Header["Cc"] = strings.Join(message.cc, ";")
	Header["Bcc"] = strings.Join(message.bcc, ";")
	Header["Subject"] = message.subject
	Header["Content-Type"] = "multipart/mixed;boundary=" + boundary
	Header["Mime-Version"] = "1.0"
	Header["Date"] = time.Now().String()
	mail.writeHeader(buffer, Header)
	body := "\r\n--" + boundary + "\r\n"
	body += "Content-Type:" + message.contentType + "\r\n"
	body += "\r\n" + message.body + "\r\n"
	buffer.WriteString(body)
	if message.attachment.withFile {
		attachment := "\r\n--" + boundary + "\r\n"
		attachment += "Content-Transfer-Encoding:base64\r\n"
		attachment += "Content-Disposition:attachment\r\n"
		attachment += "Content-Type:" + message.attachment.contentType + ";name=\"" + mime.BEncoding.Encode("UTF-8", message.attachment.name) + "\"\r\n"
		buffer.WriteString(attachment)
		defer func() {
			if err := recover(); err != nil {
				log.Fatalln(err)
			}
		}()
		mail.writeFile(buffer, message.attachment.name)
	}
	to_address := MergeSlice(message.to, message.cc)
	to_address = MergeSlice(to_address, message.bcc)
	buffer.WriteString("\r\n--" + boundary + "--")
	err := smtp.SendMail(mail.host+":"+mail.port, mail.auth, message.from, to_address, buffer.Bytes())
	return err
}
func MergeSlice(s1 []string, s2 []string) []string {
	slice := make([]string, len(s1)+len(s2))
	copy(slice, s1)
	copy(slice[len(s1):], s2)
	return slice
}
func (mail SendMail) writeHeader(buffer *bytes.Buffer, Header map[string]string) string {
	header := ""
	for key, value := range Header {
		header += key + ":" + value + "\r\n"
	}
	header += "\r\n"
	buffer.WriteString(header)
	return header
}
// read and write the file to buffer
func (mail SendMail) writeFile(buffer *bytes.Buffer, fileName string) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err.Error())
	}
	payload := make([]byte, base64.StdEncoding.EncodedLen(len(file)))
	base64.StdEncoding.Encode(payload, file)
	buffer.WriteString("\r\n")
	for index, line := 0, len(payload); index < line; index++ {
		buffer.WriteByte(payload[index])
		if (index+1)%76 == 0 {
			buffer.WriteString("\r\n")
		}
	}
}
type loginAuth struct {
	username, password string
}
func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}
func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	// return "LOGIN", []byte{}, nil
	return "LOGIN", []byte(a.username), nil
}
func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		}
	}
	return nil, nil
}