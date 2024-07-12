package protocol

import (
	"context"

	"github.com/solodba/mcube/apps"
	"github.com/solodba/pgtools/apps/chkps"
)

var (
	ctx = context.Background()
)

// chkps服务结构体
type ChkpsService struct {
	chksvc chkps.Service
}

// ChkpsService服务结构体构造函数
func NewChkpsService() *ChkpsService {
	return &ChkpsService{
		chksvc: apps.GetInternalApp(chkps.AppName).(chkps.Service),
	}
}

// chkps服务启动方法
func (m *ChkpsService) Start() error {
	err := m.chksvc.CheckPostgresqlProcess(ctx)
	if err != nil {
		return err
	}
	return nil
}

// chkps服务停止方法
func (s *ChkpsService) Stop() error {
	return nil
}
