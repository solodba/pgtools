package impl_test

import (
	"context"

	"github.com/solodba/mcube/apps"
	"github.com/solodba/pgtools/apps/awr"
	"github.com/solodba/pgtools/test/tools"
)

// 全局变量
var (
	svc awr.Service
	ctx = context.Background()
)

// 初始函数
func init() {
	tools.DevelopmentSet()
	svc = apps.GetInternalApp(awr.AppName).(awr.Service)
}
