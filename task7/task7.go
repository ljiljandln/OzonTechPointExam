package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type word struct {
	str  string
	hash uint32
}

func getHash(str string) (hash uint32) {
	for i := range str {
		ch := str[i]
		hash |= 1 << ch
	}
	return
}

func getData(in *bufio.Reader) (data []word) {
	var n int
	fmt.Fscan(in, &n)
	data = make([]word, n)
	for i := 0; i < n; i++ {
		var str string
		fmt.Fscan(in, &str)
		w := word{str: str, hash: getHash(str)}
		data[i] = w
	}
	return data
}

func compareStr(s1, s2 string) bool {
	if strings.EqualFold(s1, s2) {
		return true
	}
	wasRev := false
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			if wasRev {
				return false
			} else if i == len(s1)-1 {
				return false
			} else if !(s1[i+1] == s2[i] && s1[i] == s2[i+1]) {
				return false
			} else {
				wasRev = true
				i++
			}
		}
	}
	return true
}

func doTask(inFile, outFile *os.File) {
	in := bufio.NewReader(inFile)
	out := bufio.NewWriter(outFile)
	defer out.Flush()

	d := getData(in)

	var m int
	fmt.Fscan(in, &m)
	for ; m > 0; m-- {
		var str string
		fmt.Fscan(in, &str)
		hash := getHash(str)

		wasFound := false
		for _, w := range d {
			if len(str) == len(w.str) && hash == w.hash && compareStr(str, w.str) {
				wasFound = true
				break
			}
		}
		if wasFound {
			fmt.Fprintln(out, "1")
		} else {
			fmt.Fprintln(out, "0")
		}
	}
}

func main() {
	doTask(os.Stdin, os.Stdout)
}
