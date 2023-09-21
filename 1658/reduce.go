package reduce

func MinOperations(nums []int, x int) int {
	numsCount := len(nums)
	left := 0
	right := numsCount - 1

	var steps []int

	for x > 0 && left <= right {
		x = x - nums[left]
		left += 1
	}

	if x == 0 {
		steps = append(steps, left+numsCount-1-right)
	}

	for left > 0 {
		left -= 1
		x = x + nums[left]

		if x < 0 {
			continue
		}

		for x > 0 && right > left {
			x = x - nums[right]
			right -= 1
		}

		if x == 0 {
			steps = append(steps, left+numsCount-1-right)
		}
	}

	if len(steps) == 0 {
		return -1
	}

	min := steps[0]

	for _, s := range steps {
		if s < min {
			min = s
		}
	}

	return min
}
