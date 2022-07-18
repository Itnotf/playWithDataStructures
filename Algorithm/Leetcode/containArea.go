package Leetcode

import "sort"

func containArea(t [][]int) []int {
	// 找到所有未被包含的点
	outsidePoint := markOutsidePoint(t)
	//遍历t，将包含的点全部设成1
	for i := 0; i < len(t); i++ {
		for j := 0; j < len(t[0]); j++ {
			//如果t[i][j] 在outsidePoint 里，跳过
			if outsidePoint[i][j] == 1 {
				continue
			}
			//否则，将 t[i][j] 值改为1
			t[i][j] = 1
		}
	}

	//求连通的1的最大面积
	s := areaS(t)
	sort.Ints(s)
	return s
}

func areaS(t [][]int) []int {
	var s []int
	visit := genEmptyVisit(t)
	//从第一个点开始，遍历所有连通点，点的个数就是面积
	for i := 0; i < len(t); i++ {
		for j := 0; j < len(t[0]); j++ {
			if visit[i][j] == 0 && t[i][j] == 1 {
				area := 0
				backtrack1(i, j, &area, visit, t)
				s = append(s, area)
			}
		}
	}

	return s
}

func backtrack1(x int, y int, area *int, visit [][]int, t [][]int) {
	//不在范围内
	if x < 0 || y < 0 || x >= len(t) || y >= len(t[0]) {
		return
	}
	//该点为 0 或者已经访问过
	if visit[x][y] == 1 || t[x][y] == 0 {
		return
	}

	visit[x][y] = 1
	*area++

	//看当前点周围4个方向的点
	// (x+1,y) (x-1,y) (x,y+1) (x,y-1)
	backtrack1(x+1, y, area, visit, t)
	backtrack1(x-1, y, area, visit, t)
	backtrack1(x, y+1, area, visit, t)
	backtrack1(x, y-1, area, visit, t)
}

func markOutsidePoint(t [][]int) [][]int {
	m, n := len(t), len(t[0])
	visit := genEmptyVisit(t)
	//遍历4周，如果是0，将可以连通的0加入visit
	for i, j := 0, 0; i < m; i++ {
		//遍历该点周围8个点，如果是0，加入visit
		if visit[i][j] == 0 && t[i][j] == 0 {
			outBacktrack(i, j, visit, t)
		}
	}
	for i, j := 0, n-1; i < m; i++ {
		if visit[i][j] == 0 && t[i][j] == 0 {
			outBacktrack(i, j, visit, t)
		}
	}
	for i, j := 0, 0; j < n; j++ {
		if visit[i][j] == 0 && t[i][j] == 0 {
			outBacktrack(i, j, visit, t)
		}
	}
	for i, j := m-1, 0; j < n; j++ {
		if visit[i][j] == 0 && t[i][j] == 0 {
			outBacktrack(i, j, visit, t)
		}
	}
	return visit
}

func outBacktrack(x int, y int, visit [][]int, t [][]int) {
	//不在范围内
	if x < 0 || y < 0 || x >= len(t) || y >= len(t[0]) {
		return
	}
	//该点为 1 或者已经访问过
	if visit[x][y] == 1 || t[x][y] == 1 {
		return
	}

	visit[x][y] = 1

	//看当前点周围8个点
	// (x+1,y) (x-1,y) (x,y+1) (x,y-1)
	outBacktrack(x+1, y, visit, t)
	outBacktrack(x+1, y+1, visit, t)
	outBacktrack(x-1, y, visit, t)
	outBacktrack(x-1, y-1, visit, t)
	outBacktrack(x, y+1, visit, t)
	outBacktrack(x-1, y+1, visit, t)
	outBacktrack(x, y-1, visit, t)
	outBacktrack(x-1, y-1, visit, t)
}

func genEmptyVisit(t [][]int) [][]int {
	visit := make([][]int, len(t))
	for i := 0; i < len(t); i++ {
		visit[i] = make([]int, len(t[0]))
	}
	return visit
}
