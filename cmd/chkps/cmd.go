package chkps

import (
	"github.com/solodba/pgtools/protocol"
	"github.com/spf13/cobra"
)

// 项目启动子命令
var Cmd = &cobra.Command{
	Use:     "chkps",
	Short:   "pgtools chkps service",
	Long:    "pgtools service",
	Example: `pgtools chkps -U postgres -M 127.0.0.1 -P 5432 -D postgres`,
	RunE: func(cmd *cobra.Command, args []string) error {
		srv := NewServer()
		if err := srv.Start(); err != nil {
			return err
		}
		return nil
	},
}

// 服务结构体
type Server struct {
	PgToolsService *protocol.PgToolsService
}

// 服务结构体初始化函数
func NewServer() *Server {
	return &Server{
		PgToolsService: protocol.NewPgToolsService(),
	}
}

// Server服务启动方法
func (s *Server) Start() error {
	if err := s.PgToolsService.Start(); err != nil {
		return err
	}
	return nil
}

// Server服务停止方法
func (s *Server) Stop() error {
	if err := s.PgToolsService.Stop(); err != nil {
		return err
	}
	return nil
}
