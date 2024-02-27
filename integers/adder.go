package integers

// Add takes two integers and returns the sum of them
func Add(x, y int) int {
	return x + y
}

func Sum(numbers []int) int {
	var holder int
	for _, value := range numbers {
		holder += value
	}
	return holder
}

func SumAll(numbersToSum ...[]int) []int {
	// length_of_collection := len(numbersToSum)
	var sums []int
	for _, v := range numbersToSum {
		sums = append(sums, Sum(v))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int{
	var sums []int
	for _, arr := range numbersToSum{
		sums = append(sums, Sum(arr[1:]))
	}
	return sums
}
