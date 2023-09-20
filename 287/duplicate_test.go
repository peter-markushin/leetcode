package duplicate

import "testing"

type dupeTest struct {
	Nums     []int
	Expected int
}

var dupeTests = []dupeTest{
	{[]int{1, 3, 4, 2, 2}, 2},
	{[]int{3, 1, 3, 4, 2}, 3},
	{[]int{2, 2, 2, 2, 2}, 2},
	{[]int{1, 4, 4, 2, 4}, 4},
	{[]int{7, 9, 7, 4, 2, 8, 7, 7, 1, 5}, 7},
}

func TestFindDuplicate(t *testing.T) {
	for _, test := range dupeTests {
		if result := FindDuplicate(test.Nums); result != test.Expected {
			t.Errorf("Output %d not equal to expected %d", result, test.Expected)
		}
	}
}
