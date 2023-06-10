package main

import "testing"

type Args struct {
	src []int64
}

func TestGetMaxSubarray(t *testing.T) {
	src1 := []int{-2, -3, 4, -1, -2, 1, 5, -3}
	res1 := maxSubArray(src1)
	if res1 != 7 {
		t.Errorf("getMaxSubarray(%v) exp %v got %v", src1, 7, res1)
	}

	src2 := []int{-2, -3, 4, -1, -2, 1, 5, -3, 4, -1, -2, 1, 10}
	res2 := maxSubArray(src2)
	if res2 != 16 {
		t.Errorf("getMaxSubarray(%v) exp %v got %v", src2, 16, res2)
	}

	src3 := []int{5, 4, -1, 7, 8}
	res3 := maxSubArray(src3)
	if res3 != 23 {
		t.Errorf("getMaxSubarray(%v) exp %v got %v", src3, 23, res3)
	}
}
