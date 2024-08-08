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
	// 获取当前数据库聚簇WAL Files信息
	GetPgWalFileInfo(context.Context) (*WalFileInfo, error)
	// 获取当前数据库Bg Writer信息
	GetPgBgWriter(context.Context) (*BgWriterInfo, error)
	// 获取当前所有锁信息
	GetPgLockInfo(context.Context) (*LockInfoSet, error)
	// 获取当前VACUUM信息
	GetPgVacuumInfo(context.Context) (*VacuumInfoSet, error)
	// 获取当前角色信息
	GetPgRoleInfo(context.Context) (*RoleInfoSet, error)
	// 获取当前后端会话信息
	GetPgBackendInfo(context.Context) (*BackendInfo, error)
	// 获取当前表空间信息
	GetPgTablespaceInfo(context.Context) (*TablespaceInfoSet, error)
	// 获取当前所有数据库信息
	GetPgDbInfo(context.Context) (*DbInfoSet, error)
	// 获取当前IO消耗TOP 10的SQL
	GetComsumeIoSql(context.Context, *QueryTopSqlArgs) (*ComsumeTopSqlSet, error)
	// 获取当前耗时TOP 10的SQL
	GetComsumeTimeSql(context.Context, *QueryTopSqlArgs) (*ComsumeTopSqlSet, error)
	// 获取当前消耗Buffer TOP 10的SQL
	GetComsumeBufferSql(context.Context, *QueryTopSqlArgs) (*ComsumeTopSqlSet, error)
	// 获取当前消耗temp空间的SQL
	GetComsumeTempSql(context.Context, *QueryTopSqlArgs) (*ComsumeTopSqlSet, error)
	// 获取当前各种消耗Top 10的SQL
	GetComsumeTopSql(context.Context) (*ComsumeTopSqlTotalSet, error)
	// 获取所有TOP SQL集合
	GetComsumeAllSql(context.Context) (*ComsumeAllSqlSet, error)
	// 生成AWR数据
	GenAwrData(context.Context) (*AwrData, error)
	// 生成AWR报告
	GenAwrReport(context.Context) error
}

// 查找Top Sql条件结构体
type QueryTopSqlArgs struct {
	DbName string `json:"db_name"`
}

// 查找Top Sql条件结构体初始化函数
func NewQueryTopSqlArgs() *QueryTopSqlArgs {
	return &QueryTopSqlArgs{}
}
