package problem412

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	testCases := []struct {
		num int
		ans []string
	}{
		{num: 3, ans: []string{"1", "2", "Fizz"}},
		{num: 5, ans: []string{"1", "2", "Fizz", "4", "Buzz"}},
		{
			num: 15,
			ans: []string{
				"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz",
			},
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.num)
		t.Run(testName, func(t *testing.T) {
			ans := fizzBuzz(testCase.num)
			if !reflect.DeepEqual(testCase.ans, ans) {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
