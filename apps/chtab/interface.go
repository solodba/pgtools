package chtab

import "context"

// 服务模块名称
const (
	AppName = "chtab"
)

// 服务接口
type Service interface {
	// 修改主从信息表
	UpdatePrimaryStandbyTable(context.Context) error
}
