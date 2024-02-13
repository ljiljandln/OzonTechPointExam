package main

import (
	"bufio"
	"fmt"
	"os"
)

func (d *data) checkM() bool {
	if d.m {
		return false
	}
	d.m = true
	return true
}

func (d *data) checkD() bool {
	if !d.m {
		return false
	}
	d.m = false
	return true
}

func (d *data) checkC(i int) bool {
	if !d.m {
		return false
	}
	if i == len(d.str)-1 {
		return false
	}
	d.m = false
	return d.str[i+1] == 'M'
}

func (d *data) checkR(i int) bool {
	if !d.m {
		return false
	} else if len(d.str)-i < 4 {
		return false
	}
	if d.str[i+1:i+4] != "CMD" {
		return false
	}
	d.m = false
	return true
}

type data struct {
	str string
	m   bool
}

func (d *data) solve() bool {
	for i := 0; i < len(d.str); i++ {
		switch rune(d.str[i]) {
		case 'M':
			if !d.checkM() {
				return false
			}
		case 'D':
			if !d.checkD() {
				return false
			}
		case 'C':
			if !d.checkC(i) {
				return false
			}
		default:
			if !d.checkR(i) {
				return false
			}
			i += 3
		}
	}
	return !d.m
}

func doTask(inFile *os.File, outFile *os.File) {
	in := bufio.NewReader(inFile)
	out := bufio.NewWriter(outFile)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	slice := make([]string, t)

	for i := 0; i < t; i++ {
		var str string
		fmt.Fscan(in, &str)
		d := data{str, false}
		ans := d.solve()
		if ans {
			slice[i] = "YES"
		} else {
			slice[i] = "NO"
		}
	}

	for _, ans := range slice {
		fmt.Fprintln(out, ans)
	}
}

func main() {
	doTask(os.Stdin, os.Stdout)
}
