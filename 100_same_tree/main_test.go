package problem100

import (
	"fmt"
	"testing"
)

func TestIsSameTree(t *testing.T) {
	testCases := []struct {
		p   *TreeNode
		q   *TreeNode
		ans bool
	}{
		{
			p:   &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			q:   &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			ans: true,
		},
		{
			p:   &TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
			q:   &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
			ans: false,
		},
		{
			p:   &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			q:   &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}}},
			ans: false,
		},
		{
			p:   &TreeNode{Val: 1},
			q:   &TreeNode{Val: 1},
			ans: true,
		},
		{
			p:   &TreeNode{Val: 2},
			q:   &TreeNode{Val: 1},
			ans: false,
		},
		{
			p:   nil,
			q:   &TreeNode{Val: 1},
			ans: false,
		},
		{
			p:   &TreeNode{Val: 1},
			q:   nil,
			ans: false,
		},
		{
			p:   nil,
			q:   nil,
			ans: true,
		},
	}

	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v %v", testCase.p, testCase.q)
		t.Run(testName, func(t *testing.T) {
			ans := isSameTree(testCase.p, testCase.q)
			if testCase.ans != ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
