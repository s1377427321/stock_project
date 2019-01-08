package storage

import (
	//	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
)

func Test_testOne(t *testing.T) {
	log.Println("tsttesttest")
	a := NewMysql()
	log.Println(a)
}
