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

func TestGetMxid(t *testing.T) {
	maxid, err := svc.GetMxid(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(maxid)
}

func TestGetNextMxidOffset(t *testing.T) {
	nextMxidOffset, err := svc.GetNextMxidOffset(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(nextMxidOffset)
}

func TestGetNextXid(t *testing.T) {
	nextXid, err := svc.GetNextXid(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(nextXid)
}
