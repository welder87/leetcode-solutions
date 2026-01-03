package problem1047

import "bytes"

// Time complexity O(n + n + m), space complexity O(n + m + k))
func removeDuplicates(s string) string {
	stack := make([]byte, 0, len(s))
	for i := range s {
		if len(stack) == 0 || stack[len(stack)-1] != s[i] {
			stack = append(stack, s[i])
		} else {
			val := stack[len(stack)-1]
			for len(stack) > 0 && val == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return string(stack)
}

// Time complexity O(n + n + m), space complexity O(n + m + k))
func removeDuplicatesV1(s string) string {
	stack := make([]byte, 0, len(s))
	for i := range s {
		if len(stack) == 0 || stack[len(stack)-1] != s[i] {
			stack = append(stack, s[i])
		} else {
			val := stack[len(stack)-1]
			j := len(stack) - 1
			for j >= 0 && val == stack[j] {
				j--
			}
			stack = stack[:j+1]
		}
	}
	return string(stack)
}

// Time complexity O(n + m + k + t), space complexity O(n + m + k))
func removeDuplicatesV2(s string) string {
	var top *Node = nil
	for i := range s {
		if top == nil || top.val != s[i] {
			node := Node{val: s[i], next: top}
			top = &node
		} else {
			val := top.val
			for top != nil && val == top.val {
				top = top.next
			}
		}
	}
	ans := make([]byte, len(s))
	i := len(ans) - 1
	for top != nil {
		ans[i] = top.val
		top = top.next
		i--
	}
	return string(bytes.TrimLeft(ans, "\x00"))
}

type Node struct {
	next *Node
	val  byte
}
