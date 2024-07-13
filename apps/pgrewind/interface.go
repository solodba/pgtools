package pgrewind

import "context"

// 服务模块名称
const (
	AppName = "pgrewind"
)

// 服务接口
type Service interface {
	// 修改主从信息表
	PgRewindSyncData(context.Context) error
}
