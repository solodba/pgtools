package impl

import (
	"context"
	"fmt"
)

func (i *impl) RepairPriamryKeepalived(ctx context.Context) error {
	cmd := `sed -i "s#priority 90#priority 100#" /etc/keepalived/keepalived.conf;service keepalived reload`
	_, err := i.cmdConf.RunShell(cmd)
	if err != nil {
		return err
	}
	fmt.Println("主节点keepalived服务修复完成")
	return nil
}
