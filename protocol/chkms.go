package protocol

import (
	"github.com/solodba/mcube/apps"
	"github.com/solodba/pgtools/apps/chkms"
)

// chkms服务结构体
type ChkmsService struct {
	chkmssvc chkms.Service
}

// ChkmsService服务结构体构造函数
func NewChkmsService() *ChkmsService {
	return &ChkmsService{
		chkmssvc: apps.GetInternalApp(chkms.AppName).(chkms.Service),
	}
}

// chkms服务启动方法
func (m *ChkmsService) Start() error {
	err := m.chkmssvc.GetPostgresqlPrimaryStandby(ctx)
	if err != nil {
		return err
	}
	return nil
}

// chkps服务停止方法
func (s *ChkmsService) Stop() error {
	return nil
}
