package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/smile-im/microkit-client/client/access"
	"github.com/smile-im/microkit-client/proto/accesspb"
)

var (
	accessClient  accesspb.AccessClient
	connectClient accesspb.Access_ConnectClient
)

func init() {
	var err error
	accessClient, err = access.NewClient()
	log.Println(2)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	connectClient, err = accessClient.Connect(context.Background())
	log.Println(3)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println("连接成功")
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	go Recv()

	// 发送登录消息
	err := connectClient.Send(&accesspb.ConnectRequest{
		Msg: &accesspb.Message{
			Operation: accesspb.OperationType_OperationAuth,
			Body:      nil,
		},
	})
	if err != nil {
		log.Println("发送登录消息错误", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		err := connectClient.Send(&accesspb.ConnectRequest{
			Msg: &accesspb.Message{
				From:      1,
				To:        1,
				Operation: accesspb.OperationType_OperationSendMsg,
				Body:      scanner.Bytes(),
			},
		})
		if err != nil {
			log.Println("发送消息错误", err)
			os.Exit(1)
		}
	}
}

// 接收消息
func Recv() {

	for {
		msg, err := connectClient.Recv()
		if err != nil {
			log.Println("读取消息错误", err)
			return
		}
		log.Println("接收到消息", string(msg.Msg.Body))
	}
}
