package protocol

import (
	"github.com/solodba/mcube/apps"
	"github.com/solodba/pgtools/apps/recover"
)

// RecoverService服务结构体
type RecoverService struct {
	recoversvc recover.Service
}

// RecoverService服务结构体构造函数
func NewRecoverService() *RecoverService {
	return &RecoverService{
		recoversvc: apps.GetInternalApp(recover.AppName).(recover.Service),
	}
}

// RecoverService服务启动方法
func (m *RecoverService) Start() error {
	return nil
}

// RecoverService服务停止方法
func (s *RecoverService) Stop() error {
	return nil
}
