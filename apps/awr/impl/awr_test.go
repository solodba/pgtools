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

func TestGetComsumeTempSql(t *testing.T) {
	comsumeTopSqlSet, err := svc.GetComsumeTempSql(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(comsumeTopSqlSet))
}

func TestGetPgWalFileInfo(t *testing.T) {
	walFileInfo, err := svc.GetPgWalFileInfo(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(walFileInfo))
}

func TestGetPgLockInfo(t *testing.T) {
	lockInfoSet, err := svc.GetPgLockInfo(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(lockInfoSet))
}

func TestGetPgVacuumInfo(t *testing.T) {
	vacuumInfo, err := svc.GetPgVacuumInfo(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(vacuumInfo))
}

func TestGetPgRoleInfo(t *testing.T) {
	roleInfoSet, err := svc.GetPgRoleInfo(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(roleInfoSet))
}
