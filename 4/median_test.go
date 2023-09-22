package median

import "testing"

type medianTest struct {
	Nums1    []int
	Nums2    []int
	Expected float64
}

var medianTestData = []medianTest{
	{[]int{1, 3}, []int{2}, 2},
	{[]int{1, 2}, []int{3, 4}, 2.5},
}

func TestFindMedianSortedArrays(t *testing.T) {
	for _, test := range medianTestData {
		if result := FindMedianSortedArrays(test.Nums1, test.Nums2); result != test.Expected {
			t.Errorf("Expected %f for %+v, %+v, but got %f", test.Expected, test.Nums1, test.Nums2, result)
		}
	}
}

func BenchmarkFindMedianSortedArrays(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindMedianSortedArrays(medianTestData[1].Nums1, medianTestData[1].Nums2)
	}
}
