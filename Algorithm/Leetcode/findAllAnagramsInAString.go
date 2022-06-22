package Leetcode

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}

	var need, windows [26]int
	var ret []int
	for _, c := range p {
		need[c-'a']++
	}

	i, j := 0, 0
	for i < len(p)-1 {
		k := s[i] - 'a'
		if need[k] > 0 {
			windows[k]++
		}
		i++
	}
	for i < len(s) {
		front := s[i] - 'a'
		end := s[j] - 'a'

		if need[front] > 0 {
			windows[front]++
		}
		if need == windows {
			ret = append(ret, j)
		}
		if need[end] > 0 {
			windows[end]--
		}

		i++
		j++
	}
	return ret
}
