package median

var ptr1, ptr2, len1, len2 int

func getMin(nums1 []int, nums2 []int) int {
	var res int
	if ptr1 < len1 && ptr2 < len2 {
		if nums1[ptr1] < nums2[ptr2] {
			res = nums1[ptr1]
			ptr1 += 1
		} else {
			res = nums2[ptr2]
			ptr2 += 1
		}
	} else if ptr1 == len1 {
		res = nums2[ptr2]
		ptr2 += 1
	} else if ptr2 == len2 {
		res = nums1[ptr1]
		ptr1 += 1
	}

	return res
}

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	ptr1, ptr2 = 0, 0
	len1, len2 = len(nums1), len(nums2)

	if (len1+len2)%2 == 1 {
		for i := 0; i < (len1+len2)/2; i += 1 {
			getMin(nums1, nums2)
		}

		return float64(getMin(nums1, nums2))
	}

	for i := 0; i < ((len1+len2)/2)-1; i += 1 {
		getMin(nums1, nums2)
	}

	return (float64(getMin(nums1, nums2)) + float64(getMin(nums1, nums2))) / 2
}
