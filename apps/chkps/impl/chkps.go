package impl

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

func (i *impl) CheckPostgresqlProcess(context.Context) error {
	cmd := "ps -ef |grep -v grep|grep -i postgres: | wc -l"
	result, err := i.cmdConf.RunShell(cmd)
	if err != nil {
		return err
	}
	processNum, err := strconv.Atoi(strings.Trim(result, "\n"))
	if err != nil {
		return err
	}
	if processNum < 6 {
		return fmt.Errorf("[%s]节点postgresql进程异常", i.cmdConf.Syshost)
	}
	fmt.Printf("[%s]节点postgresql进程正常\n", i.cmdConf.Syshost)
	return nil
}
