package problem2710

import (
	"fmt"
	"testing"
)

func TestRemoveTrailingZeros(t *testing.T) {
	testCases := []struct {
		num string
		ans string
	}{
		{
			num: "51230100",
			ans: "512301",
		},
		{
			num: "123",
			ans: "123",
		},
		{
			num: "10000",
			ans: "1",
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.num)
		t.Run(testName, func(t *testing.T) {
			ans := removeTrailingZeros(testCase.num)
			if testCase.ans != ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
