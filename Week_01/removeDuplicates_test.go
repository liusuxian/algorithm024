package Week_01

import "testing"

// 删除排序数组中的重复项

// 解法：双指针
// 时间复杂度 O(n)
// 空间复杂度 O(1)
func removeDuplicates(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	p := 0
	for q := 1; q < len(nums); q++ {
		if nums[p] != nums[q] {
			p++
			if q-p > 0 {
				nums[p] = nums[q]
			}
		}
	}

	return p + 1
}

func TestRemoveDuplicates(t *testing.T) {
	t.Log(removeDuplicates([]int{}))
	t.Log(removeDuplicates([]int{1, 1, 2}))
	t.Log(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
