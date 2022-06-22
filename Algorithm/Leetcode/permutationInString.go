package Leetcode

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	var m0, m [26]int
	for _, c := range s1 {
		m0[c-'a']++
	}

	i, j := 0, 0
	for j < len(s1)-1 {
		if m0[s2[j]-'a'] > 0 {
			m[s2[j]-'a']++
		}
		j++
	}

	for j < len(s2) {
		charI := s2[i]
		charJ := s2[j]
		if m0[charJ-'a'] > 0 {
			m[charJ-'a']++
		}
		if m0 == m {
			return true
		}

		if m0[charI-'a'] > 0 {
			m[charI-'a']--
		}
		i++
		j++
	}

	return false
}
