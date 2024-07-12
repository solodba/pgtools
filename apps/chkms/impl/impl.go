package impl

import (
	"database/sql"

	"github.com/solodba/mcube/apps"
	"github.com/solodba/pgtools/apps/chkms"
	"github.com/solodba/pgtools/conf"
)

// 实现Service的结构体
type impl struct {
	db *sql.DB
}

// 实现Ioc中心Name方法
func (i *impl) Name() string {
	return chkms.AppName
}

// 实现Ioc中心Conf方法
func (i *impl) Conf() error {
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
