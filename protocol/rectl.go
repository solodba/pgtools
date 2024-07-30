package protocol

import (
	"fmt"

	"github.com/solodba/mcube/apps"
	"github.com/solodba/pgtools/apps/rectl"
)

// RectlService服务结构体
type RectlService struct {
	rectlsvc rectl.Service
}

// RectlService服务结构体构造函数
func NewRectlService() *RectlService {
	return &RectlService{
		rectlsvc: apps.GetInternalApp(rectl.AppName).(rectl.Service),
	}
}

// RectlService服务启动方法
func (m *RectlService) Start() error {
	fmt.Println("=====================================================================================================================================")
	_, err := m.rectlsvc.GenRebuildControlFileCmd(ctx)
	if err != nil {
		return err
	}
	fmt.Println("=====================================================================================================================================")
	return nil
}

// RectlService服务停止方法
func (s *RectlService) Stop() error {
	return nil
}
