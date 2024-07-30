package impl_test

import (
	"fmt"
	"testing"
)

func TestGetNextWal(t *testing.T) {
	nextWalName, err := svc.GetNextWal(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(nextWalName)
}

func TestGetMxid(t *testing.T) {
	maxid, err := svc.GetMxid(ctx)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(maxid)
}
