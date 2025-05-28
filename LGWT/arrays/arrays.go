package main

func main() {}
func Sum(numbers []int) int {
	var sum int
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return sum
}
func Sumrange(numbers []int) int {
	sum := 0
	for _, i := range numbers {
		sum += i
	}
	return sum
}
func SumAll(numbersToSum ...[]int) []int {
	sums := make([]int, len(numbersToSum))
	for i, v := range numbersToSum {
		sums[i] = Sum(v)
	}
	return sums
}
