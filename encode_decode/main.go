package main

import (
	"path/filepath"
	"os"
	"fmt"
	"strings"
	"io/ioutil"

	"bufio"
)

var reader  *bufio.Reader

var suffix =[]string{"rm","rmvb","mpeg","mov","mtv","dat","wmv","avi","3gp","amv","dmv","flv","3GP","mts","mp4"}

func main() {
	reader = bufio.NewReader(os.Stdin)
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Errorf("errror",err)
		return
	}
	dir = strings.Replace(dir, "\\", "/", -1)
	for {
		fmt.Println("1加密   2删除多余文件")
		if strBytes, _, err := reader.ReadLine();err == nil{
			if strings.Contains(string(strBytes),"1") == true{
				fmt.Println("请输入要加密文件或文件夹")
				if strBytes, _, err = reader.ReadLine();err == nil{
					filePath := dir +"/"+ string(strBytes)
					GetDirOrFileSize(filePath)
				}
			}else if strings.Contains(string(strBytes),"2") == true {
				fmt.Println("请输入要删除文件或文件夹")
				if strBytes, _, err = reader.ReadLine();err == nil{
					filePath := dir +"/"+ string(strBytes)
					DeleteNotMovie(filePath)
				}
			}

			fmt.Println("All OVER")

		}else {
			return
		}
	}
}

func GetDirOrFileSize(pathOrfile string )  {
	flist, e := ioutil.ReadDir(pathOrfile)
	if e != nil {
		f, e := os.Stat(pathOrfile)
		if e == nil && f.IsDir() == false {
			EnCodeDeCode(pathOrfile)
		}
		return
	}

	for _, f := range flist {
		if f.IsDir() {
			GetDirOrFileSize(pathOrfile+"/"+f.Name())
		}else {
			EnCodeDeCode(pathOrfile+"/"+f.Name())
		}
	}
}


func EnCodeDeCode(readFilePath string)  {
	fmt.Println("____ start ",readFilePath)
	cout:=0
	ro, err := os.OpenFile(readFilePath,0x00002,0666)
	if err != nil {
		fmt.Println("can't opened this file")
		return
	}
	var readIndex int64=0
	defer ro.Close()
	s := make([]byte, 1024)
	for {
		switch nr, err := ro.Read(s[:]); true {
		case nr < 0:
			fmt.Println(os.Stderr, "cat: error reading: %s\n", err.Error(),readFilePath)
			return
		case nr == 0: // EOF
			fmt.Println( "read success "+readFilePath)
			return
		case nr > 0:
			if cout%2 ==0 {
				gg:=NagationBytes(s)
				_,err=ro.WriteAt(gg[0:nr],readIndex)
				if err !=nil{
					fmt.Println("____________",err,readFilePath)
					return
				}

			}
			readIndex+=int64(nr)
			cout = cout+1
		}
	}
}

func DeleteNotMovie(pathOrfile string)  {
	flist, e := ioutil.ReadDir(pathOrfile)
	if e != nil {
		f, e := os.Stat(pathOrfile)
		if e == nil && f.IsDir() == false {
			EnCodeDeCode(pathOrfile)
			DeleteFile(pathOrfile)
		}
		return
	}

	for _, f := range flist {
		if f.IsDir() {
			DeleteNotMovie(pathOrfile+"/"+f.Name())
		}else {
			//EnCodeDeCode(pathOrfile+"/"+f.Name())
			DeleteFile(pathOrfile+"/"+f.Name())
		}
	}
}

func DeleteFile(file string)  {
	var isFind  = false
	for _,v:=range suffix{
		big:=strings.ToUpper(v)
		if strings.Contains(file,big) {
			isFind = true
			break
		}

		small:=strings.ToLower(v)
		if strings.Contains(file,small) {
			isFind = true
			break
		}

	}

	if isFind == false {
		os.Remove(file)
	}
}
