package findthedifference

func FindTheDifference(s string, t string) byte {
	sMap := make(map[byte]int)
	tMap := make(map[byte]int)
	var val byte
	for i := 0; i < len(s); i++ {
		val = s[i]
		if count, ok := sMap[val]; ok {
			sMap[val] = count + 1
		} else {
			sMap[val] = 1
		}
	}
	for i := 0; i < len(t); i++ {
		val = t[i]
		if count, ok := tMap[val]; ok {
			tMap[val] = count + 1
		} else {
			tMap[val] = 1
		}
	}
	var foundValue byte
	for key, tCount := range tMap {
		if sCount, ok := sMap[key]; !ok || tCount != sCount {
			foundValue = key
		}
	}
	return foundValue
}

func FindTheDifference1(s string, t string) byte {
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
