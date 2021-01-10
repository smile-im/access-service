package services

import (
	"log"

	"github.com/micro-kit/micro-common/logger"
	"github.com/smile-im/access-service/program/logic"
	"github.com/smile-im/microkit-client/proto/accesspb"
)

/* 提供给客户端使用的rpc */

// Foreground 实现grpc客户端rpc接口
type Foreground struct {
	Base
}

// NewForeground 创建客户端rpc对象
func NewForeground() *Foreground {
	return &Foreground{
		Base: NewBase(),
	}
}

// Connect 连接实时消息
func (s *Foreground) Connect(stream accesspb.Access_ConnectServer) error {
	log.Println("有连接")
	defer func() {
		logger.Logger.Debugw("连接断开")
	}()
	return logic.RegGrpcLogic(stream)
	// for {
	// 	in, err := stream.Recv()
	// 	if err == io.EOF {
	// 		logger.Logger.Debugw("读取到流结束", "err", err)
	// 		return nil
	// 	}
	// 	if err != nil {
	// 		logger.Logger.Errorw("读取新流错误", "err", err)
	// 		return nil
	// 	}
	// 	log.Println(in.Msg)
	// 	out := strings.Replace(string(in.Msg.Body), "我", "你", -1)
	// 	out = strings.Replace(out, "？", "！", -1)
	// 	stream.Send(&accesspb.ConnectReply{
	// 		Msg: &accesspb.Message{
	// 			Body: []byte(out),
	// 		},
	// 	})
	// }
	// return nil
}
