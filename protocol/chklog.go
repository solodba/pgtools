package protocol

import (
	"github.com/solodba/mcube/apps"
	"github.com/solodba/pgtools/apps/chklog"
)

// chklog服务结构体
type ChklogService struct {
	chklogsvc chklog.Service
}

// ChklogService服务结构体构造函数
func NewChklogService() *ChklogService {
	return &ChklogService{
		chklogsvc: apps.GetInternalApp(chklog.AppName).(chklog.Service),
	}
}

// chklog服务启动方法
func (m *ChklogService) Start() error {
	err := m.chklogsvc.CheckPromoteLog(ctx)
	if err != nil {
		return err
	}
	return nil
}

// chkps服务停止方法
func (s *ChklogService) Stop() error {
	return nil
}
