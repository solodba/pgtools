package impl_test

import (
	"testing"

	"github.com/solodba/pgtools/test/tools"
)

func TestGetSystemInfo(t *testing.T) {
	_, err := svc.GetSystemInfo(ctx)
	if err != nil {
		t.Fatal()
	}
}

func TestGetComsumeIoSql(t *testing.T) {
	comsumeTopSqlSet, err := svc.GetComsumeIoSql(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(comsumeTopSqlSet))
}

func TestGetComsumeTimeSql(t *testing.T) {
	comsumeTopSqlSet, err := svc.GetComsumeTimeSql(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(comsumeTopSqlSet))
}

func TestGetComsumeBufferSql(t *testing.T) {
	comsumeTopSqlSet, err := svc.GetComsumeBufferSql(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(comsumeTopSqlSet))
}
