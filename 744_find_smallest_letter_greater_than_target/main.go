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
