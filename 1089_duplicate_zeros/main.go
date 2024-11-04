package problem1086

func DuplicateZeros(arr []int) {
	if len(arr) == 1 || len(arr) == 0 {
		return
	}
	i := 0
	counter := 0
	for counter < len(arr) {
		if arr[i] == 0 {
			counter += 2
		} else {
			counter++
		}
		i++
	}
	if i == len(arr)-1 {
		return
	}
	j := len(arr) - 1
	for i < j && i >= 0 {
		if arr[i] == 0 {
			for k := 0; k < 2; k++ {
				arr[j] = 0
				j--
			}
			i--
		} else {
			arr[j] = arr[i]
			i--
			j--
		}
	}
}
