package impl_test

import (
	"context"

	"github.com/solodba/mcube/apps"
	"github.com/solodba/pgtools/apps/chkms"
	"github.com/solodba/pgtools/test/tools"
)

// 全局变量
var (
	svc chkms.Service
	ctx = context.Background()
)

// 初始函数
func init() {
	tools.DevelopmentSet()
	svc = apps.GetInternalApp(chkms.AppName).(chkms.Service)
}
