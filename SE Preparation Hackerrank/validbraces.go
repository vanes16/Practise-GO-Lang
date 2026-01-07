package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'areBracketsProperlyMatched' function below.
 *
 * The function is expected to return a BOOLEAN.
 * The function accepts STRING code_snippet as parameter.
 */

func areBracketsProperlyMatched(code_snippet string) bool {
	// Write your code here
	stack := []rune{}

	// map
	pairs := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for _, ch := range code_snippet {
		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, ch)
			continue
		}
		if ch == ')' || ch == ']' || ch == '}' {
			if len(stack) == 0 {
				return false
			}
			if pairs[stack[len(stack)-1]] != ch {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	code_snippet := readLine(reader)

	result := areBracketsProperlyMatched(code_snippet)

	fmt.Printf("%s\n", (map[bool]string{true: "1", false: "0"})[result])
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
