package mq

import (
	"errors"
	"os"

	proto "github.com/golang/protobuf/proto"
	"github.com/micro-kit/micro-common/logger"
	"github.com/micro-kit/micro-common/mq/nsqlib"
	"github.com/nsqio/go-nsq"
	"github.com/smile-im/microkit-client/proto/accesspb"
)

/* nsq 实现 */

type Nsq struct {
	BaseMq
	producer  *nsq.Producer
	recvTopic string
	sendTopic string
}

// 初始化订阅mq消息和发送客户端
func InitNsq() Mq {
	if os.Getenv("MQ_TYPE") != MqTypeNsq {
		return nil
	}
	// 生产者
	producer, err := nsqlib.NewProducer()
	if err != nil {
		logger.Logger.Panicw("创建nsq生产者错误", "err", err)
	}
	return &Nsq{
		BaseMq:    baseMq,
		producer:  producer,
		recvTopic: os.Getenv("NSQ_RECV_TOPIC"),
		sendTopic: os.Getenv("NSQ_SEND_TOPIC"),
	}
}

// Recv 接收客户端发来的消息
func (n *Nsq) Recv() (err error) {
	err = nsqlib.NewConsumer(n.recvTopic, n.srvId, n)
	return
}

// 处理一条nsq消息
func (n *Nsq) HandleMessage(message *nsq.Message) (err error) {
	logger.Logger.Debugw("收到nsq消息", "msg", message)
	if message == nil {
		logger.Logger.Warnw("nsq 消息为nil", "srv_id", n.srvId)
		return errors.New("nsq 消息为nil")
	}
	mqMsg := new(accesspb.MqMessage)
	err = proto.Unmarshal(message.Body, mqMsg)
	if err != nil {
		logger.Logger.Warnw("解析nsq消息内容错误", "srv_id", n.srvId, "body", message.Body)
		return
	}
	n.msgChan <- mqMsg
	return
}

// Send 向用户发送消息
func (n *Nsq) Send(msg *accesspb.Message) (err error) {
	mqMessage := &accesspb.MqMessage{
		Msg:   msg,
		SrvId: n.srvId,
	}
	body, err := proto.Marshal(mqMessage)
	if err != nil {
		logger.Logger.Errorw("编码mq消息对象错误", "srv_id", n.srvId, "body", mqMessage)
		return
	}
	err = n.producer.Publish(n.sendTopic, body)
	if err != nil {
		logger.Logger.Errorw("发送mq消息到nsq错误", "err", err, "srv_id", n.srvId, "body", mqMessage)
	}
	return
}
