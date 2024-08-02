package protocol

import (
	"github.com/solodba/mcube/apps"
	"github.com/solodba/pgtools/apps/awr"
)

// AwrService服务结构体
type AwrService struct {
	awrsvc awr.Service
}

// AwrService服务结构体构造函数
func NewAwrService() *AwrService {
	return &AwrService{
		awrsvc: apps.GetInternalApp(awr.AppName).(awr.Service),
	}
}

// AwrService服务启动方法
func (m *AwrService) Start() error {
	return nil
}

// AwrService服务停止方法
func (s *AwrService) Stop() error {
	return nil
}
