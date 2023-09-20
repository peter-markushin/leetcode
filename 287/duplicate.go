package duplicate

func FindDuplicate(nums []int) int {
	for i, n := range nums {
		for _, c := range nums[i+1:] {
			if c == n {
				return c
			}
		}
	}

	return -1
}
