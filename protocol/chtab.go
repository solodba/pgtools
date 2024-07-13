package protocol

import (
	"github.com/solodba/mcube/apps"
	"github.com/solodba/pgtools/apps/chtab"
)

// chtab服务结构体
type ChtabService struct {
	chtabsvc chtab.Service
}

// ChtabService服务结构体构造函数
func NewChtabService() *ChtabService {
	return &ChtabService{
		chtabsvc: apps.GetInternalApp(chtab.AppName).(chtab.Service),
	}
}

// chtab服务启动方法
func (m *ChtabService) Start() error {
	err := m.chtabsvc.UpdatePrimaryStandbyTable(ctx)
	if err != nil {
		return err
	}
	return nil
}

// chtab服务停止方法
func (s *ChtabService) Stop() error {
	return nil
}
