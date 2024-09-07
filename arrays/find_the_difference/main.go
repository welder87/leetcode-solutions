package findthedifference

func FindTheDifference(s string, t string) byte {
	vector1 := quickSort([]byte(s))
	vector2 := quickSort([]byte(t))
	latestIndexOfVector1 := len(vector1) - 1
	var foundValue byte
	for idx := 0; idx < len(vector2); idx++ {
		second_val := vector2[idx]
		if idx > latestIndexOfVector1 || vector1[idx] != second_val {
			foundValue = second_val
			break
		}
	}
	return foundValue
}

func quickSort(vector []byte) []byte {
	lengthOfVector := len(vector)
	if lengthOfVector <= 1 {
		return vector
	}
	pivot := vector[lengthOfVector/2]
	var left []byte
	var middle []byte
	var right []byte

	for _, val := range vector {
		if val < pivot {
			left = append(left, val)
		} else if val > pivot {
			right = append(right, val)
		} else {
			middle = append(middle, val)
		}
	}

	new := make([]byte, 0, lengthOfVector)
	new = append(new, quickSort(left)...)
	new = append(new, middle...)
	new = append(new, quickSort(right)...)
	return new
}
