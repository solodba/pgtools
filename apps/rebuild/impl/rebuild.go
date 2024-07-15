package impl

import (
	"context"
	"fmt"
)

func (i *impl) RebuildStandby(ctx context.Context) error {
	switch i.cmdConf.PgType {
	case "pg11":
		cmd := "mv /data/postgres/data /data/postgres/data_bak"
		_, err := i.cmdConf.RunShell(cmd)
		if err != nil {
			return err
		}
		cmd = fmt.Sprintf(`su - postgres -c "/data/postgres/bin/pg_basebackup -h %s -p %d -U postgres -D /data/postgres/data -X stream -P"`, i.cmdConf.PrimaryIp, i.cmdConf.PrimaryPort)
		_, err = i.cmdConf.RunShell(cmd)
		if err != nil {
			return err
		}
		cmd = fmt.Sprintf(`cat >> /data/postgres/data/recovery.conf <<EOF
standby_mode = on
primary_conninfo = 'host=%s port=%d user=postgres'
recovery_target_timeline = 'latest'
EOF`, i.cmdConf.PrimaryIp, i.cmdConf.PrimaryPort)
		_, err = i.cmdConf.RunShell(cmd)
		if err != nil {
			return err
		}
		cmd = "chown postgres.postgres /data/postgres/data/recovery.conf"
		_, err = i.cmdConf.RunShell(cmd)
		if err != nil {
			return err
		}
		cmd = `su - postgres -c "pg_ctl -D /data/postgres/data start > /dev/null 2>&1"`
		_, err = i.cmdConf.RunShell(cmd)
		if err != nil {
			return err
		}
		fmt.Println("postgresql备库修复成功")
		return nil
	case "pg13":
		cmd := "mv /data/postgres/data /data/postgres/data_bak"
		_, err := i.cmdConf.RunShell(cmd)
		if err != nil {
			return err
		}
		cmd = fmt.Sprintf(`su - postgres -c "/data/postgres/bin/pg_basebackup -h %s -p %d -U postgres -D /data/postgres/data -X stream -P"`, i.cmdConf.PrimaryIp, i.cmdConf.PrimaryPort)
		_, err = i.cmdConf.RunShell(cmd)
		if err != nil {
			return err
		}
		cmd = "touch /data/postgres/data/standby.signal;chown postgres.postgres /data/postgres/data/standby.signal"
		_, err = i.cmdConf.RunShell(cmd)
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
	default:
		return fmt.Errorf("[%s]该类型数据库不支持", i.cmdConf.PgType)
	}
}
