package problem1833

func maxIceCream(costs []int, coins int) int {
	maxCost, minCost := -1, costs[0]
	for _, cost := range costs {
		if maxCost < cost {
			maxCost = cost
		}
		if minCost > cost {
			minCost = cost
		}
	}
	if coins < minCost {
		return 0
	}
	counter := make([]int, maxCost+1)
	for _, cost := range costs {
		counter[cost]++
	}
	sortedCosts := make([]int, 0, len(costs))
	for cost, cnt := range counter {
		if cnt == 0 {
			continue
		}
		for i := 0; i < cnt; i++ {
			sortedCosts = append(sortedCosts, cost)
		}
	}
	sum := 0
	i := 0
	for i < len(sortedCosts) {
		sum += sortedCosts[i]
		if sum > coins {
			break
		}
		i++
	}
	return i
}
