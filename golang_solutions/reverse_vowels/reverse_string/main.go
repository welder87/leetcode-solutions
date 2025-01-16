package reversevowels

func ReverseVowels(s string) string {
	vowels := map[byte]struct{}{
		[]byte("a")[0]: {},
		[]byte("e")[0]: {},
		[]byte("i")[0]: {},
		[]byte("o")[0]: {},
		[]byte("u")[0]: {},

		[]byte("A")[0]: {},
		[]byte("E")[0]: {},
		[]byte("I")[0]: {},
		[]byte("O")[0]: {},
		[]byte("U")[0]: {},
	}
	vectorOfBytes := []byte(s)
	i := 0
	j := len(s) - 1
	for i < j {
		_, ok1 := vowels[vectorOfBytes[i]]
		_, ok2 := vowels[vectorOfBytes[j]]
		if ok1 && !ok2 {
			j--
		} else if !ok1 && ok2 {
			i++
		} else if ok1 && ok2 {
			vectorOfBytes[i], vectorOfBytes[j] = vectorOfBytes[j], vectorOfBytes[i]
			i++
			j--
		} else {
			i++
			j--
		}
	}
	return string(vectorOfBytes)
}
