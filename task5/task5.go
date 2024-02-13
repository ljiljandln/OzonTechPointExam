package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type dir struct {
	Name     string   `json:"dir"`
	Files    []string `json:"files"`
	Folders  []dir    `json:"folders"`
	infected bool
}

func (d *dir) search(infected bool) int {
	count := 0
	d.infected = infected
	if d.infected {
		count += len(d.Files)
	} else {
		for _, file := range d.Files {
			ext := filepath.Ext(file)
			if strings.EqualFold(ext, ".hack") {
				count += len(d.Files)
				d.infected = true
				break
			}
		}
	}
	for _, dir := range d.Folders {
		count += dir.search(d.infected)
	}
	return count
}

func getData(in *bufio.Reader) []byte {
	var n int
	fmt.Fscan(in, &n)
	var data []byte
	in.ReadLine()
	for i := 0; i < n; i++ {
		line, _, _ := in.ReadLine()
		data = append(data, line...)
	}
	return data
}

func doTask(inFile *os.File, outFile *os.File) {
	in := bufio.NewReader(inFile)
	out := bufio.NewWriter(outFile)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		var d dir
		json.Unmarshal(getData(in), &d)
		fmt.Fprintln(out, d.search(false))
	}
}

func main() {
	doTask(os.Stdin, os.Stdout)
}
