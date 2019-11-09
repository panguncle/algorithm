package helper

import "testing"

func TestIntSliceEqual(t *testing.T) {
	testcases := []struct {
		Source1  []int
		Source2  []int
		Excepted bool
	}{
		{
			Source1:  nil,
			Source2:  nil,
			Excepted: true,
		},
		{
			Source1:  []int{1, 2},
			Source2:  []int{1, 2},
			Excepted: true,
		},
		{
			Source1:  []int{1, 2},
			Source2:  []int{2, 1},
			Excepted: false,
		},
		{
			Source1:  []int{1},
			Source2:  []int{1, 2},
			Excepted: false,
		},
		{
			Source1:  nil,
			Source2:  []int{},
			Excepted: true,
		},
	}

	for _, testcase := range testcases {
		ret := IntSliceEqual(testcase.Source1, testcase.Source2)
		if ret != testcase.Excepted {
			t.Fatalf("testcase:%+v, excepted:%t, got:%t\n", testcase, testcase.Excepted, ret)
		}
	}
}
