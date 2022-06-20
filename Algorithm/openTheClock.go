package algorithm

func openLock(deadends []string, target string) int {
	deadendsMap := make(map[string]int)
	for _, deadend := range deadends {
		deadendsMap[deadend] = 1
	}
	start := "0000"
	if _, ok := deadendsMap[start]; ok {
		return -1
	}

	var q []string
	q = append(q, start)
	deadendsMap[start] = 1

	step := 0
	for len(q) > 0 {
		for _, s := range q {
			q = q[1:]
			if s == target {
				return step
			}
			for _, v := range nextRounds(s, deadendsMap) {
				q = append(q, v)
			}
		}
		step++
	}

	return -1
}

func nextRounds(s string, deadendsMap map[string]int) []string {
	var nextS []string
	newArr := make([]byte, 4)
	arr := []byte(s)

	for i := 0; i < 4; i++ {
		copy(newArr, arr)
		newArr[i] = plus(newArr[i])
		plus := string(newArr)
		if _, ok := deadendsMap[plus]; !ok {
			nextS = append(nextS, plus)
			deadendsMap[plus] = 1
		}

		newArr[i] = sub(sub(newArr[i]))
		sub := string(newArr)
		if _, ok := deadendsMap[sub]; !ok {
			nextS = append(nextS, sub)
			deadendsMap[sub] = 1
		}

	}
	return nextS
}

func plus(n byte) byte {
	if n == '9' {
		return '0'
	}
	return n + 1
}

func sub(n byte) byte {
	if n == '0' {
		return '9'
	}
	return n - 1
}
