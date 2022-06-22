package Leetcode

func minWindow(s string, t string) string {
	left, right, count := 0, 0, 0
	minStr := ""
	need := genNeed(t)
	window := map[int32]int{}

	for right < len(s) {
		rightChar := int32(s[right])

		right++
		if need[rightChar] > 0 {
			window[rightChar]++
			if window[rightChar] == need[rightChar] {
				count++
			}
		}

		for count == len(need) {
			leftChar := int32(s[left])
			if need[leftChar] > 0 {
				//[left,right] 是一个子串
				if minStr == "" || len(minStr) > right-left {
					minStr = s[left:right]
				}

				window[leftChar]--
				if window[leftChar] < need[leftChar] {
					count--
					left++
					break
				}
			}

			left++
		}
	}

	return minStr
}

func genNeed(t string) map[int32]int {
	m := make(map[int32]int)
	for _, b := range t {
		m[b]++
	}
	return m
}
