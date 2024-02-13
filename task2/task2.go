package main

import (
	"bufio"
	"fmt"
	"os"
)

func doTask(inFile, outFile *os.File) {
	in := bufio.NewReader(inFile)
	out := bufio.NewWriter(outFile)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	slice := make([]float64, t)

	for i := 0; i < t; i++ {
		var n, p, sum int
		fmt.Fscan(in, &n, &p)
		for ; n > 0; n-- {
			var a int
			fmt.Fscan(in, &a)
			sum += (p * a) % 100
		}
		slice[i] = float64(sum) / 100.0
	}

	for _, res := range slice {
		fmt.Fprintf(out, "%.2f\n", res)
	}
}

func main() {
	doTask(os.Stdin, os.Stdout)
}
