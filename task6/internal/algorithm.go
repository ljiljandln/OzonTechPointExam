package internal

import (
	"bufio"
	"fmt"
)

const (
	upLeft    = 0
	up        = 1
	upRight   = 2
	right     = 3
	downRight = 4
	down      = 5
	downLeft  = 6
	left      = 7
)

var moveMatrix = [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}

type DpData struct {
	dp1  [][]uint8
	dp2  [][]uint8
	dp3  [][]uint8
	dp4  [][]uint8
	n, m int
}

func NewDpData(in *bufio.Reader) DpData {
	var n, m int
	fmt.Fscan(in, &n, &m)
	table := getTable(in, n, m)

	dp1 := getDp1(table)
	dp2 := getDp2(table)
	dp3 := getDp3(table)
	dp4 := getDp4(table)

	return DpData{dp1: dp1, dp2: dp2, dp3: dp3, dp4: dp4, n: n, m: m}
}

func (d *DpData) Solve() (int, int) {
	var res uint8
	var row, col int
	for i := 1; i <= d.n; i++ {
		for j := 1; j <= d.m; j++ {
			var curr uint8
			if i == 1 {
				curr = d.firstLastRowHandler(d.dp4, d.dp3, i, j, downRight, downLeft)
			} else if i == d.n {
				curr = d.firstLastRowHandler(d.dp2, d.dp1, i, j, upRight, upLeft)
			} else if j == 1 {
				curr = min2(prev(d.dp2, i, j, upRight), prev(d.dp4, i, j, upLeft))
			} else if j == d.m {
				curr = min2(prev(d.dp1, i, j, upLeft), prev(d.dp3, i, j, downLeft))
			} else {
				m1 := min2(prev(d.dp1, i, j, upLeft), prev(d.dp2, i, j, upRight))
				m2 := min2(prev(d.dp3, i, j, downLeft), prev(d.dp4, i, j, downRight))
				curr = min2(m1, m2)
			}
			if curr > res {
				res = curr
				row = i
				col = j
			}
		}
	}
	return row, col
}

func (d *DpData) firstLastRowHandler(dp1, dp2 [][]uint8, i, j int, move1, move2 int) (curr uint8) {
	if j == 1 {
		curr = prev(dp1, i, j, move1)
	} else if j == d.m {
		curr = prev(dp2, i, j, move2)
	} else {
		curr = min2(prev(dp1, i, j, move1), prev(dp2, i, j, move2))
	}
	return
}

func getTable(in *bufio.Reader, n, m int) [][]uint8 {
	table := make([][]uint8, n+1)
	for i := 1; i <= n; i++ {
		table[i] = make([]uint8, m+1)
		var str string
		fmt.Fscan(in, &str)

		for j := 1; j <= m; j++ {
			ch := str[j-1] - '0'
			table[i][j] = ch
		}
	}
	return table
}

func prev(dp [][]uint8, i, j int, move int) uint8 {
	i += moveMatrix[move][0]
	j += moveMatrix[move][1]
	return dp[i][j]
}

func dpHelper(dp, t [][]uint8, i, j int, row, col int, move1, move2 int) {
	num := t[i][j]
	if i == row && j == col {
		dp[i][j] = num
	} else if i == row && j != col {
		dp[i][j] = min2(num, prev(dp, i, j, move1))
	} else if i != row && j == col {
		dp[i][j] = min2(num, prev(dp, i, j, move2))
	} else {
		dp1 := prev(dp, i, j, move1)
		dp2 := prev(dp, i, j, move2)
		dp[i][j] = min3(num, dp1, dp2)
	}
}

func getDp1(t [][]uint8) [][]uint8 {
	n, m := len(t), len(t[1])
	dp := make([][]uint8, n)
	for i := 1; i < n; i++ {
		dp[i] = make([]uint8, m)
		for j := 1; j < m; j++ {
			dpHelper(dp, t, i, j, 1, 1, left, up)
		}
	}
	return dp
}

func getDp2(t [][]uint8) [][]uint8 {
	n, m := len(t), len(t[1])
	dp := make([][]uint8, n)
	for i := 1; i < n; i++ {
		dp[i] = make([]uint8, m)
		for j := m - 1; j > 0; j-- {
			dpHelper(dp, t, i, j, 1, m-1, right, up)
		}
	}
	return dp
}

func getDp3(t [][]uint8) [][]uint8 {
	n, m := len(t), len(t[1])
	dp := make([][]uint8, n)
	for i := n - 1; i > 0; i-- {
		dp[i] = make([]uint8, m)
		for j := 1; j < m; j++ {
			dpHelper(dp, t, i, j, n-1, 1, left, down)
		}
	}
	return dp
}

func getDp4(t [][]uint8) [][]uint8 {
	n, m := len(t), len(t[1])
	dp := make([][]uint8, n)
	for i := n - 1; i > 0; i-- {
		dp[i] = make([]uint8, m)
		for j := m - 1; j > 0; j-- {
			dpHelper(dp, t, i, j, n-1, m-1, right, down)
		}
	}
	return dp
}

func min2(a, b uint8) uint8 {
	if a <= b {
		return a
	}
	return b
}

func min3(a, b, c uint8) uint8 {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	}
	return c
}
