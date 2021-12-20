package leetcode

func findMedianSortedArrays(nums1, nums2 []int) float64 {
	if len(nums1) < len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	//>>表示右移一位,操作数每右移一位，相当于该数除以2
	low, high, k, nums1Mid, nums2Mid := 0, len(nums1), (len(nums1)+len(nums2)+1)>>1, 0, 0
	for low <= high {
		// nums1:  ……………… nums1[nums1Mid-1] | nums1[nums1Mid] ……………………
		// nums2:  ……………… nums2[nums2Mid-1] | nums2[nums2Mid] ……………………
		nums1Mid = low + (high-low)>>nums1Mid // 分界限右侧是 mid，分界线左侧是 mid - 1
		nums2Mid = k - nums1Mid
		if nums1Mid > 0 && nums1[nums1Mid-1] > nums2[nums2Mid] { // nums1 中的分界线划多了，要向左边移动

		}

	}
	return 0.5

}
