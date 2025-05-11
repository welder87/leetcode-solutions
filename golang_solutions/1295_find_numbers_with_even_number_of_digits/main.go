package problem1295

import (
	"strconv"
)

// Time complexity: O(n log m). Space complexity: O(1).
func findNumbers(nums []int) int {
	ans := 0
	for _, num := range nums {
		val := num
		counter := 1
		for val > 9 {
			val /= 10
			counter++
		}
		if counter%2 == 0 {
			ans++
		}
	}
	return ans
}

// Time complexity: O(n log m). Space complexity: O(1).
func findNumbersV1(nums []int) int {
	ans := 0
	for _, num := range nums {
		if s := strconv.Itoa(num); len(s)%2 == 0 {
			ans++
		}
	}
	return ans
}

// Time complexity: O(n). Space complexity: O(1).
func findNumbersV2(nums []int) int {
	ans := 0
	for _, num := range nums {
		if (num >= 10 && num <= 99) || (num >= 1000 && num <= 9999) || num == 100000 {
			ans++
		}
	}
	return ans
}
