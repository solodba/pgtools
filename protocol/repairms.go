package protocol

import (
	"github.com/solodba/mcube/apps"
	"github.com/solodba/pgtools/apps/repairms"
)

// RepairmsService服务结构体
type RepairMsService struct {
	repairmssvc repairms.Service
}

// RepairmsService服务结构体构造函数
func NewRepairMsService() *RepairMsService {
	return &RepairMsService{
		repairmssvc: apps.GetInternalApp(repairms.AppName).(repairms.Service),
	}
}

// RepairmsService服务启动方法
func (m *RepairMsService) Start() error {
	err := m.repairmssvc.RepairPrimaryStandby(ctx)
	if err != nil {
		return err
	}
	return nil
}

// chtab服务停止方法
func (s *RepairMsService) Stop() error {
	return nil
}
