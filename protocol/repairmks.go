package protocol

import (
	"github.com/solodba/mcube/apps"
	"github.com/solodba/pgtools/apps/repairmka"
)

// RepairmkaService服务结构体
type RepairmkaService struct {
	repairmkasvc repairmka.Service
}

// RepairmkaService服务结构体构造函数
func NewRepairmkaService() *RepairmkaService {
	return &RepairmkaService{
		repairmkasvc: apps.GetInternalApp(repairmka.AppName).(repairmka.Service),
	}
}

// RepairskaService服务启动方法
func (m *RepairmkaService) Start() error {
	err := m.repairmkasvc.RepairPriamryKeepalived(ctx)
	if err != nil {
		return err
	}
	return nil
}

// chtab服务停止方法
func (s *RepairmkaService) Stop() error {
	return nil
}
