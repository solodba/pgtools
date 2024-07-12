package tools

import (
	"github.com/solodba/mcube/apps"
	"github.com/solodba/mcube/logger"
	_ "github.com/solodba/pgtools/apps/all"
	"github.com/solodba/pgtools/conf"
)

func LoadConfig() {
	conf.Conf = conf.NewDefaultConfig()
	conf.Conf.PostgreSQL.Username = "postgres"
	conf.Conf.PostgreSQL.Password = "postgres"
	conf.Conf.PostgreSQL.Host = "192.168.1.140"
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
