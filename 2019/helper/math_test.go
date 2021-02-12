package helper

import (
	"testing"
)

func TestMaxInt(t *testing.T) {
	testcases := []struct {
		Source []int
		Ret    int
	}{
		{
			[]int{1},
			1,
		},
		{
			[]int{1, 2, 3, 4},
			4,
		},
		{
			[]int{4, 3, 5, 2, 1},
			5,
		},
	}

	for _, testcase := range testcases {
		ret := MaxInt(testcase.Source...)
		if ret != testcase.Ret {
			t.Fatalf("testcase:%#v, excepted:%d, got:%d\n", testcase, testcase.Ret, ret)
		}
	}

	testcaseE := []int{}
	defer func() {
		if e := recover(); e != nil {
			if err, ok := e.(error); ok && err == ErrListEmpty {
				return
			}
		}

		t.Fatalf("testcase:%#v, excepted panic, but not", testcaseE)
	}()

	MaxInt(testcaseE...)
}

func TestMinInt(t *testing.T) {
	testcases := []struct {
		Input []int
		Min   int
	}{
		{
			Input: []int{1},
			Min:   1,
		},
		{
			Input: []int{1, 3, 4, 2, 0, -1},
			Min:   -1,
		},
	}

	for _, testcase := range testcases {
		ret := MinInt(testcase.Input...)
		if ret != testcase.Min {
			t.Fatalf("testcase:%#v, excepted:%d, got:%d\n", testcase, testcase.Min, ret)
		}
	}

	testcaseE := []int{}
	defer func() {
		if e := recover(); e != nil {
			if err, ok := e.(error); ok && err == ErrListEmpty {
				return
			}
		}

		t.Fatalf("testcase:%#v, excepted panic, but not", testcaseE)
	}()

	MinInt(testcaseE...)
}
