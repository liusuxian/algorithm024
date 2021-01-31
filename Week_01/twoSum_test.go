package Week_01

import "testing"

// 两数之和

// 暴力求解
// 时间复杂度 O(n^2)
// 空间复杂度 O(1)
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

// 空间换时间
// 时间复杂度 O(n)
// 空间复杂度 O(n)
func twoSum1(nums []int, target int) []int {
	m := map[int]int{}
	for i, num := range nums {
		if v, ok := m[target-num]; ok {
			return []int{v, i}
		}
		m[num] = i
	}

	return []int{}
}

func TestTwoSum(t *testing.T) {
	t.Log(twoSum([]int{2, 2, 7, 11, 15}, 9))
	t.Log(twoSum([]int{3, 2, 4}, 6))
	t.Log(twoSum([]int{3, 3}, 6))
	t.Log(twoSum([]int{3, 3}, 7))
	t.Log(twoSum([]int{-3, 4, 3, 90}, 0))
	t.Log("------------------------")
	t.Log(twoSum1([]int{2, 2, 7, 11, 15}, 9))
	t.Log(twoSum1([]int{3, 2, 4}, 6))
	t.Log(twoSum1([]int{3, 3}, 6))
	t.Log(twoSum1([]int{3, 3}, 7))
	t.Log(twoSum1([]int{-3, 4, 3, 90}, 0))
}
