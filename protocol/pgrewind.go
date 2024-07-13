package protocol

import (
	"github.com/solodba/mcube/apps"
	"github.com/solodba/pgtools/apps/pgrewind"
)

// PgrewindService服务结构体
type PgrewindService struct {
	pgrewindsvc pgrewind.Service
}

// PgrewindService服务结构体构造函数
func NewPgrewindService() *PgrewindService {
	return &PgrewindService{
		pgrewindsvc: apps.GetInternalApp(pgrewind.AppName).(pgrewind.Service),
	}
}

// PgrewindService服务启动方法
func (m *PgrewindService) Start() error {
	err := m.pgrewindsvc.PgRewindSyncData(ctx)
	if err != nil {
		return err
	}
	return nil
}

// PgrewindService服务停止方法
func (s *PgrewindService) Stop() error {
	return nil
}
