package Leetcode

func lengthOfLongestSubstring(s string) int {
	//双指针i,j
	//当前窗口map
	//最长的长度maxLen
	//如果s[i]不在窗口内，将s[i]加入窗口，maxLen = max(len(windows),maxLen)，i++
	//如果s[i]在窗口中，
	// -- 如果s[i] == s[j]，将s[j]移出窗口，j++ ,跳出循环
	// -- 如果s[i] != s[j]，将s[j]移出窗口，j++，继续循环

	i, j, maxLen := 0, 0, 0
	windows := map[byte]int{}

	for i < len(s) {
		chi := s[i]
		if _, ok := windows[chi]; !ok {
			windows[chi] = 1
			maxLen = max(i-j+1, maxLen)
			i++
			continue
		}

		for s[i] != s[j] {
			delete(windows, s[j])
			j++
		}
		delete(windows, s[j])
		j++
	}
	return maxLen
}
