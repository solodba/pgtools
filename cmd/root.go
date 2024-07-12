package cmd

import (
	"github.com/solodba/mcube/apps"
	"github.com/solodba/mcube/logger"
	"github.com/solodba/mcube/version"
	_ "github.com/solodba/pgtools/apps/all"
	"github.com/solodba/pgtools/cmd/chkps"
	"github.com/solodba/pgtools/conf"
	"github.com/spf13/cobra"
)

// 全局参数
var (
	showVersion bool
	SysUser     string
	SysPwd      string
	SysHost     string
	SysPort     int32
	PgUser      string
	PgPwd       string
	PgHost      string
	PgDb        string
	PgPort      int32
	PgType      string
)

// 根命令
var RootCmd = &cobra.Command{
	Use:     "pgtools [chkps|repairps]",
	Short:   "pgtools service",
	Long:    "pgtools service",
	Example: "pgtools -v",
	RunE: func(cmd *cobra.Command, args []string) error {
		if showVersion {
			logger.L().Info().Msgf(version.ShortVersion())
			return nil
		}
		return cmd.Help()
	},
}

// 加载全局配置
func LoadConfigFromCmd() {
	conf.Conf = conf.NewDefaultConfig()
	conf.Conf.PostgreSQL.Username = PgUser
	conf.Conf.PostgreSQL.Password = PgPwd
	conf.Conf.PostgreSQL.Host = PgHost
	conf.Conf.PostgreSQL.Port = PgPort
	conf.Conf.PostgreSQL.DB = PgDb
	conf.Conf.CmdConf.Sysuser = SysUser
	conf.Conf.CmdConf.Syspwd = SysPwd
	conf.Conf.CmdConf.Syshost = SysHost
	conf.Conf.CmdConf.Sysport = SysPort
}

// 初始化函数
func Initial() {
	LoadConfigFromCmd()
	err := apps.InitInternalApps()
	cobra.CheckErr(err)
}

// 执行函数
func Execute() {
	cobra.OnInitialize(Initial)
	RootCmd.AddCommand(chkps.Cmd)
	err := RootCmd.Execute()
	cobra.CheckErr(err)
}

// 初始化函数
func init() {
	RootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "show project binlog parse version")
	RootCmd.PersistentFlags().StringVarP(&SysUser, "sysuser", "u", "root", "connect host machine username")
	RootCmd.PersistentFlags().StringVarP(&SysPwd, "syspwd", "w", "123456", "connect host machine password")
	RootCmd.PersistentFlags().StringVarP(&SysHost, "syshost", "m", "127.0.0.1", "connect host machine ip address")
	RootCmd.PersistentFlags().Int32VarP(&SysPort, "sysport", "p", 22, "connect host machine port")
	RootCmd.PersistentFlags().StringVarP(&PgUser, "pguser", "U", "postgres", "connect postgresql username")
	RootCmd.PersistentFlags().StringVarP(&PgPwd, "pgpwd", "W", "postgres", "connect postgresql password")
	RootCmd.PersistentFlags().StringVarP(&PgHost, "pghost", "M", "127.0.0.1", "connect postgresql host")
	RootCmd.PersistentFlags().Int32VarP(&PgPort, "pgport", "P", 5432, "connect postgresql port")
	RootCmd.PersistentFlags().StringVarP(&PgDb, "pgdb", "D", "postgres", "connect postgresql database")
	RootCmd.PersistentFlags().StringVarP(&PgType, "pgtype", "T", "pg13", "connect postgresql database type")
}
