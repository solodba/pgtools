package conf

import (
	"fmt"
	"net"
	"time"

	"golang.org/x/crypto/ssh"
)

// 创建SSH连接
func (s *CmdConf) CreateSSHConn() (*ssh.Client, error) {
	config := ssh.ClientConfig{
		User: s.Sysuser,
		Auth: []ssh.AuthMethod{ssh.Password(s.Syspwd)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Timeout: 60 * time.Minute,
	}
	addr := fmt.Sprintf("%s:%d", s.Syshost, s.Sysport)
	client, err := ssh.Dial("tcp", addr, &config)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// 创建SSH会话用于执行shell命令
func (s *CmdConf) RunShell(shell string) (string, error) {
	client, err := s.CreateSSHConn()
	if err != nil {
		return "", err
	}
	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	output, err := session.CombinedOutput(shell)
	if err != nil {
		return "", err
	}
	return string(output), nil
}
