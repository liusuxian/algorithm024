package Week_01

import "testing"

// 移动零

// 时间复杂度: O(n)
// 空间复杂度: O(1)
// 2次循环
func moveZeroes(nums []int) {
	if nums == nil {
		return
	}

	//第一次遍历的时候，j指针记录非0的个数，只要是非0的统统都赋给nums[j]
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}
	//非0元素统计完了，剩下的都是0了
	//所以第二次遍历把末尾的元素都赋为0即可
	for i := j; i < len(nums); i++ {
		nums[i] = 0
	}
}

// 时间复杂度: O(n)
// 空间复杂度: O(1)
// 1次循环
func moveZeroes1(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			if i != j {
				nums[i] = 0
			}
			j++
		}
	}
}

// 时间复杂度: O(n)
// 空间复杂度: O(1)
// 1次循环 i、j互相交换
func moveZeroes2(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}
}

func TestMoveZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	t.Log(nums)
	moveZeroes1(nums)
	t.Log(nums)
	moveZeroes2(nums)
	t.Log(nums)
}
