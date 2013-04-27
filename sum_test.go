package gotest

import "testing"

func TestSum(t *testing.T) {
	sum := Sum(1, 2)
	if sum != 3 {
		t.Error("error")
	}
}

func TestMultiSum(t *testing.T) {
	sum := MultiSum(1, 2, 3)
	if sum != 6 {
		t.Errorf("%d", sum)
	}
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(i, i)
	}
}

func BenchmarkMultiSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MultiSum(i, i, i, i, i, i)
	}
}
