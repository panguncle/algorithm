package helper

import "testing"

func TestIntPointer(t *testing.T) {
	i := 1
	pi := IntPtr(i)
	if *pi != i {
		t.Fatalf("excepted:%+v, got:%+v", &i, pi)
	}

	pi2 := IntPtr(2)
	if *pi2 != 2 {
		t.Fatalf("excepted:%d, got:%d", 2, *pi2)
	}
}
