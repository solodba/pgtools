package impl

import (
	"github.com/solodba/mcube/apps"
	"github.com/solodba/pgtools/apps/recover"
	"github.com/solodba/pgtools/conf"
)

// 实现Service的结构体
type impl struct {
	cmdConf *conf.CmdConf
}

// 实现Ioc中心Name方法
func (i *impl) Name() string {
	return recover.AppName
}

// 实现Ioc中心Conf方法
func (i *impl) Conf() error {
	i.cmdConf = conf.Conf.CmdConf
	return nil
}

// 注册初始化函数
func init() {
	apps.RegistryInternalApp(&impl{})
}
