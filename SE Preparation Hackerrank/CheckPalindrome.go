package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'isAlphabeticPalindrome' function below.
 *
 * The function is expected to return a BOOLEAN.
 * The function accepts STRING code as parameter.
 */

func isAlphabeticPalindrome(code string) bool {
	// Write your code
	var a byte = 'a'
	var b byte = 'A'
	fmt.Println(a)
	fmt.Println(b)

	extString := make([]byte, 0)

	for i := 0; i < len(code); i++ {
		c := code[i]

		if c >= 'A' && c <= 'Z' {
			c = c + 32
			extString = append(extString, c)
		} else if c >= 'a' && c <= 'z' {
			extString = append(extString, c)
		}
	}

	left, right := 0, len(extString)-1
	for left < right {
		if extString[left] != extString[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	code := readLine(reader)

	result := isAlphabeticPalindrome(code)

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
