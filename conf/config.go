package conf

import (
	"database/sql"
	"sync"
)

// 全局配置参数
var (
	Conf *Config
)

// Config结构体
type Config struct {
	PostgreSQL *PostgreSQL
	CmdConf    *CmdConf
}

// PostgreSQL结构体
type PostgreSQL struct {
	Username    string
	Password    string
	Host        string
	Port        int32
	DB          string
	MaxOpenConn int64
	MaxIdleConn int64
	MaxLifeTime int64
	MaxIdleTime int64
	lock        sync.Mutex
	db          *sql.DB
}

// PostgreSQL结构体构造函数
func NewDefaultPostgreSQL() *PostgreSQL {
	return &PostgreSQL{
		Username: "root",
		Password: "123456",
		Host:     "127.0.0.1",
		Port:     3306,
		DB:       "test",
	}
}

// CmdConf结构体
type CmdConf struct {
	Sysuser string
	Syspwd  string
	Syshost string
	Sysport int32
}

// CmdConf结构体构造函数
func NewDefaultCmdConf() *CmdConf {
	return &CmdConf{}
}

// Config结构体构造函数
func NewDefaultConfig() *Config {
	return &Config{
		PostgreSQL: NewDefaultPostgreSQL(),
		CmdConf:    NewDefaultCmdConf(),
	}
}
