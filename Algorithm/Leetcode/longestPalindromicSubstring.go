package Leetcode

func longestPalindrome(s string) string {

	maxL, maxR := 0, 0

	for i := 0; i < len(s); i++ {

		l := i
		r := i

		for l-1 >= 0 && r+1 <= len(s)-1 && s[l-1] == s[r+1] {
			l--
			r++
		}

		if r-l > maxR-maxL {
			maxL = l
			maxR = r
		}

	}

	for i := 0; i < len(s)-1; i++ {
		//如果有两个一样的
		if s[i+1] == s[i] {
			l := i
			r := i + 1

			for l-1 >= 0 && r+1 <= len(s)-1 && s[l-1] == s[r+1] {
				l--
				r++
			}

			if r-l > maxR-maxL {
				maxL = l
				maxR = r
			}
		}

	}

	return s[maxL : maxR+1]
}
