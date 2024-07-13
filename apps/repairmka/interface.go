package repairmka

import "context"

// 服务模块名称
const (
	AppName = "repairmka"
)

// 服务接口
type Service interface {
	// 修复主节点keepalived
	RepairPriamryKeepalived(context.Context) error
}
