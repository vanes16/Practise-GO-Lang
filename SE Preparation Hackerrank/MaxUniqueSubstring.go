package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'maxDistinctSubstringLengthInSessions' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING sessionString as parameter.
 */

func maxDistinctSubstringLengthInSessions(sessionString string) int32 {
	// Write your code here
	if sessionString == "*" || sessionString == "" || sessionString == " " {
		return 0
	}
	last := make([]int, 256)
	for i := 0; i < 256; i++ {
		last[i] = -1
	}

	left, maxLen := 0, 0

	for right := 0; right < len(sessionString); right++ {

		if sessionString[right] == '*' {
			for i := 0; i < 256; i++ {
				last[i] = -1
			}
			left = right + 1
			continue
		}

		if last[sessionString[right]] >= left {
			left = last[sessionString[right]] + 1
		}
		last[sessionString[right]] = right

		if right-left+1 > maxLen {
			maxLen = right - left + 1

		}

	}
	return int32(maxLen)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	sessionString := readLine(reader)

	result := maxDistinctSubstringLengthInSessions(sessionString)

	fmt.Printf("%d\n", result)
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
