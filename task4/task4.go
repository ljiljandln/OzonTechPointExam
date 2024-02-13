package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const (
	A = 'a'
	B = 'b'
)

type data struct {
	table          [][]rune
	n, m           int
	xa, ya, xb, yb int
}

func newData(in *bufio.Reader) (d data) {
	fmt.Fscan(in, &d.n, &d.m)
	d.table = make([][]rune, d.n)

	for i := range d.table {
		d.table[i] = make([]rune, d.m)
		var line string
		fmt.Fscan(in, &line)
		for j := range d.table[i] {
			ch := rune(line[j])
			if ch == 'A' {
				d.xa = i
				d.ya = j
			} else if ch == 'B' {
				d.xb = i
				d.yb = j
			}
			d.table[i][j] = ch
		}
	}
	return d
}

func (d *data) printRes(out *bufio.Writer) {
	for _, row := range d.table {
		fmt.Fprintln(out, string(row))
	}
}

func (d *data) wayUpAndLeft(x, y int, ch rune) {
	if x == 0 && y == 0 {
		return
	}
	if y%2 == 1 {
		y--
		d.table[x][y] = ch
	}
	for x > 0 {
		x--
		d.table[x][y] = ch
	}
	for y > 0 {
		y--
		d.table[x][y] = ch
	}
}

func (d *data) wayDownAndRight(x, y int, ch rune) {
	if x == d.n-1 && y == d.m-1 {
		return
	}
	if y%2 == 1 {
		y++
		d.table[x][y] = ch
	}
	for x < d.n-1 {
		x++
		d.table[x][y] = ch
	}
	for y < d.m-1 {
		y++
		d.table[x][y] = ch
	}
}

func (d *data) findWay() {
	ra := math.Sqrt(float64(d.xa*d.xa + d.ya*d.ya))
	rb := math.Sqrt(float64(d.xb*d.xb + d.yb*d.yb))
	if ra < rb {
		d.wayUpAndLeft(d.xa, d.ya, A)
		d.wayDownAndRight(d.xb, d.yb, B)
	} else {
		d.wayUpAndLeft(d.xb, d.yb, B)
		d.wayDownAndRight(d.xa, d.ya, A)
	}
}

func doTask(inFile *os.File, outFile *os.File) {
	in := bufio.NewReader(inFile)
	out := bufio.NewWriter(outFile)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for ; t > 0; t-- {
		d := newData(in)
		d.findWay()
		d.printRes(out)
	}
}

func main() {
	doTask(os.Stdin, os.Stdout)
}
