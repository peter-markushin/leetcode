package product

import (
	"reflect"
	"testing"
)

type ProductTest struct {
	Nums     []int
	Expected []int
}

var productTestData = []ProductTest{
	{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
	{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
}

func TestProductExceptSelf(t *testing.T) {
	for _, test := range productTestData {
		if result := ProductExceptSelf(test.Nums); !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("Expected %+v for %+v, but got %+v", test.Expected, test.Nums, result)
		}
	}
}

func BenchmarkProductExceptSelf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProductExceptSelf(productTestData[1].Nums)
	}
}
