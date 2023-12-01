package product

func ProductExceptSelf(nums []int) []int {
	var result = make([]int, len(nums))

	result[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		result[i] = nums[i] * result[i-1]
	}

	postfix := 1

	for i := len(nums) - 1; i > 0; i-- {
		result[i] = result[i-1] * postfix
		postfix *= nums[i]
	}
	result[0] = postfix

	return result
}
