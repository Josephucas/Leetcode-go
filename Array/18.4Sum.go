package Array

import "sort"

func FourSum(nums []int, target int) (quadruplets [][]int) {
	//把这个放在最前面，连sort的时间都省掉了
	if len(nums) < 4 {
		return

	}

	sort.Ints(nums)
	n := len(nums)

	if nums[n-1]+nums[n-2]+nums[n-3]+nums[n-4] < target {
		return
	}

	//这里是一定要小于等于的，不能只是小于号
	//在确定第一个数之后，如果 nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target，
	//说明此时剩下的三个数无论取什么值，四数之和一定大于 target，因此退出第一重循环；

	for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		//在确定第一个数之后，如果nums[i]+nums[n−3]+nums[n−2]+nums[n−1]<target，
		//说明此时剩下的三个数无论取什么值，四数之和一定小于target，因此第一重循环直接进入下一轮，枚举 nums[i+1]；
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}

		//在确定前两个数之后，如果nums[i]+nums[j]+nums[j+1]+nums[j+2]>target，
		//说明此时剩下的两个数无论取什么值，四数之和一定大于target，因此退出第二重循环；
		for j := i + 1; j < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			//在确定前两个数之后，如果nums[i]+nums[j]+nums[n−2]+nums[n−1]<target，
			//说明此时剩下的两个数无论取什么值，四数之和一定小于target，因此第二重循环直接进入下一轮，枚举nums[j+1]。

			//这里是一定不能加等号的，啊啊啊
			if j > i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[n-1]+nums[n-2] < target {
				continue

			}
			
			//到这一步的时候，已经不需要再剪枝了
			for left, right := j+1, n-1; left < right; {
				if sum := nums[i] + nums[j] + nums[left] + nums[right]; sum == target {
					quadruplets = append(quadruplets, []int{nums[i], nums[j], nums[left], nums[right]})

					//找到答案时，双指针时同时靠近的！！！于此同时对重复的进行剪枝
					for left++; left < right && nums[left] == nums[left-1]; left++ {

					}
					for right--; left < right && nums[right] == nums[right+1]; right-- {

					}

				} else if sum < target {
					left++

				} else {
					right--
				}

			}
		}

	}
	return
}
