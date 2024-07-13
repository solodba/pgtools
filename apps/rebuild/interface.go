package rebuild

import "context"

// 服务模块名称
const (
	AppName = "rebuild"
)

// 服务接口
type Service interface {
	// 重构备库
	RebuildStandby(context.Context) error
}
