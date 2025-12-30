package problem1047

// Time complexity O(n + n), space complexity O(n + m))
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

// Time complexity O(n + n), space complexity O(n + m))
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
