package logic

import (
	"errors"
	"fmt"
	"sync"

	"github.com/micro-kit/micro-common/logger"
	"github.com/smile-im/access-service/program/mq"
	"github.com/smile-im/microkit-client/proto/accesspb"
	"github.com/smile-im/microkit-client/proto/authpb"
)

/* 统一处理 */

var (
	AllImLogic = make([]*sync.Map, 0)
)

// RegImLogic 注册多种接入对象
func RegImLogic(imLogic *sync.Map) {
	AllImLogic = append(AllImLogic, imLogic)
}

type ImLogic interface {
	// Login 登录
	Login(msg *accesspb.Message) (*authpb.UserInfo, error)
	// Recv 接收客户端发来的消息
	Recv(msg *accesspb.Message) error
	// Send 向用户发送消息
	Send(msg *accesspb.Message) error
}

// BaseLogic 基础公共操作
type BaseLogic struct {
	mq mq.Mq
}

var (
	baseLogic BaseLogic
)

func init() {
	baseLogic.mq = mq.MqClient
	go func() {
		for {
			msg := <-mq.MqClient.MessageChan()
			// logger.Logger.Debugw("收到mq消息", "msg", msg)
			for _, v := range AllImLogic {
				// logger.Logger.Debugw("存在连接?", "msg", msg)
				// 找到用户对应连接，如果存在则发送消息
				if val, ok := v.Load(msg.Msg.To); ok {
					// logger.Logger.Debugw("存在连接", "msg", msg, "val", val)
					if imLogic, ok := val.(ImLogic); ok {
						err := imLogic.Send(msg.Msg)
						if err != nil {
							logger.Logger.Errorw("发送mq订阅到的消息到客户端错误", "err", err, "msg", msg)
						} else {
							logger.Logger.Infow("发送mq订阅到的消息到客户端成功", "err", err, "msg", msg)
						}
					}
				}
			}
		}
	}()
}

var XXX int64

// 登录，将连接第一条消息调用此函数
func (bl *BaseLogic) Login(msg *accesspb.Message) (userInfo *authpb.UserInfo, err error) {
	if msg == nil {
		logger.Logger.Infow("第一条消息为nil", "msg", msg)
		return nil, errors.New("第一条消息为nil")
	}
	if msg.Operation != accesspb.OperationType_OperationAuth {
		logger.Logger.Infow("第一条消息不是登录消息", "msg", msg)
		return nil, errors.New("第一条消息不是登录消息")
	}
	// 通过rpc验证登录信息中token是否有效
	XXX++
	userInfo = &authpb.UserInfo{
		UserId:   XXX,
		Nickname: "哈哈" + fmt.Sprint(XXX),
	}
	// TODO 发布用户登录mq消息，让其它连接断开，消息中包含环境变量 SVC_ID
	// TODO 登录成功更新redis连接数

	return
}

// Recv 接收客户端发来的消息
func (bl *BaseLogic) Recv(msg *accesspb.Message) (err error) {
	// 发送至mq
	err = bl.mq.Send(msg)
	if err != nil {
		logger.Logger.Errorw("接收客户端发来的消息发送到mq错误", "err", err, "msg", msg)
	}
	return
}
