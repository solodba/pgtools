package impl

import (
	"context"
	"fmt"
)

func (i *impl) UpdatePrimaryStandbyTable(ctx context.Context) error {
	switch i.cmdConf.PgType {
	case "pg13":
		sql := `update public.pg_cluster_takeover set (new_primary,takeover_happen,status)=('0.0.0.0','no','clear')`
		_, err := i.db.ExecContext(ctx, sql)
		if err != nil {
			return err
		}
		fmt.Println("更新主从信息表public.pg_cluster_takeover成功")
		return nil
	default:
		return fmt.Errorf("[%s]该类型数据库不支持", i.cmdConf.PgType)
	}
}
