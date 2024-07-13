package impl

import (
	"context"
	"fmt"
)

func (i *impl) RepairStandbyKeepalived(ctx context.Context) error {
	cmd := `sed -i "s#priority 100#priority 90#" /etc/keepalived/keepalived.conf;service keepalived start`
	_, err := i.cmdConf.RunShell(cmd)
	if err != nil {
		return err
	}
	fmt.Println("备节点keepalived服务修复完成")
	return nil
}
