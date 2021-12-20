package leetcode

func removeDuplicateElement(number []int) []int {
	result := make([]int, 0, len(number))
	temp := map[int]struct{}{}
	for _, item := range number {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

//func RemoveDuplicates() {
//
//}
//
//func RemoveDuplicates(nums []int) int {
//
//}
