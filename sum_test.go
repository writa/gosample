package gotest

import "testing"

func TestSum(t *testing.T) {
	sum := Sum(1, 2)
	if sum != 3 {
		t.Error("error")
	}
}
