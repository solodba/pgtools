package impl

import (
	"context"

	"github.com/solodba/pgtools/apps/recover"
)

// 获取当前lsn号和当前lsn号所在的wal文件
func (i *impl) GetLsnAndFile(ctx context.Context) error {
	return nil
}

// 切换lsn日志
func (i *impl) SwitchWalLog(context.Context) error {
	return nil
}

// 获取误删事务号
func (i *impl) GetDeleteTxid(ctx context.Context, fileInfo *recover.FileInfo) (*recover.TxInfo, error) {
	return nil, nil
}

// 关闭数据库
func (i *impl) StopDb(ctx context.Context) error {
	return nil
}

// 恢复数据库到指定事务号
func (i *impl) RecoverDbToTxid(ctx context.Context, txInfo *recover.TxInfo) error {
	return nil
}

// 启动数据库
func (i *impl) StartDb(ctx context.Context) error {
	return nil
}
