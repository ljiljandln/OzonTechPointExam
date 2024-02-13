package main

import (
	"bufio"
	"fmt"
	"os"
	"task6/internal"
)

func main() {
	doTask(os.Stdin, os.Stdout)
}

func doTask(inFile, outFile *os.File) {
	in := bufio.NewReader(inFile)
	out := bufio.NewWriter(outFile)
	defer out.Flush()

	var t int

	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		d := internal.NewDpData(in)
		row, col := d.Solve()
		fmt.Fprintf(out, "%d %d\n", row, col)
	}
}
