package services

import (
	"context"

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
func (s *Foreground) Connect(ctx context.Context, req *accesspb.ConnectRequest) (*accesspb.ConnectReply, error) {
	// 验证参数是否错误

	// TODO 逻辑代码

	// 返回结果
	reply := &accesspb.ConnectReply{}
	return reply, nil
}
