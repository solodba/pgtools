package impl_test

import (
	"context"

	"github.com/solodba/mcube/apps"
	"github.com/solodba/pgtools/apps/chklog"
	"github.com/solodba/pgtools/test/tools"
)

// 全局变量
var (
	svc chklog.Service
	ctx = context.Background()
)

// 初始函数
func init() {
	tools.DevelopmentSet()
	svc = apps.GetInternalApp(chklog.AppName).(chklog.Service)
}
