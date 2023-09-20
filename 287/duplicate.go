package duplicate

func FindDuplicate(nums []int) int {
	knownNums := make([]bool, len(nums))

	for _, n := range nums {
		if knownNums[n] {
			return n
		}

		knownNums[n] = true
	}

	return -1
}
