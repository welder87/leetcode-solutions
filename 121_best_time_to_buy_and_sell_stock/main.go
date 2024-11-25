package problem121

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	mi := prices[0]
	mx := 0
	for i := 1; i < len(prices); i++ {
		if profit := prices[i] - mi; profit > mx {
			mx = profit
		} else if mi > prices[i] {
			mi = prices[i]
		}
	}
	return mx
}
