package Leetcode

func romanToInt(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	sum := 0
	for i := len(s) - 1; i >= 0; i-- {
		if i+1 < len(s) {
			if (s[i] == 'I' && (s[i+1] == 'V' || s[i+1] == 'X')) || (s[i] == 'X' && (s[i+1] == 'L' || s[i+1] == 'C')) || (s[i] == 'C' && (s[i+1] == 'D' || s[i+1] == 'M')) {
				sum -= m[s[i]]
			} else {
				sum += m[s[i]]
			}
		} else {
			sum += m[s[i]]
		}
	}
	return sum
}
