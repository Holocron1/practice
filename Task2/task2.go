package main

import "fmt"

// FilterEven возвращает новый срез, содержащий

// только чётные числа из nums в том же порядке.

func FilterEven(nums []int) []int {

	result := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			result = append(result, nums[i])
		}
	}

	return result

}

func main() {

	in := []int{5, 4, 9, 2, 7, 6}
	fmt.Println(FilterEven(in))
	// Ожидаемый вывод: [4 2 6]

}
