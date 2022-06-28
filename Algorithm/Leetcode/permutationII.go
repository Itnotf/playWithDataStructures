package Leetcode

import "sort"

func permutation(s string) []string {
	result := map[string]int{}

	backtrack(result, "", s)

	var ret []string
	for key, _ := range result {
		ret = append(ret, key)
	}
	sort.Strings(ret)
	return ret
}

func backtrack(result map[string]int, path string, s string) {
	if len(s) == 0 {
		result[path] = 1
	}

	for i := 0; i < len(s); i++ {
		path += s[i : i+1]
		backtrack(result, path, s[0:i]+s[i+1:])
		path = path[:len(path)-1]
	}
}
