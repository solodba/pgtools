package impl

import (
	"context"
	"fmt"
)

func (i *impl) GetPostgresqlPrimaryStandby(ctx context.Context) error {
	sql := `select count(*) from pg_stat_replication`
	row := i.db.QueryRowContext(ctx, sql)
	var recordCount int
	err := row.Scan(&recordCount)
	if err != nil {
		return err
	}
	fmt.Println(recordCount)
	return nil
}
