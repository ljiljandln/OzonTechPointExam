package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

const (
	ActualOut = "test_data/out.txt"
)

type testData struct {
	number      int
	in          string
	expectedOut string
}

func getTestData() []testData {
	data := make([]testData, 41)
	for i := 1; i <= 41; i++ {
		in := fmt.Sprintf("test_data/%d", i)
		out := fmt.Sprintf("test_data/%d.a", i)
		data[i-1] = testData{i, in, out}
	}
	return data
}

func TestTask(t *testing.T) {
	tests := getTestData()

	for _, test := range tests {
		solve(test.in)
		expected, err := os.Open(test.expectedOut)
		if err != nil {
			t.Errorf("can't open file [%s]", test.expectedOut)
		}

		actual, err := os.Open(ActualOut)
		if err != nil {
			t.Errorf("can't open file [%s]", ActualOut)
		}

		compareResult(t, expected, actual, test.number)

		expected.Close()
		actual.Close()
	}
}

func compareResult(t *testing.T, expected, actual *os.File, testNumber int) {
	exp := bufio.NewReader(expected)
	act := bufio.NewReader(actual)

	line := 1
	for {
		s1, _, err := exp.ReadLine()
		if err != nil {
			return
		}
		s2, _, _ := act.ReadLine()
		str1 := string(s1)
		str2 := string(s2)
		if !strings.EqualFold(str1, str2) {
			t.Errorf("testNumber: [%d], line [%d], need: [%s], get: [%s]\n", testNumber, line, str1, str2)
		}
		line++
	}
}

func solve(path string) {
	inFile, _ := os.Open(path)
	defer inFile.Close()

	outFile, _ := os.Create(ActualOut)
	defer outFile.Close()

	doTask(inFile, outFile)
}
