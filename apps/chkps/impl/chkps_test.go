package impl_test

import "testing"

func TestCheckPostgresqlProcess(t *testing.T) {
	err := svc.CheckPostgresqlProcess(ctx)
	if err != nil {
		t.Fatal(err)
	}
}
