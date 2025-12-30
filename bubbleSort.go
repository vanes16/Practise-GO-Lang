package main

import "fmt"

func bubbleSort(data []int) {
	dataLen := len(data)

	for x := 0; x < dataLen-1; x++ {
		swap := false

		for y := 0; y < dataLen-x-1; y++ {
			if data[y] > data[y+1] {
				data[y], data[y+1] = data[y+1], data[y]
				swap = true

				fmt.Println(data)

			}
		}
		if !swap {
			break
		}
	}

}

func main() {
	data := []int{5, 4, 3, 2, 1, 7, 6}

	// 4 3 2 1 5 6 7  // iteration 0 loop 5x
	// 3 2 1 4 5 6 7 // iteration 1 loop 4x
	// 2 1 3 4 5 6 7 // iteration 2 loop 3x
	// 1 2 3 4 5 6 7 // iteration 3 loop 2x

	lenData := len(data)
	bubbleSort(data)
	fmt.Println(lenData)
	fmt.Println(data)
}
