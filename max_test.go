package gosample

import "testing"

func Test_Max(t *testing.T) {
	nums := []int{0, 2, 3, 5, 4, 6}
	max := Max(nums)
	if max != 6 {
		t.Error("error")
	}
}
