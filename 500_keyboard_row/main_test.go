package problem500

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindWords(t *testing.T) {
	testCases := []struct {
		words []string
		ans   []string
	}{
		{
			words: []string{"Hello", "Alaska", "Dad", "Peace"},
			ans:   []string{"Alaska", "Dad"},
		},
		{
			words: []string{"omk"},
			ans:   []string{},
		},
		{
			words: []string{"adsdf", "sfd"},
			ans:   []string{"adsdf", "sfd"},
		},
		{
			words: []string{"qws", "xdr"},
			ans:   []string{},
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.words)
		t.Run(testName, func(t *testing.T) {
			ans := findWords(testCase.words)
			if !reflect.DeepEqual(testCase.ans, ans) {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
