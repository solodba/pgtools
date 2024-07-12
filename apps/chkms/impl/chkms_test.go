package impl_test

import "testing"

func TestGetPostgresqlPrimaryStandby(t *testing.T) {
	err := svc.GetPostgresqlPrimaryStandby(ctx)
	if err != nil {
		t.Fatal(err)
	}
}
