package Array

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
//https://leetcode.cn/problems/container-with-most-water/solution/sheng-zui-duo-shui-de-rong-qi-by-leetcode-solution/

//这里采用双指针,然后左右两个短板中，短的那一板移动，
//因为高的那一帮子不管怎么移动，水槽里面的水永远比最初的要少，所以为了找到最大的边界，必须移动短板

func MaxArea(height []int) int {
	//这个减去一不能忘记
	max, start, end := 0, 0, len(height)-1

	for start < end {
		width := height[end] - height[start]
		higt := 0
		if height[start] < height[end] {
			higt = height[start]
			start++
		} else {
			higt = height[end]
			end--
		}
		temp := higt * width
		if temp > max {
			max = temp
		}

	}
	return max

}
