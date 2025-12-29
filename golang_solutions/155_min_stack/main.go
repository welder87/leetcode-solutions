package problem155

// Amortized time complexity: O(1). Amortized space complexity: O(n + n).
// Two-stack solution.
type MinStack struct {
	stack    []int
	minimums []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.minimums) == 0 {
		this.minimums = append(this.minimums, val)
	} else {
		this.minimums = append(
			this.minimums,
			min(val, this.minimums[len(this.minimums)-1]),
		)
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	this.stack = this.stack[:len(this.stack)-1]
	this.minimums = this.minimums[:len(this.minimums)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.minimums)-1]
}

func (this *MinStack) GetMin() int {
	return this.minimums[len(this.minimums)-1]
}
