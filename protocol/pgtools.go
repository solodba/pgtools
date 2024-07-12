package protocol

import (
	"github.com/solodba/mcube/apps"
	"github.com/solodba/pgtools/apps/chkps"
)

// pgtools服务结构体
type PgToolsService struct {
	chksvc chkps.Service
}

// pgtools服务结构体构造函数
func NewPgToolsService() *PgToolsService {
	return &PgToolsService{
		chksvc: apps.GetInternalApp(chkps.AppName).(chkps.Service),
	}
}

// pgtools服务启动方法
func (m *PgToolsService) Start() error {

	return nil
}

// pgtools服务停止方法
func (s *PgToolsService) Stop() error {
	return nil
}
