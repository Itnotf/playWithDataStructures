package Leetcode

import (
	"sort"
)

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i int, j int) bool {
		return points[i][0] < points[j][0]
	})
	cur := points[0]
	count := 1
	for i := 1; i < len(points); i++ {
		if points[i][0] <= cur[1] {
			//第一段边界缩减为短的那个
			if points[i][1] < cur[1] {
				cur[1] = points[i][1]
			}
			continue
		}
		count++
		cur = points[i]
	}
	return count
}
