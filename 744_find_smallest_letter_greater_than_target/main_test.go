package problem744

import (
	"fmt"
	"testing"
)

func TestNextGreatestLetter(t *testing.T) {
	testCases := []struct {
		letters []byte
		target  byte
		ans     byte
	}{
		{
			letters: []byte("cfj"),
			target:  byte('a'),
			ans:     byte('c'),
		},
		{
			letters: []byte("cfj"),
			target:  byte('c'),
			ans:     byte('f'),
		},
		{
			letters: []byte("xxyy"),
			target:  byte('z'),
			ans:     byte('x'),
		},
		{
			letters: []byte("xy"),
			target:  byte('z'),
			ans:     byte('x'),
		},
		{
			letters: []byte("abd"),
			target:  byte('c'),
			ans:     byte('d'),
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v %v", testCase.letters, testCase.target)
		t.Run(testName, func(t *testing.T) {
			ans := nextGreatestLetter(testCase.letters, testCase.target)
			if testCase.ans != ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
