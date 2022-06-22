package Leetcode

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for _, value := range nums1 {
		if _, ok := m[value]; ok {
			continue
		}
		m[value] = 1
	}
	for _, value := range nums2 {
		if mValue, ok := m[value]; ok && mValue == 1 {
			m[value] = 2
		}
	}
	var ret []int
	for key, value := range m {
		if value == 2 {
			ret = append(ret, key)
		}
	}
	return ret
}
