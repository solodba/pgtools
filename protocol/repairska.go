package protocol

import (
	"github.com/solodba/mcube/apps"
	"github.com/solodba/pgtools/apps/repairska"
)

// RepairskaService服务结构体
type RepairskaService struct {
	repairskasvc repairska.Service
}

// RepairskaService服务结构体构造函数
func NewRepairskaService() *RepairskaService {
	return &RepairskaService{
		repairskasvc: apps.GetInternalApp(repairska.AppName).(repairska.Service),
	}
}

// RepairskaService服务启动方法
func (m *RepairskaService) Start() error {
	err := m.repairskasvc.RepairStandbyKeepalived(ctx)
	if err != nil {
		return err
	}
	return nil
}

// chtab服务停止方法
func (s *RepairskaService) Stop() error {
	return nil
}
