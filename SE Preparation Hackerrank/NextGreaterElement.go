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
 * Complete the 'findNextGreaterElementsWithDistance' function below.
 *
 * The function is expected to return a 2D_INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY readings as parameter.
 */

func findNextGreaterElementsWithDistance(readings []int32) [][]int32 {
	// Write your code here
	result := [][]int32{}
	var distance int32
	for i := 0; i < len(readings); i++ {
		distance = 0
		found := false
		for j := i; j < len(readings); j++ {
			if readings[j] > readings[i] {
				found = true
				distance = int32(j - i)
				result = append(result, []int32{readings[j], distance})
				break
			}
			if j == len(readings)-1 {
				if found == false {
					result = append(result, []int32{-1, -1})
				}
			}
		}
	}

	return result

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	readingsCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var readings []int32

	for i := 0; i < int(readingsCount); i++ {
		readingsItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		readingsItem := int32(readingsItemTemp)
		readings = append(readings, readingsItem)
	}

	result := findNextGreaterElementsWithDistance(readings)

	for i, rowItem := range result {
		for j, colItem := range rowItem {
			fmt.Printf("%d", colItem)

			if j != len(rowItem)-1 {
				fmt.Printf(" ")
			}
		}

		if i != len(result)-1 {
			fmt.Printf("\n")
		}
	}

	fmt.Printf("\n")
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
