package main

import (
	"os"
	"fmt"
	"encoding/base64"
	)

var IsDecode = false


func ReadBigFile(fileName string, handle func([]byte,bool)) error {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("can't opened this file")
		return err
	}
	defer f.Close()
	s := make([]byte, 4096)
	for {
		switch nr, err := f.Read(s[:]); true {
		case nr < 0:
			fmt.Fprintf(os.Stderr, "cat: error reading: %s\n", err.Error())
			os.Exit(1)
		case nr == 0: // EOF
			handle(nil,true)
		case nr > 0:
			handle(s[0:nr],false)
		}
	}
	return nil
}

func WriteBigFile(fileName string, handle func([]byte,bool)) error {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("can't opened this file")
		return err
	}
	defer f.Close()
	s := make([]byte, 4096)
	for {
		switch nr, err := f.Read(s[:]); true {
		case nr < 0:
			fmt.Fprintf(os.Stderr, "cat: error reading: %s\n", err.Error())
			os.Exit(1)
		case nr == 0: // EOF
			handle(nil,true)
		case nr > 0:
			handle(s[0:nr],false)
		}
	}
	return nil
}

func Do() {
	var aeskey = []byte("321423u9y8d2fwfl")
	pass := []byte("vdncloud123456")
	xpass, err := AesEncrypt(pass, aeskey)
	if err != nil {
		fmt.Println(err)
		return
	}

	pass64 := base64.StdEncoding.EncodeToString(xpass)
	fmt.Printf("加密后:%v\n", pass64)

	bytesPass, err := base64.StdEncoding.DecodeString(pass64)
	if err != nil {
		fmt.Println(err)
		return
	}

	tpass, err := AesDecrypt(bytesPass, aeskey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("解密后:%s\n", tpass)
}

func NagationBytes(a []byte) []byte {
	newB:=make([]byte,len(a))
	for i:=0;i<len(a);i++ {
		newB[i]= nagationByte(a[i])
	}

	return  newB// Go语言取反方式和C语言不同，Go语言不支持~符号。
}


// 变换符号
func nagationByte(a byte) byte {
	// 注意: C语言中是 ~a+1这种方式
	return ^a + 1 // Go语言取反方式和C语言不同，Go语言不支持~符号。
}

// 变换符号
func nagation(a int) int {
	// 注意: C语言中是 ~a+1这种方式
	return ^a + 1 // Go语言取反方式和C语言不同，Go语言不支持~符号。
}