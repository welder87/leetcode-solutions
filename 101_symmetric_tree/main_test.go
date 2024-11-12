package problem100

import (
	"fmt"
	"testing"
)

func TestIsSymmetric(t *testing.T) {
	testCases := []struct {
		root *TreeNode
		ans  bool
	}{
		{
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
				Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}},
			},
			ans: true,
		},
		{
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
				Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
			},
			ans: false,
		},
	}

	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v", testCase.root)
		t.Run(testName, func(t *testing.T) {
			ans := isSymmetric(testCase.root)
			if testCase.ans != ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
