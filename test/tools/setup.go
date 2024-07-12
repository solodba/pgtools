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
	conf.Conf.CmdConf.Syshost = "192.168.1.140"
	conf.Conf.CmdConf.Sysport = 22
}

func DevelopmentSet() {
	LoadConfig()
	err := apps.InitInternalApps()
	if err != nil {
		logger.L().Panic().Msgf("initial object config error, err: %s", err.Error())
	}
}
