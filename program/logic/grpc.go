package logic

import (
	"errors"
	"io"
	"sync"

	proto "github.com/golang/protobuf/proto"
	"github.com/micro-kit/micro-common/logger"
	"github.com/smile-im/microkit-client/proto/accesspb"
)

/* grpc 流模式 */

type GrpcLogic struct {
	BaseLogic
	stream accesspb.Access_ConnectServer
}

var (
	grpcStreamMap *sync.Map
)

func init() {
	grpcStreamMap = new(sync.Map)
	RegImLogic(grpcStreamMap)
}

// RegGrpcLogic 注册连接
func RegGrpcLogic(stream accesspb.Access_ConnectServer) error {
	if stream == nil {
		logger.Logger.Infow("注册grpc连接，流对象为nil")
		return errors.New("注册grpc连接，流对象为nil")
	}
	in, err := stream.Recv()
	if err != nil {
		logger.Logger.Errorw("读取登录消息错误", "err", err)
		return err
	}
	var grpcLogic ImLogic
	grpcLogic = &GrpcLogic{
		stream:    stream,
		BaseLogic: baseLogic,
	}
	// 校验登录 - 发生登录响应
	userInfo, err := grpcLogic.Login(in.GetMsg())
	authResp := &accesspb.Message{
		Ver:       MessageVer,
		Operation: accesspb.OperationType_OperationAuthResp,
		Seq:       in.Msg.Seq,
	}
	authBody := &accesspb.AuthBody{}
	if err != nil {
		authBody.State = accesspb.AuthBody_AuthFail
		authBody.Message = err.Error()
		authBodyBytes, _ := proto.Marshal(authBody)
		authResp.Body = authBodyBytes
		_err := grpcLogic.Send(authResp)
		if _err != nil {
			logger.Logger.Errorw("fail发送登录回复错误", "err", _err)
		}
		return err
	}
	authResp.To = userInfo.UserId
	authBody.State = accesspb.AuthBody_AuthSuccess
	authBody.Message = "success"
	authBodyBytes, _ := proto.Marshal(authBody)
	authResp.Body = authBodyBytes
	err = grpcLogic.Send(authResp)
	if err != nil {
		logger.Logger.Errorw("fail发送登录回复错误", "err", err)
		return err
	}
	// 存储连接信息
	grpcStreamMap.Store(userInfo.UserId, grpcLogic)
	defer func() {
		grpcStreamMap.Delete(userInfo.UserId)
	}()
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			logger.Logger.Debugw("读取到流结束", "err", err)
			return nil
		}
		if err != nil {
			logger.Logger.Errorw("读取新流错误", "err", err)
			return nil
		}
		err = grpcLogic.Recv(in.GetMsg())
		if err != nil {
			logger.Logger.Warnw("处理接收到消息错误", "err", err, "msg", in.GetMsg())
			return err
		}
	}
}

// Send 向用户发送消息
func (gl *GrpcLogic) Send(msg *accesspb.Message) (err error) {
	err = gl.stream.Send(&accesspb.ConnectReply{
		Msg: msg,
	})
	return
}
