package common

import (
	"path/filepath"
	"os"
	"github.com/astaxie/beego/logs"
	"strings"
	"fmt"
	"errors"
	"github.com/astaxie/beego"
	"reflect"
	"time"
	"strconv"
)

type IntSlice []int

func (c IntSlice) Len() int {
	return len(c)
}
func (c IntSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c IntSlice) Less(i, j int) bool {
	return  c[i] < c[j]
}


func CreateBeegoLog(fileName string) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		logs.Error(err)
	}
	dir = strings.Replace(dir, "\\", "/", -1)

	filePath := dir

	SetLogOption(filePath, fileName, 6, true)
}

//设置日志
func SetLogOption(logDir, namePrefix string, logLevel int, logToConsole bool) error {
	fmt.Println("logDir ", logDir)
	fmt.Println("namePrefix ", namePrefix)
	if fi, err := os.Stat(logDir); os.IsNotExist(err) {
		err = os.Mkdir(logDir, 0666)
		if err != nil {
			panic(err)
		}

		fmt.Println("The log dir:", logDir, "doesn't exist, create it!")
	} else {
		if !fi.IsDir() {
			panic(errors.New(fmt.Sprintf("The file:", logDir, "is not a directory!")))
		}
	}

	if logLevel < beego.LevelEmergency || logLevel > beego.LevelDebug {
		panic(errors.New(fmt.Sprintf("Invalid logLevel:", logLevel)))
	}

	beego.BeeLogger.SetLevel(logLevel)

	str := fmt.Sprintf(`{"filename":"%s/%s.log", "maxlines":1000000,"maxsize":2000000000,"daily":true,"maxdays":7}`, logDir, namePrefix)
	err := beego.BeeLogger.SetLogger("file", str)
	if err != nil {
		panic(err)
	}

	if logToConsole {
		err = beego.BeeLogger.SetLogger("console", str)
		if err != nil {
			panic(err)
		}
	} else {
		beego.BeeLogger.DelLogger("console")
	}

	beego.SetLogFuncCall(true)

	return nil
}




//用map填充结构
func FillStruct(data map[string]interface{}, obj interface{}) error {
	for k, v := range data {
		err := SetField(obj, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

//用map的值替换结构的值
func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()        //结构体属性值
	structFieldValue := structValue.FieldByName(name) //结构体单个属性值

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type() //结构体的类型
	val := reflect.ValueOf(value)              //map值的反射值

	var err error
	if structFieldType != val.Type() {
		val, err = TypeConversion(fmt.Sprintf("%v", value), structFieldValue.Type().Name()) //类型转换
		if err != nil {
			return err
		}
	}

	structFieldValue.Set(val)
	return nil
}

//类型转换
func TypeConversion(value string, ntype string) (reflect.Value, error) {
	if ntype == "string" {
		return reflect.ValueOf(value), nil
	} else if ntype == "time.Time" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "Time" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "int" {
		i, err := strconv.Atoi(value)
		return reflect.ValueOf(i), err
	} else if ntype == "int8" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int8(i)), err
	} else if ntype == "int32" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int64(i)), err
	} else if ntype == "int64" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(i), err
	} else if ntype == "float32" {
		i, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(float32(i)), err
	} else if ntype == "float64" {
		i, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(i), err
	}

	//else if .......增加其他一些类型的转换

	return reflect.ValueOf(value), errors.New("未知的类型：" + ntype)
}

