package Week_01

import "testing"

// 时间复杂度 O(n)
// 空间复杂度 O(n)
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			return digits
		}
	}

	digits = make([]int, len(digits)+1)
	digits[0] = 1
	return digits
}

// 时间复杂度 O(n)
// 空间复杂度 O(n)
func plusOne1(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		digits[i] %= 10
		if digits[i] != 0 {
			return digits
		}
	}

	digits = make([]int, len(digits)+1)
	digits[0] = 1
	return digits
}

func TestPlusOne(t *testing.T) {
	t.Log(plusOne([]int{1, 2, 3}))
	t.Log(plusOne([]int{9, 9, 9}))
	t.Log(plusOne1([]int{1, 2, 3}))
	t.Log(plusOne1([]int{9, 9, 9}))
}
