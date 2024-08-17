package numberofemployees

func NumberOfEmployeesWhoMetTarget(hours []int, target int) int {
	counter := 0
	for _, val := range hours {
		if val >= target {
			counter++
		}
	}
	return counter
}
