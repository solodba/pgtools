package impl

import (
	"context"
	"fmt"
	"strings"
)

func (i *impl) GetPostgresqlPrimaryStandby(ctx context.Context) error {
	sql := `select pg_is_in_recovery();`
	var pgRole string
	row := i.db.QueryRowContext(ctx, sql)
	err := row.Scan(&pgRole)
	if err != nil {
		return err
	}
	if strings.Trim(pgRole, "\n") == "false" {
		fmt.Printf("当前节点为主节点\n")
	}
	if strings.Trim(pgRole, "\n") == "true" {
		fmt.Printf("当前节点为备节点\n")
	}
	return nil
}
