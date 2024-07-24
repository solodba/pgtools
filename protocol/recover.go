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
	fileInfo, err := m.recoversvc.GetLsnAndFile(ctx)
	if err != nil {
		return err
	}
	txInfo, err := m.recoversvc.GetDeleteTxid(ctx, fileInfo)
	if err != nil {
		return err
	}
	err = m.recoversvc.StopDb(ctx)
	if err != nil {
		return err
	}
	err = m.recoversvc.RecoverDbToTxid(ctx, txInfo)
	if err != nil {
		return err
	}
	err = m.recoversvc.StartDb(ctx)
	if err != nil {
		return err
	}
	return nil
}

// RecoverService服务停止方法
func (s *RecoverService) Stop() error {
	return nil
}
