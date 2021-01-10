# AccessService
用于长连接接入，前期支持grpc流模式，后期支持websocket


## TODO
- 是否登录和上传文件也放这里
- 支持系统消息，向某人或某房间某人或某房间广播消息
- 房间消息，每个服务都订阅，使用本地map存储房间连接

 micro-cli project addrpc -svc=access -rpc=Connect -comment=获取 -type=foreground -croot=github.com/smile-im/microkit-client -root=github.com/smile-im
 
 micro-cli project addrpc -svc=access -rpc=CreateRoom -comment=创建房间 -type=admin -croot=github.com/smile-im/microkit-client -root=github.com/smile-im