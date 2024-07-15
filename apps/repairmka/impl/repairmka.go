package impl

import (
	"context"
	"fmt"
)

func (i *impl) RepairPriamryKeepalived(ctx context.Context) error {
	switch i.cmdConf.PgType {
	case "pg11":
		cmd := `sed -i "s#priority 90#priority 100#" /etc/keepalived/keepalived.conf;service keepalived reload`
		_, err := i.cmdConf.RunShell(cmd)
		if err != nil {
			return err
		}
		fmt.Println("主节点keepalived服务修复完成")
		return nil
	default:
		return fmt.Errorf("[%s]该类型数据库不支持", i.cmdConf.PgType)
	}
}
