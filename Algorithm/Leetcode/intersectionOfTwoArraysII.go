package Leetcode

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for _, value := range nums1 {
		m[value] += 1
	}
	var ret []int
	for _, value := range nums2 {
		if _, ok := m[value]; ok && m[value] > 0 {
			m[value] -= 1
			ret = append(ret, value)
		}
	}
	return ret
}
