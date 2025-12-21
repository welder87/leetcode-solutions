package problem2446

import (
	"fmt"
	"testing"
)

func TestHaveConflict(t *testing.T) {
	testHaveConflict(t, haveConflict)
}

func testHaveConflict(t *testing.T, fn func([]string, []string) bool) {
	testCases := []struct {
		event1 []string
		event2 []string
		ans    bool
	}{
		// preset cases
		// |-----|
		//       |-----|
		{[]string{"01:15", "02:00"}, []string{"02:00", "03:00"}, true},
		// |-----|
		//     |-----|
		{[]string{"01:00", "02:00"}, []string{"01:20", "03:00"}, true},
		// |-----|
		//         |-----|
		{[]string{"10:00", "11:00"}, []string{"14:00", "15:00"}, false},
		// |-------|
		//   |--|
		{[]string{"01:37", "14:20"}, []string{"05:06", "06:17"}, true},
		// common cases
		//      |-----|
		// |-----|
		{[]string{"01:20", "03:00"}, []string{"01:00", "02:00"}, true},
		// |-----|
		//       |-----|
		{[]string{"01:15", "01:17"}, []string{"01:17", "02:00"}, true},
		{[]string{"21:15", "21:17"}, []string{"21:17", "23:00"}, true},
		//       |-----|
		// |-----|
		{[]string{"01:17", "02:00"}, []string{"01:15", "01:17"}, true},
		{[]string{"21:17", "23:00"}, []string{"21:15", "21:17"}, true},
		// |-----|
		// |-----|
		{[]string{"01:15", "02:00"}, []string{"01:15", "02:00"}, true},
		// |-----|
		//         |-----|
		{[]string{"23:15", "23:16"}, []string{"23:00", "23:14"}, false},
		//         |-----|
		// |-----|
		{[]string{"23:00", "23:14"}, []string{"23:15", "23:16"}, false},
		//    |--|
		// |-------|
		{[]string{"14:15", "14:16"}, []string{"01:00", "23:55"}, true},
		// |-------|
		//   |--|
		{[]string{"01:00", "23:55"}, []string{"14:15", "14:16"}, true},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v %v", testCase.event1, testCase.event2)
		t.Run(testName, func(t *testing.T) {
			ans := fn(testCase.event1, testCase.event2)
			if testCase.ans != ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
