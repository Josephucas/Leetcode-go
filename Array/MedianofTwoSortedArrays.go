package Array

import "math"

//题目大意 #
//给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
//
//请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
//
//你可以假设 nums1 和 nums2 不会同时为空。
//
//解题思路 #
//给出两个有序数组，要求找出这两个数组合并以后的有序数组中的中位数。要求时间复杂度为 O(log (m+n))。
//
//这一题最容易想到的办法是把两个数组合并，然后取出中位数。但是合并有序数组的操作是 O(m+n) 的，
//不符合题意。看到题目给的 log 的时间复杂度，很容易联想到二分搜索。
//
//由于要找到最终合并以后数组的中位数，两个数组的总大小也知道，所以中间这个位置也是知道的。只需要二分搜索一个数组中切分的位置，
//另一个数组中切分的位置也能得到。为了使得时间复杂度最小，所以二分搜索两个数组中长度较小的那个数组。
//
//关键的问题是如何切分数组 1 和数组 2 。其实就是如何切分数组 1 。先随便二分产生一个 midA，切分的线何时算满足了中位数的条件呢？
//即，线左边的数都小于右边的数，即，nums1[midA-1] ≤ nums2[midB] && nums2[midB-1] ≤ nums1[midA] 。
//如果这些条件都不满足，切分线就需要调整。如果 nums1[midA] < nums2[midB-1]，说明 midA 这条线划分出来左边的数小了，切分线应该右移；
//如果 nums1[midA-1] > nums2[midB]，说明 midA 这条线划分出来左边的数大了，切分线应该左移。经过多次调整以后，切分线总能找到满足条件的解。
//
//假设现在找到了切分的两条线了，数组 1 在切分线两边的下标分别是 midA - 1 和 midA。数组 2 在切分线两边的下标分别是 midB - 1 和 midB。
//最终合并成最终数组，如果数组长度是奇数，那么中位数就是 max(nums1[midA-1], nums2[midB-1])。如果数组长度是偶数，
//那么中间位置的两个数依次是：max(nums1[midA-1], nums2[midB-1]) 和 min(nums1[midA], nums2[midB])，
//那么中位数就是 (max(nums1[midA-1], nums2[midB-1]) + min(nums1[midA], nums2[midB])) / 2

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//如果 m>n ，那么得出的 jj 有可能是负数。如果m更大没必要取到m
	if len(nums1) > len(nums2) {
		return FindMedianSortedArrays(nums2, nums1)
	}
	m, n := len(nums1), len(nums2)
	left, right := 0, m
	median1, median2 := 0, 0
	for left <= right {
		i := (left + right) / 2
		j := (m+n+1)/2 - i
		//无限小
		nums_im1 := math.MinInt32
		if i != 0 {
			nums_im1 = nums1[i-1]
		}
		//无限大
		nums_i := math.MaxInt32
		if i != m {
			nums_i = nums1[i]
		}
		nums_jm1 := math.MinInt32
		if j != 0 {
			nums_jm1 = nums2[j-1]
		}
		nums_j := math.MaxInt32
		if j != n {
			nums_j = nums2[j]
		}
		//保证数组1的左边最大的数字比数组2右边最小的数字要小，如果大了就得往小二分
		if nums_im1 <= nums_j {
			median1 = max(nums_im1, nums_jm1)
			median2 = min(nums_i, nums_j)
			left = i + 1
		} else {
			right = i - 1
		}
	}

	//如果是偶数就取值中间的平均值
	if (m+n)%2 == 0 {
		return float64(median1+median2) / 2.0
	}
	return float64(median1)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
