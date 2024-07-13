package protocol

import (
	"github.com/solodba/mcube/apps"
	"github.com/solodba/pgtools/apps/rebuild"
)

// RebuildService服务结构体
type RebuildService struct {
	rebuildsvc rebuild.Service
}

// RebuildService服务结构体构造函数
func NewRebuildService() *RebuildService {
	return &RebuildService{
		rebuildsvc: apps.GetInternalApp(rebuild.AppName).(rebuild.Service),
	}
}

// RebuildService服务启动方法
func (m *RebuildService) Start() error {
	err := m.rebuildsvc.RebuildStandby(ctx)
	if err != nil {
		return err
	}
	return nil
}

// RebuildService服务停止方法
func (s *RebuildService) Stop() error {
	return nil
}
