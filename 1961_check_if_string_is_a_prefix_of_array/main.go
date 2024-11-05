package problem1961

import "errors"

func IsPrefixString(s string, words []string) bool {
	if len(s) < len(words[0]) {
		return false
	}
	wi := WordIter{
		index1: 0,
		index2: 0,
		words:  &words,
		word:   words[0],
	}
	i := 0
	for i < len(s) {
		if val, err := wi.GetNext(); err != nil || s[i] != val {
			return false
		}
		i++
	}
	return true
}

type WordIter struct {
	index1 int
	index2 int
	words  *[]string
	word   string
}

func (wi *WordIter) GetNext() (symbol byte, err error) {
	for {
		if wi.index1 < len(*wi.words) {
			if wi.index2 < len(wi.word) {
				symbol := wi.word[wi.index2]
				wi.index2++
				return symbol, nil
			} else {
				wi.index1++
				if wi.index1 < len(*wi.words) {
					wi.word = (*wi.words)[wi.index1]
					wi.index2 = 0
				} else {
					return byte(0), errors.New("end of words")
				}
			}
		} else {
			return byte(0), errors.New("end of words")
		}
	}
}
