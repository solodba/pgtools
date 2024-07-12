package chkps

import "context"

// 服务模块名称
const (
	AppName = "chkps"
)

// 服务接口
type Service interface {
	// 检查服务器上postgresql进程
	CheckPostgresqlProcess(context.Context) error
}
