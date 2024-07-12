package chkms

import "context"

// 服务模块名称
const (
	AppName = "chkms"
)

// 服务接口
type Service interface {
	// 获取主从关系
	GetPostgresqlPrimaryStandby(context.Context) error
}
