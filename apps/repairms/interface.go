package repairms

import "context"

// 服务模块名称
const (
	AppName = "repairms"
)

// 服务接口
type Service interface {
	// 修复主备
	RepairPrimaryStandby(context.Context) error
}
