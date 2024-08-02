package impl_test

import "testing"

func TestGetSystemInfo(t *testing.T) {
	_, err := svc.GetSystemInfo(ctx)
	if err != nil {
		t.Fatal()
	}
}
