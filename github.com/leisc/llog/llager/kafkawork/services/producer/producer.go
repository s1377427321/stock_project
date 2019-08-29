package producer

import (
	"fmt"

	"github.com/Shopify/sarama"
)

type TopicProducer struct {
	producer sarama.AsyncProducer
}

func CreateProducer(servers []string) *TopicProducer {
	fmt.Print("init kafka producer, it may take a few seconds to init the connection\n")

	var err error

	topicproducer := &TopicProducer{}

	mqConfig := sarama.NewConfig()
	mqConfig.Producer.Return.Successes = true
	mqConfig.Producer.RequiredAcks = sarama.WaitForAll
	mqConfig.Version = sarama.V0_10_0_0

	topicproducer.producer, err = sarama.NewAsyncProducer(servers, mqConfig)
	if err != nil {
		msg := fmt.Sprintf("Kafak producer create fail. err: %v", err)
		fmt.Println(msg)
		panic(msg)
	}
	return topicproducer
}

func (p *TopicProducer) Produce(topic string, key string, content string) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(content),
	}

	/*_, _, err := producer.SendMessage(msg)
	if err != nil {
		msg := fmt.Sprintf("Send Error topic: %v. key: %v. content: %v", topic, key, content)
		fmt.Println(msg)
		return err
	}
	fmt.Printf("Send OK topic:%s key:%s value:%s\n", topic, key, content)
	*/
	p.producer.Input() <- msg

	select {
	case suc := <-p.producer.Successes():
		fmt.Printf("offset: %d,  timestamp: %s\n", suc.Offset, suc.Timestamp.String())
	case fail := <-p.producer.Errors():
		fmt.Printf("err: %s\n", fail.Err.Error())
	}
	return nil
}
