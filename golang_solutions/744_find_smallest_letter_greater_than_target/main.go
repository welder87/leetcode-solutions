package problem744

// O(n)
func nextGreatestLetter(letters []byte, target byte) byte {
	for i := 0; i < len(letters); i++ {
		if target < letters[i] {
			return letters[i]
		}
	}
	return letters[0]
}

func nextGreatestLetterV1(letters []byte, target byte) byte {
	if target < letters[0] || target >= letters[len(letters)-1] {
		return letters[0]
	}
	left := 0
	right := len(letters)
	for left < right {
		mid := left + (right-left)/2
		if letters[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return letters[left%len(letters)]
}
