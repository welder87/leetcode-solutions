package kidswithcandies

func KidsWithCandies(candies []int, extraCandies int) []bool {
	max := 0
	for _, val := range candies {
		if val > max {
			max = val
		}
	}
	results := make([]bool, len(candies))
	for idx, val := range candies {
		results[idx] = val + extraCandies >= max
	}
	return results
}
