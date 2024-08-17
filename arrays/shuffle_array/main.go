package shufflearray

func Shuffle(nums []int, n int) []int {
	vector := make([]int, len(nums))
	g := 0
	j := n
	for i := 0; i < len(vector); i++ {
		if i%2 == 0 {
			vector[i] = nums[g]
			g++
		} else {
			vector[i] = nums[j]
			j++
		}
	}
	return vector
}
