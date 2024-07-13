package impl_test

import "testing"

func TestCheckPromoteLog(t *testing.T) {
	err := svc.CheckPromoteLog(ctx)
	if err != nil {
		t.Fatal(err)
	}
}
