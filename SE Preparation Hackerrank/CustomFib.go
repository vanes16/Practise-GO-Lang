package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'getAutoSaveInterval' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts INTEGER n as parameter.
 */

func getAutoSaveInterval(n int32) int64 {
	// Write your code here

	if n == 0 {
		return 1
	} else if n == 1 {
		return 2
	}
	result := make([]int, 2, n)
	result[0] = 1
	result[1] = 2
	for i := 1; i < int(n); i++ {
		result = append(result, result[i]+result[i-1])
	}
	return int64(result[n])
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	result := getAutoSaveInterval(n)

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
