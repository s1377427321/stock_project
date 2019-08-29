package commoncache

import (
	"github.com/spf13/viper"
	"gopkg.in/redis.v5"
	"github.com/astaxie/beego/logs"
	"fmt"
	"time"
	"strconv"
)

var client *redis.Client

const (
	//用于保存话题发布时间
	IDPUBLISHTIME string = "id:publish:time"
)

func CommonCacheInit() {
	addr := viper.GetString("redis.first.addr")
	password := viper.GetString("redis.first.passwd")
	index := viper.GetInt("redis.first.index")
	client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       index,
	})
	_, err := client.Ping().Result()
	if err != nil {
		logs.Error("CommonCacheInit client.Ping().Result()", err)
		return
	}
	logs.Info("CommonCacheInit success")
}


func SaveIDPublishTime(id int64, val int64) {
	key := fmt.Sprintf("%s:%d", IDPUBLISHTIME, id)
	value := fmt.Sprintf("%d", val)
	dismissTime := viper.GetInt("redpackage.DefaultSaveDay")
	client.Set(key, value, time.Second*time.Duration(dismissTime))
}

func GetIDPublishTime(id int64) int64 {
	//key := fmt.Sprintf("%s:%d", IDPUBLISHTIME, id)
	//valstr :=client.Get(key).Val()
	//val, err := strconv.ParseInt(valstr, 10, 64)
	//if err != nil {
	//	logs.Error(nil, "GetIDPublishTime.strconv.ParseInt error  %v  %v  %v", id, valstr, err)
	//	return 0
	//}
	//
	//return val

	key := fmt.Sprintf("%s:%d", IDPUBLISHTIME, id)
	valstr ,err:=client.Get(key).Result()

	if err != nil {
		logs.Error(nil, "GetIDPublishTime.strconv.ParseInt error  %v  %v  %v", id, valstr, err)
		return 0
	}
	val, err := strconv.ParseInt(valstr, 10, 64)

	return val
}
