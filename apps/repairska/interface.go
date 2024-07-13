package repairska

import "context"

// 服务模块名称
const (
	AppName = "repairska"
)

// 服务接口
type Service interface {
	// 修复节点keepalived
	RepairStandbyKeepalived(context.Context) error
}
