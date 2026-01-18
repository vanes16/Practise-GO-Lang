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
 * Complete the 'countAffordablePairs' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY prices
 *  2. INTEGER budget
 */

func countAffordablePairs(prices []int32, budget int32) int32 {
	// Write your code here

	n := len(prices)
	if n < 2 {
		return 0
	}
	left, right := 0, n-1
	var output int64 = 0
	for left < right {
		if prices[left]+prices[right] <= budget {
			output += int64(right - left)
			left++
		} else {
			right--
		}

	}
	return int32(output)

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	pricesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var prices []int32

	for i := 0; i < int(pricesCount); i++ {
		pricesItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		pricesItem := int32(pricesItemTemp)
		prices = append(prices, pricesItem)
	}

	budgetTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	budget := int32(budgetTemp)

	result := countAffordablePairs(prices, budget)

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
