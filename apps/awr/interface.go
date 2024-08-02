package awr

import "context"

// 服务模块名称
const (
	AppName = "awr"
)

// 服务接口
type Service interface {
	// 获取当前系统信息
	GetSystemInfo(context.Context) (*SystemInfo, error)
}
