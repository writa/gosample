// gotest is go sample
package gotest

// Sum is i + j 
func Sum(i int, j int) int {
	return i + j
}

// Sum is i...
func MultiSum(i ...int) int {
	sum := 0
	for _, c := range i {
		sum = sum + c
	}
	return sum
}
