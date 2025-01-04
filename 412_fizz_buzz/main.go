package problem412

import "strconv"


func fizzBuzz(n int) []string {
	result := make([]string, 0, n+1)
	for i := 1; i <= n; i++ {
		var spam string
		if i % 3 == 0 && i % 5 == 0 {
			spam = "FizzBuzz"
		} else if i % 3 == 0 {
			spam = "Fizz"
		} else if i % 5 == 0 {
			spam = "Buzz"
		} else {
			spam = strconv.Itoa(i)
		}
		result = append(result, spam)
	}
	return result
}
