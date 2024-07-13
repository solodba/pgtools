package repairms

import (
	"github.com/solodba/pgtools/protocol"
	"github.com/spf13/cobra"
)

// 项目启动子命令
var Cmd = &cobra.Command{
	Use:     "repairms",
	Short:   "pgtools repairms service",
	Long:    "pgtools service",
	Example: `pgtools repairms -u root -w Root_123 -m 127.0.0.1 -p 22`,
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
	RepairMsService *protocol.RepairMsService
}

// 服务结构体初始化函数
func NewServer() *Server {
	return &Server{
		RepairMsService: protocol.NewRepairMsService(),
	}
}

// Server服务启动方法
func (s *Server) Start() error {
	if err := s.RepairMsService.Start(); err != nil {
		return err
	}
	return nil
}

// Server服务停止方法
func (s *Server) Stop() error {
	if err := s.RepairMsService.Stop(); err != nil {
		return err
	}
	return nil
}
