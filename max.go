package gosample

func Max(nums []int) int {
	tmp := 0
	for _, i := range nums {
		if tmp < i {
			tmp = i
		}
	}
	return tmp
}
