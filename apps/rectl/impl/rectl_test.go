package impl_test

import (
	"testing"
)

func TestGetNextWal(t *testing.T) {
	nextWalName, err := svc.GetNextWal(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(nextWalName)
}
