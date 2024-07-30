package rectl

import (
	"github.com/solodba/pgtools/protocol"
	"github.com/spf13/cobra"
)

// 项目启动子命令
var Cmd = &cobra.Command{
	Use:     "rectl",
	Short:   "pgtools rectl service",
	Long:    "pgtools rectl",
	Example: `pgtools rectl -u root -w Root_123 -m 127.0.0.1 -p 22`,
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
	RectlService *protocol.RectlService
}

// 服务结构体初始化函数
func NewServer() *Server {
	return &Server{
		RectlService: protocol.NewRectlService(),
	}
}

// Server服务启动方法
func (s *Server) Start() error {
	if err := s.RectlService.Start(); err != nil {
		return err
	}
	return nil
}

// Server服务停止方法
func (s *Server) Stop() error {
	if err := s.RectlService.Stop(); err != nil {
		return err
	}
	return nil
}
