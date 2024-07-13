package impl

import (
	"context"
	"fmt"
)

func (i *impl) CheckPromoteLog(ctx context.Context) error {
	cmd := `grep -i promote /etc/keepalived/takeover_pg.log`
	result, err := i.cmdConf.RunShell(cmd)
	if err != nil {
		return fmt.Errorf("takeover_pg.log切换日志不存在")
	}
	fmt.Println(result)
	return nil
}
