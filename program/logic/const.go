package logic

/* 本包使用的常量 */

var (
	MessageVer int32 = 202101
)

// redis key
const (
	// 用于记录用户连接所在客户端，参数为用户id，值为map类型对应每种平台的连接所在
	RedisUserConnKey = "user:conn:%d"
)

// 平台
const (
	PlatformH5  = "h5"  // h5网页
	PlatformApp = "app" // 手机app
	PlatformPC  = "pc"  // 电脑
)
