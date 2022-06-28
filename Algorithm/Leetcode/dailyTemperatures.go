package Leetcode

func dailyTemperatures(temperatures []int) []int {
	//stack为下标+温度的数组
	var stack [][]int

	ret := make([]int, len(temperatures))
	for index, temperature := range temperatures {
		//如果栈顶小于当前温度，出栈，计算天数
		for len(stack) > 0 && stack[len(stack)-1][1] < temperature {
			ret[stack[len(stack)-1][0]] = index - stack[len(stack)-1][0]
			stack = stack[0 : len(stack)-1]
		}

		//否则，入栈
		stack = append(stack, []int{index, temperature})
	}
	return ret
}
