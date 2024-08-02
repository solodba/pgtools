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
	// 获取当前数据库聚簇信息
	GetPgClusterInfo(context.Context) (*PgClusterInfo, error)
	// 获取当前IO消耗TOP 10的SQL
	GetComsumeIoSql(context.Context) (*ComsumeTopSqlSet, error)
	// 生成AWR数据
	GenAwrData(context.Context) (*AwrData, error)
	// 生成AWR报告
	GenAwrReport(context.Context) error
}
