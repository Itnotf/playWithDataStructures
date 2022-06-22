package Leetcode

func lengthOfLongestSubstring(s string) int {
	var window [256]int
	i, j, longest := 0, 0, 0
	for i < len(s) {
		keyI := s[i]
		window[keyI]++

		if window[keyI] > 1 {
			keyJ := s[j]
			for keyJ != keyI {
				window[keyJ]--
				j++
				keyJ = s[j]
			}
			window[keyJ]--
			j++
		} else {
			if longest < i-j+1 {
				longest = i - j + 1
			}
		}
		i++
	}
	return longest
}
