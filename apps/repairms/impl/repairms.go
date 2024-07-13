package impl

import (
	"context"
	"fmt"
)

func (i *impl) RepairPrimaryStandby(ctx context.Context) error {
	cmd := "touch /data/postgres/data/standby.signal;chown postgres.postgres /data/postgres/data/standby.signal"
	_, err := i.cmdConf.RunShell(cmd)
	if err != nil {
		return err
	}
	cmd = "sed -i '/^primary_conninfo/ s/^/#/' /data/postgres/data/postgresql.conf"
	_, err = i.cmdConf.RunShell(cmd)
	if err != nil {
		return err
	}
	newPrimary := fmt.Sprintf(`primary_conninfo = \'user=postgres channel_binding=disable host=%s port=%d sslmode=disable sslcompression=0 ssl_min_protocol_version=TLSv1.2 gssencmode=disable krbsrvname=postgres target_session_attrs=any\'`, i.cmdConf.PrimaryIp, i.cmdConf.PrimaryPort)
	cmd = fmt.Sprintf(`echo %s >> /data/postgres/data/postgresql.conf`, newPrimary)
	_, err = i.cmdConf.RunShell(cmd)
	if err != nil {
		return err
	}
	cmd = `service postgresql start`
	_, err = i.cmdConf.RunShell(cmd)
	if err != nil {
		return err
	}
	fmt.Println("postgresql备库修复成功")
	return nil
}
