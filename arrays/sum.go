package arrays

// Sum accepts a slice of integers and returns the sum
func Sum(numbers []int) int {
	// []int = slice, [x]int = array - two different types
	// so a slice is really just an un-sized array?
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

// SumAll accepts any number of slices of integers and returns a slice of their sums
func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}
