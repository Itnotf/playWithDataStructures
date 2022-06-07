package algorithm

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	stack := make([]int, 0)
	for _, k := range s {
		if k == '(' || k == '[' || k == '{' {
			stack = push(stack, int(k))
		} else {
			if len(stack) == 0 {
				return false
			}
			s, e := pop(stack)
			stack = s
			if (k == ')' && e != '(') || (k == ']' && e != '[') || (k == '}' && e != '{') {
				return false
			}
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

func pop(arr []int) ([]int, int) {
	length := len(arr)
	return arr[0 : length-1], arr[length-1]
}

func push(arr []int, e int) []int {
	return append(arr, e)
}
