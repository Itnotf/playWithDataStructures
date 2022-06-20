package algorithm

func minWindow(s string, t string) string {
	left, right, count := 0, 0, 0
	minStr := ""
	need := genNeed(t)
	window := map[int32]int{}

	for right < len(s) {
		rightChar := int32(s[right])
		//right char not in t
		if need[rightChar] == 0 {
			right++
			continue
		}
		//right char in t
		window[rightChar]++

		if window[rightChar] != need[rightChar] {
			right++
			continue
		}
		count++
		if count != len(need) {
			right++
			continue
		}

		//s[left:right+1] contains t
		for left <= right {
			leftChar := int32(s[left])
			if need[leftChar] == 0 {
				left++
				continue
			}
			//[left,right] 是一个子串
			if minStr == "" || len(minStr) > right-left+1 {
				minStr = s[left : right+1]
			}

			window[leftChar]--
			if window[leftChar] < need[leftChar] {
				count--
				left++
				break
			}
			left++
		}
		right++
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
