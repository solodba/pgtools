package impl

import (
	"context"
	"fmt"
)

func (i *impl) RepairStandbyKeepalived(ctx context.Context) error {
	switch i.cmdConf.PgType {
	case "pg11":
		cmd := `sed -i "s#priority 100#priority 90#" /etc/keepalived/keepalived.conf;service keepalived start`
		_, err := i.cmdConf.RunShell(cmd)
		if err != nil {
			return err
		}
		fmt.Println("备节点keepalived服务修复完成")
		return nil
	case "pg13":
		cmd := `service keepalived start`
		_, err := i.cmdConf.RunShell(cmd)
		if err != nil {
			return err
		}
		fmt.Println("备节点keepalived服务修复完成")
		return nil
	default:
		return fmt.Errorf("[%s]该类型数据库不支持", i.cmdConf.PgType)
	}
}
