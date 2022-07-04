package Leetcode

func maxArea(height []int) int {
	// 双指针i,j
	// 面积 maxS
	// 如果 height[i]<height[j] , i++ 否则 j--

	i, j := 0, len(height)-1
	var maxS int

	for i != j {
		maxS = max(calcS(height, i, j), maxS)

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return maxS
}

func calcS(height []int, left int, right int) int {
	return min(height[left], height[right]) * (right - left)
}
