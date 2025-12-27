package sumsubarray

func MaximumSubarraySum(numbers []int) int {
	currentSum := 0
	highSum := 0

	for _, num := range numbers {
		currentSum += num

		if currentSum < 0 {
			currentSum = 0
		}

		if currentSum > highSum {
			highSum = currentSum
		}
	}

	return highSum
}
