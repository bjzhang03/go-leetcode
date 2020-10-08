package medium_0740

import (
	"testing"
)

func TestDeleteAndEarn(t *testing.T) {

	var cases = []struct {
		nums     []int
		expected int
	}{
		{[]int{3, 4, 2}, 6},
		{[]int{2, 2, 3, 3, 3, 4}, 9},
		{[]int{1, 8, 5, 9, 6, 9, 4, 1, 7, 3, 3, 6, 3, 3, 8, 2, 6, 3, 2, 2, 1, 2, 9, 8, 7, 1, 1, 10, 6, 7, 3, 9, 6, 10, 5, 4, 10, 1, 6, 7, 4, 7, 4, 1, 9, 5, 1, 5, 7, 5},
			138},
		{[]int{12, 32, 93, 17, 100, 72, 40, 71, 37, 92, 58, 34, 29, 78, 11, 84, 77, 90, 92, 35, 12, 5, 27, 92, 91, 23, 65, 91, 85, 14, 42, 28, 80, 85, 38, 71, 62, 82,
			66, 3, 33, 33, 55, 60, 48, 78, 63, 11, 20, 51, 78, 42, 37, 21, 100, 13, 60, 57, 91, 53, 49, 15, 45, 19, 51, 2, 96, 22, 32, 2, 46, 62, 58, 11, 29, 6, 74, 38, 70,
			97, 4, 22, 76, 19, 1, 90, 63, 55, 64, 44, 90, 51, 36, 16, 65, 95, 64, 59, 53, 93}, 3451},
		{[]int{37, 6, 8, 34, 67, 54, 13, 26, 41, 54, 58, 34, 96, 40, 52, 59, 95, 61, 39, 30, 76, 99, 93, 34, 63, 77, 37, 47, 74, 65, 85, 93, 20, 43, 29, 60, 63,
			17, 28, 73, 49, 1, 71, 99, 93, 46, 29, 1, 44, 93, 64, 84, 73, 2, 10, 22, 87, 14, 70, 32, 58, 20, 87, 57, 17, 55, 55, 15, 16, 38, 67, 98, 78, 61, 13, 92, 32,
			75, 64, 78, 25, 85, 34, 51, 28, 100, 30, 10, 45, 65, 52, 13, 80, 35, 8, 84, 1, 60, 11, 54, 92, 22, 26, 54, 30, 97, 54, 62, 59, 92, 64, 21, 69, 88, 27, 73,
			20, 42, 5, 52, 93, 46, 71, 75, 63, 77, 18, 27, 14, 45, 72, 80, 36, 30, 89, 49, 79, 18, 24, 39, 9, 30, 27, 69, 7, 100, 56, 30, 77, 89, 97, 20, 65, 38, 17, 46,
			19, 92, 84, 99, 21, 49, 62, 52, 19, 78, 47, 62, 79, 29, 64, 36, 7, 9, 69, 80, 20, 24, 78, 93, 54, 79, 54, 96, 72, 76, 5, 63, 33, 20, 32, 36, 69, 69, 11, 35,
			71, 79, 66, 46}, 6238},
	}

	for _, val := range cases {
		actual := deleteAndEarn(val.nums)
		if actual != val.expected {
			t.Errorf("Test Failed! input := %v, expected := %d, actual:= %d", val.nums, val.expected, actual)
		}
	}
}
