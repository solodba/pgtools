package chklog

import "context"

// 服务模块名称
const (
	AppName = "chklog"
)

// 服务接口
type Service interface {
	// 查看切换日志
	CheckPromoteLog(context.Context) error
}
