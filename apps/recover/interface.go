package recover

import "context"

// 服务模块名称
const (
	AppName = "recover"
)

// 服务接口
type Service interface {
	// 获取当前lsn号和当前lsn号所在的wal文件
	GetLsnAndFile(context.Context) (*FileInfo, error)
	// 切换lsn日志
	SwitchWalLog(context.Context) error
	// 获取误删事务号
	GetDeleteTxid(context.Context, *FileInfo) (*TxInfo, error)
	// 关闭数据库
	StopDb(context.Context) error
	// 恢复数据库到指定事务号
	RecoverDbToTxid(context.Context, *TxInfo) error
	// 启动数据库
	StartDb(context.Context) error
}

// FileInfo结构体
type FileInfo struct {
	// 文件路径
	FilePath string `json:"file_path"`
	// 当前lsn号
	LsnNum string `json:"lsn_num"`
}

// FileInfo结构体初始化函数
func NewFileInfo() *FileInfo {
	return &FileInfo{}
}
