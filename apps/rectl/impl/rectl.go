package impl

import (
	"context"

	"github.com/solodba/pgtools/apps/rectl"
)

// 获取最新wal日志名
func (i *impl) GetNextWal(ctx context.Context) (string, error) {
	return "", nil
}

// 获取下一个多事务ID和最旧多事务ID
func (i *impl) GetMxid(ctx context.Context) (*rectl.Mxid, error) {
	return nil, nil
}

// 获取下一个多事务处理偏移量
func (i *impl) GetNextMxidOffset(ctx context.Context) (string, error) {
	return "", nil
}

// 获取下一个事务ID
func (i *impl) GetNextXid(ctx context.Context) (string, error) {
	return "", nil
}
