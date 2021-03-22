/**
提供发布订阅功能，客户端订阅服务端
采用golang-nsq方式进行处理
*/

package subpub

import (
	"encoding/json"
	"github.com/nsqio/go-nsq"

	"time"
)

var (
	nsqServerOpt NsqServerOptions
	nsqProducer  *nsq.Producer
	nsqComsumer  *nsq.Consumer
)

type NsqServerOptions struct {
	IP   string
	Port string
}

type NsqMessage struct {
	ID        int
	UUID      string
	Name      string
	Content   interface{} //消息内容
	PayLoad   []byte      // 消息负载
	TimeStamp int64
}

func NsqSeverInit() {
	nsqServerOpt = NsqServerOptions{
		IP:   "",
		Port: "",
	}
}
func NsqProducerSetUp(options NsqServerOptions) (*nsq.Producer, error) {
	var (
		nsqProducer *nsq.Producer
		err         error
	)
	config := nsq.NewConfig()
	nsqProducer, err = nsq.NewProducer(options.IP+":"+options.Port, config)
	if err != nil {
		return nil, err
	}

	/*	err := w.Publish("write_test", []byte("test"))
		if err != nil {
			return nsqProducer err
		}

		w.Stop()*/
	return nsqProducer, nil
}

func NsqConsumerSetup(options NsqServerOptions, topic string, channel string, hander nsq.Handler) (*nsq.Consumer, error) {
	var (
		nsqConsumer *nsq.Consumer
		err         error
	)
	config := nsq.NewConfig()
	//todo  根据 需求设置config配置
	config.MaxAttempts = 10
	// Maximum number of messages to allow in flight
	config.MaxInFlight = 5
	// Maximum duration when REQueueing
	config.MaxRequeueDelay = time.Second * 900
	config.DefaultRequeueDelay = time.Second * 0

	nsqConsumer, err = nsq.NewConsumer(topic, channel, config)
	if err != nil {
		return nil, err
	}
	nsqComsumer.AddHandler(hander)
	/*	nsqComsumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		//log.Printf("Got a message: %v", message)

		//wg.Done()
		return nil
	}))*/
	//todo 多服务端server的情况下采用ConnectToNSQLookupds()方法进行实例寻找
	//err = nsqComsumer.ConnectToNSQLookupds(options.IP+":"+options.Port)
	err = nsqComsumer.ConnectToNSQLookupd(options.IP + ":" + options.Port)
	if err != nil {
		return nil, err
	}

	return nsqConsumer, nil
}

func NsqProduerPublish(producer *nsq.Producer, topic string, payLoadData []byte) error {

	return producer.Publish(topic, payLoadData)

}

func NsqProducerStop(producer *nsq.Producer) {

	producer.Stop()
	//return nil
}

func NsqConsumerStop(consumer *nsq.Consumer) {

	consumer.Stop()
}

// demo of hander message
func (Msg *NsqMessage) HandleMessage(message *nsq.Message) error {
	var (
		err error
	)
	//todo
	//Process the Message
	var nsqMsg NsqMessage
	if err := json.Unmarshal(message.Body, &nsqMsg); err != nil {

		// Returning a non-nil error will automatically send a REQ command to NSQ to re-queue the message.
		return err
	}
	return err
}
