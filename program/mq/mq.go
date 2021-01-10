package mq

import (
	"os"

	"github.com/micro-kit/micro-common/logger"
	"github.com/smile-im/microkit-client/proto/accesspb"
)

/* im mq相关操作 */

const (
	MqTypeNsq      = "nsq"
	MqTypeRabbitmq = "rabbitmq"
)

var (
	MqClient Mq
)

type Mq interface {
	// Recv 接收客户端发来的消息
	Recv() error
	// Send 向用户发送消息
	Send(msg *accesspb.Message) error
	// 返回mq消息chan
	MessageChan() chan *accesspb.MqMessage
}

// mq基础对象
type BaseMq struct {
	srvId   string
	msgChan chan *accesspb.MqMessage
}

var (
	baseMq = BaseMq{
		msgChan: make(chan *accesspb.MqMessage),
		srvId:   os.Getenv("SVC_ID"),
	}
)

// 获取消息管道
func (bm *BaseMq) MessageChan() chan *accesspb.MqMessage {
	return bm.msgChan
}

func init() {
	srvId := os.Getenv("SVC_ID")
	switch os.Getenv("MQ_TYPE") {
	case MqTypeNsq:
		MqClient = InitNsq()
	case MqTypeRabbitmq:

	default:
		logger.Logger.Infow("不支持的mq类型", "srv_id", srvId, "mq_type", os.Getenv("MQ_TYPE"))
	}
	if MqClient != nil {
		go func() {
			err := MqClient.Recv()
			if err != nil {
				logger.Logger.Errorw("订阅mq消息错误", "err", err, "srv_id", srvId, "mq_type", os.Getenv("MQ_TYPE"))
			}
		}()
	}
}
