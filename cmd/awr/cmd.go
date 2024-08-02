package awr

import (
	"github.com/solodba/pgtools/protocol"
	"github.com/spf13/cobra"
)

// 项目启动子命令
var Cmd = &cobra.Command{
	Use:     "awr",
	Short:   "pgtools awr service",
	Long:    "pgtools awr",
	Example: `pgtools awr -U postgres -M 127.0.0.1 -P 5432 -W postgres -D postgres`,
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
	AwrService *protocol.AwrService
}

// 服务结构体初始化函数
func NewServer() *Server {
	return &Server{
		AwrService: protocol.NewAwrService(),
	}
}

// Server服务启动方法
func (s *Server) Start() error {
	if err := s.AwrService.Start(); err != nil {
		return err
	}
	return nil
}

// Server服务停止方法
func (s *Server) Stop() error {
	if err := s.AwrService.Stop(); err != nil {
		return err
	}
	return nil
}
