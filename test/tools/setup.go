package tools

import (
	"github.com/solodba/mcube/apps"
	"github.com/solodba/mcube/logger"
	_ "github.com/solodba/pgtools/apps/all"
	"github.com/solodba/pgtools/conf"
)

func LoadConfig() {
	conf.Conf = conf.NewDefaultConfig()
	conf.Conf.CmdConf.Sysuser = "root"
	conf.Conf.CmdConf.Syspwd = "Root_123"
	conf.Conf.CmdConf.Syshost = "119.45.225.100"
	conf.Conf.CmdConf.Sysport = 22
	conf.Conf.PostgreSQL.Username = "postgres"
	conf.Conf.PostgreSQL.Password = "postgres"
	conf.Conf.PostgreSQL.Host = "119.45.61.68"
	conf.Conf.PostgreSQL.Port = 5432
	conf.Conf.PostgreSQL.DB = "postgres"
}

func DevelopmentSet() {
	LoadConfig()
	err := apps.InitInternalApps()
	if err != nil {
		logger.L().Panic().Msgf("initial object config error, err: %s", err.Error())
	}
}
