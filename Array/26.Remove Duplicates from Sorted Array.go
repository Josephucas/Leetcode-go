package Array

//给你一个 升序排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。
//
//由于在某些语言中不能改变数组的长度，所以必须将结果放在数组nums的第一部分。更规范地说，如果在删除重复项之后有 k 个元素，那么 nums 的前 k 个元素应该保存最终结果。
//
//将最终结果插入 nums 的前 k 个位置后返回 k 。
//
//不要使用额外的空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
//
//判题标准:
//
//系统会用下面的代码来测试你的题解:
//
//int[] nums = [...]; // 输入数组
//int[] expectedNums = [...]; // 长度正确的期望答案
//
//int k = removeDuplicates(nums); // 调用
//
//assert k == expectedNums.length;
//for (int i = 0; i < k; i++) {
//assert nums[i] == expectedNums[i];
//}
//如果所有断言都通过，那么您的题解将被 通过。
//
//
//
//示例 1：
//
//输入：nums = [1,1,2]
//输出：2, nums = [1,2,_]
//解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。
//示例 2：
//
//输入：nums = [0,0,1,1,1,2,2,3,3,4]
//输出：5, nums = [0,1,2,3,4]
//解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。不需要考虑数组中超出新长度后面的元素。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/remove-duplicates-from-sorted-array
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//思路
//抄leetcode官网的

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

func removeDuplicates1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1

	}
	//如果只有数组中只有一个数字或者数组中的数字为自然数数列，后来发现这样时不对的，因为你用一头一尾相减也无法保证这是个自然数组
	//if nums[n-1]-nums[0] == n-1 {
	//	return n
	//}
	slow := 0
	for fast := 1; fast < n; fast++ {
		//如果fast指针后面的数字都是一样的，那么直接返回结果
		//但是这里有一个问题当你的fast=n-1的时候，下面的式子是一定成立的而如果fast和fast-1的结果不一样，那么按照算法，你得把他们调过头过来
		//所以干脆把fast==n-1的情况剔除
		if fast < n-1 && nums[fast] == nums[n-1] {
			//这里还有一个情况时如果fast和fast-1的值是一样的，那么就只能是slow+1而不是slow+2！！！！
			if nums[fast-1] == nums[fast] {
				return slow + 1
			}
			//这里我也忘记调换了！！！
			nums[slow+1] = nums[fast]
			return slow + 2
		}

		if nums[fast] != nums[fast-1] {
			nums[slow+1] = nums[fast]
			slow++
		}

	}
	return slow + 1
}
