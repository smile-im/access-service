module github.com/smile-im/access-service

go 1.15

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.4
	github.com/micro-kit/micro-common => ../../micro-kit/micro-common
	github.com/micro-kit/microkit => ../../micro-kit/microkit
	github.com/smile-im/microkit-client/client/access => ../microkit-client/client/access
	github.com/smile-im/microkit-client/proto/accesspb => ../microkit-client/proto/accesspb
	go.etcd.io/bbolt => github.com/coreos/bbolt v1.3.4
	google.golang.org/grpc => google.golang.org/grpc v1.26.0 // grpc对etcd依赖问题
)

require (
	github.com/micro-kit/micro-common v0.0.0-00010101000000-000000000000
	github.com/micro-kit/microkit v0.0.0-00010101000000-000000000000
	github.com/smile-im/microkit-client/proto/accesspb v0.0.0-00010101000000-000000000000
	go.uber.org/zap v1.16.0
	google.golang.org/grpc v1.28.1
)
