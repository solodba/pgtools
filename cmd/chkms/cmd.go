package chkms

import (
	"github.com/solodba/pgtools/protocol"
	"github.com/spf13/cobra"
)

// 项目启动子命令
var Cmd = &cobra.Command{
	Use:     "chkms",
	Short:   "pgtools chkms service",
	Long:    "pgtools service",
	Example: `pgtools chkms -U postgres -W postgres -M 127.0.0.1 -P 5432 -D postgres`,
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
	ChkmsService *protocol.ChkmsService
}

// 服务结构体初始化函数
func NewServer() *Server {
	return &Server{
		ChkmsService: protocol.NewChkmsService(),
	}
}

// Server服务启动方法
func (s *Server) Start() error {
	if err := s.ChkmsService.Start(); err != nil {
		return err
	}
	return nil
}

// Server服务停止方法
func (s *Server) Stop() error {
	if err := s.ChkmsService.Stop(); err != nil {
		return err
	}
	return nil
}
