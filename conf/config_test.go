package conf_test

import (
	"testing"

	"github.com/solodba/pgtools/conf"
)

func TestGetDbConn(t *testing.T) {
	conf.Conf = conf.NewDefaultConfig()
	conf.Conf.PostgreSQL.Username = "postgres"
	conf.Conf.PostgreSQL.Password = "postgres"
	conf.Conf.PostgreSQL.Host = "192.168.1.140"
	conf.Conf.PostgreSQL.Port = 5432
	conf.Conf.PostgreSQL.DB = "postgres"
	conf.Conf.PostgreSQL.MaxOpenConn = 50
	conf.Conf.PostgreSQL.MaxIdleConn = 10
	conf.Conf.PostgreSQL.MaxLifeTime = 600
	conf.Conf.PostgreSQL.MaxIdleTime = 300
	conn, err := conf.Conf.PostgreSQL.GetDbConn()
	if err != nil {
		t.Fatal(err)
	}
	row, err := conn.Query("select client_addr,state,sync_state from pg_stat_replication;")
	if err != nil {
		t.Fatal(err)
	}
	defer row.Close()
	for row.Next() {
		var clientAddr, state, syncState string
		err = row.Scan(&clientAddr, &state, &syncState)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(clientAddr, state, syncState)
	}
}
