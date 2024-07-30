package rectl

import "context"

// 服务模块名称
const (
	AppName = "rectl"
)

// 服务接口
type Service interface {
	// 获取最新wal日志名
	GetNextWal(context.Context) (string, error)
	// 获取下一个多事务ID和最旧多事务ID
	GetMxid(context.Context) (*Mxid, error)
	// 获取下一个多事务处理偏移量
	GetNextMxidOffset(context.Context) (string, error)
	// 获取下一个事务ID
	GetNextXid(context.Context) (string, error)
	// 生成重建控制文件语句
	GenRebuildControlFileCmd(context.Context) (string, error)
}
