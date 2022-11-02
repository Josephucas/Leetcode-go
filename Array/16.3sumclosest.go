package Array

import (
	"math"
	"sort"
)

func ThreeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var (
		n    = len(nums)
		best = math.MaxInt32
	)
	abs := func(x int) int {
		if x < 0 {
			return -1 * x
		} else {
			return x
		}
	}
	// 根据差值的绝对值来更新答案
	update := func(cur int) {
		if abs(cur-target) < abs(best-target) {
			best = cur
		}
	}
	// 枚举 a
	for i := 0; i < n; i++ {
		// 保证和上一次枚举的元素不相等
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 使用双指针枚举 b 和 c
		j, k := i+1, n-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			// 如果和为 target 直接返回答案
			if sum == target {
				return target
			}
			//这个地方我确实没想到：1.怎么取绝对值，2.把差值的绝对值函数放到循环外面，作为一个闭包函数
			update(sum)
			// 如果和大于 target，移动 c 对应的指针
			if sum > target {
				k0 := k - 1
				//这里没有问题，但是为什么要造出k0这个中间值呢
				// 移动到下一个不相等的元素
				for j < k0 && nums[k0] == nums[k] {
					k0--

				}
				k = k0

			} else {
				// 如果和小于 target，移动 b 对应的指针
				j0 := j + 1
				// 移动到下一个不相等的元素
				for j0 < k && nums[j0] == nums[j] {
					j0++
				}
				j = j0
			}

		}

	}
	return best
}

//以下是我的升级版
//这是我第一次击败100%的用户！！！！
//开心啊哈哈哈
func ThreeSumClosest1(nums []int, target int) int {
	abs := func(x int) int {
		if x < 0 {
			return -1 * x
		} else {
			return x
		}
	}
	min := func(x, y int) int {
		if abs(x-target) < abs(y-target) {
			return x
		}
		return y
	}

	sort.Ints(nums)
	var (
		n = len(nums)

		//对于best的取值是很有技巧的，我就是靠着这个击败100%的，这是我的第二次修剪
		best = min(nums[0]+nums[1]+nums[2], nums[n-1]+nums[n-2]+nums[n-3])
	)

	// 根据差值的绝对值来更新答案
	update := func(cur int) {
		if abs(cur-target) < abs(best-target) {
			best = cur
		}
	}

	//修剪一些特殊情况，如target过大和过小，这是我的第一次修剪
	if target <= nums[0]+nums[1]+nums[2] {
		return nums[0] + nums[1] + nums[2]
	}
	if nums[n-1]+nums[n-2]+nums[n-3] <= target {
		return nums[n-1] + nums[n-2] + nums[n-3]
	}

	// 枚举 a
	for i := 0; i < n; i++ {
		// 保证和上一次枚举的元素不相等
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 使用双指针枚举 b 和 c
		j, k := i+1, n-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			// 如果和为 target 直接返回答案
			if sum == target {
				return target
			}
			//这个地方我确实没想到：1.怎么取绝对值，2.把差值的绝对值函数放到循环外面，作为一个闭包函数
			update(sum)
			// 如果和大于 target，移动 c 对应的指针
			if sum > target {
				k = k - 1
				//这里没有问题，但是为什么要造出k0这个中间值呢
				// 移动到下一个不相等的元素
				for j < k && nums[k] == nums[k+1] {
					k--

				}

			} else {
				// 如果和小于 target，移动 b 对应的指针
				j = j + 1
				// 移动到下一个不相等的元素
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			}

		}

	}
	return best
}
