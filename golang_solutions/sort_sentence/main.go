package sortsentence

import (
	"strconv"
	"strings"
)

func SortSentence(s string) string {
	sentence := strings.Split(s, " ")
	source := make([]string, len(sentence))
	for _, val := range sentence {
		lastIndex := len(val) - 1
		lastSymbol := string(val[lastIndex])
		order, _ := strconv.Atoi(lastSymbol)
		source[order-1] = val[:lastIndex]
	}
	return strings.Join(source, " ")
}

func SortSentence1(s string) string {
	sentence := strings.Split(s, " ")
	sentence = quickSort(sentence)
	source := make([]string, 0, len(sentence))
	for _, val := range sentence {
		source = append(source, val[:len(val)-1])
	}
	return strings.Join(source, " ")
}

func quickSort(vector []string) []string {
	lengthOfVector := len(vector)
	if len(vector) <= 1 {
		return vector
	}
	middleString := vector[lengthOfVector/2]
	middleSymbol := middleString[len(middleString)-1]
	left := make([]string, 0, lengthOfVector)
	middle := make([]string, 0, lengthOfVector)
	right := make([]string, 0, lengthOfVector)
	for _, val := range vector {
		lastSymbol := val[len(val)-1]
		if lastSymbol < middleSymbol {
			left = append(left, val)
		} else if lastSymbol > middleSymbol {
			right = append(right, val)
		} else {
			middle = append(middle, val)
		}
	}
	new := make([]string, 0, lengthOfVector)
	new = append(new, quickSort(left)...)
	new = append(new, middle...)
	new = append(new, quickSort(right)...)
	return new
}
