package impl

import (
	"database/sql"

	"github.com/solodba/mcube/apps"
	"github.com/solodba/pgtools/apps/chtab"
	"github.com/solodba/pgtools/conf"
)

// 实现Service的结构体
type impl struct {
	cmdConf *conf.CmdConf
	db      *sql.DB
}

// 实现Ioc中心Name方法
func (i *impl) Name() string {
	return chtab.AppName
}

// 实现Ioc中心Conf方法
func (i *impl) Conf() error {
	i.cmdConf = conf.Conf.CmdConf
	db, err := conf.Conf.PostgreSQL.GetDbConn()
	if err != nil {
		return err
	}
	i.db = db
	return nil
}

// 注册初始化函数
func init() {
	apps.RegistryInternalApp(&impl{})
}
