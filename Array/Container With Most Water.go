package leetcode

//题目大意 #
//给出一个非负整数数组 a1，a2，a3，…… an，每个整数标识一个竖立在坐标轴 x 位置的一堵高度为 ai 的墙，选择两堵墙，和 x 轴构成的容器可以容纳最多的水。
//
//解题思路 #
//这一题也是对撞指针的思路。首尾分别 2 个指针，每次移动以后都分别判断长宽的乘积是否最大。

//在每个状态下，无论长板或短板向中间收窄一格，都会导致水槽 底边宽度 -1−1​ 变短：
//
//若向内 移动短板 ，水槽的短板 min(h[i], h[j])min(h[i],h[j]) 可能变大，因此下个水槽的面积 可能增大 。
//若向内 移动长板 ，水槽的短板 min(h[i], h[j])min(h[i],h[j])​ 不变或变小，因此下个水槽的面积 一定变小 。
//因此，初始化双指针分列水槽左右两端，循环每轮将短板向内移动一格，并更新面积最大值，直到两指针相遇时跳出；即可获得最大面积
//
//作者：jyd
//链接：https://leetcode-cn.com/problems/container-with-most-water/solution/container-with-most-water-shuang-zhi-zhen-fa-yi-do/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func maxArea(height []int) int {
	max, start, end := 0, 0, len(height)-1
	//如果这里的start等于end说明max为0，或者一堆的0，所以这里不能相等
	for start < end {
		width := end - start
		high := 0
		if height[start] < height[end] {
			high = height[start]
			start++
		} else {
			high = height[end]
			end--
		}
		temp := width * high
		if temp > max {
			max = temp
		}
	}
	return max

}
func maxArea1(height []int) int {
	max, start, end := 0, 0, len(height)-1
	//如果这里的start等于end说明max为0，或者一堆的0，所以这里不能相等
	for start < end {
		width := end - start
		high := 0
		if height[start] < height[end] {
			high = height[start]

			start++
		} else {
			high = height[end]
			end--
		}
		temp := width * high
		if temp > max {
			max = temp
		}
	}
	return max

}
