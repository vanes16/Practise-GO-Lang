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
 * Complete the 'countResponseTimeRegressions' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY responseTimes as parameter.
 */

func countResponseTimeRegressions(responseTimes []int32) int32 {
	if len(responseTimes) < 2 {
		return 0
	}

	var count int32 = 0
	var sum int64 = int64(responseTimes[0])

	for i := 1; i < len(responseTimes); i++ {
		avg := sum / int64(i)
		if int64(responseTimes[i]) > avg {
			count++
		}
		sum += int64(responseTimes[i])
	}

	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	responseTimesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var responseTimes []int32

	for i := 0; i < int(responseTimesCount); i++ {
		responseTimesItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		responseTimesItem := int32(responseTimesItemTemp)
		responseTimes = append(responseTimes, responseTimesItem)
	}

	result := countResponseTimeRegressions(responseTimes)

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
