package kafkawork

import (
	"strconv"
	"time"

	"github.com/lexkong/log/lager/kafkawork/services/configs"
	"github.com/lexkong/log/lager/kafkawork/services/producer"
)

var (
	apilogCfg       *configs.KfkConfig
	apilog_producer *producer.TopicProducer
)

func InitProducer() {
	apilogCfg = &configs.KfkConfig{}
	configs.LoadJsonConfig(apilogCfg, "./log.json")

	apilog_producer = producer.CreateProducer(apilogCfg.Servers)
}

func Init() {
	InitProducer()
}

func PublishLog(logmsg string) error {
	if apilog_producer == nil {
		InitProducer()
	}
	key := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	return apilog_producer.Produce(apilogCfg.Topics[0], key, logmsg)
}
