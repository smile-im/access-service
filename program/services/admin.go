package services

import (
	"context"

	"github.com/smile-im/microkit-client/proto/accesspb"
)

/* 提供给管理后台使用的rpc */

// Admin 实现grpc管理端rpc接口
type Admin struct {
	Base
}

// NewAdmin 创建管理后台rpc对象
func NewAdmin() *Admin {
	return &Admin{
		Base: NewBase(),
	}
}

// CreateRoom 创建房间
func (s *Admin) CreateRoom(ctx context.Context, req *accesspb.CreateRoomRequest) (*accesspb.CreateRoomReply, error) {
	// 验证参数是否错误

	// TODO 逻辑代码

	// 返回结果
	reply := &accesspb.CreateRoomReply{}
	return reply, nil
}
