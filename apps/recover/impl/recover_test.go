package impl_test

import (
	"fmt"
	"testing"
)

func TestGetLsnAndFile(t *testing.T) {
	fileInfo, err := svc.GetLsnAndFile(ctx)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fileInfo)
}

func TestGetDeleteTxid(t *testing.T) {
	fileInfo, err := svc.GetLsnAndFile(ctx)
	if err != nil {
		t.Fatal(err)
	}
	txInfo, err := svc.GetDeleteTxid(ctx, fileInfo)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(txInfo)
}

func TestRecover(t *testing.T) {
	fileInfo, err := svc.GetLsnAndFile(ctx)
	if err != nil {
		t.Fatal(err)
	}
	txInfo, err := svc.GetDeleteTxid(ctx, fileInfo)
	if err != nil {
		t.Fatal(err)
	}
	err = svc.StopDb(ctx)
	if err != nil {
		t.Fatal(err)
	}
	err = svc.RecoverDbToTxid(ctx, txInfo)
	if err != nil {
		t.Fatal(err)
	}
	err = svc.StartDb(ctx)
	if err != nil {
		t.Fatal(err)
	}
}
