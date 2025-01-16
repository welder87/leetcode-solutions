package problem500

import "strings"

func findWords(words []string) []string {
	rows := []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
	rowsSet := make([]map[byte]struct{}, 0, len(rows))
	for _, s := range rows {
		set := make(map[byte]struct{}, len(s))
		for i := 0; i < len(s); i++ {
			set[s[i]] = struct{}{}
		}
		rowsSet = append(rowsSet, set)
	}
	rows = rows[:0]

	ans := make([]string, 0, len(rows))
	for _, word := range words {
		set, isFound := &rowsSet[0], false
		wordInLower := strings.ToLower(word)
		for i := 0; i < len(rowsSet); i++ {
			if _, ok := rowsSet[i][wordInLower[0]]; ok {
				set = &rowsSet[i]
				isFound = true
				break
			}
		}
		if !isFound {
			continue
		}
		inRow := true
		for i := 1; i < len(wordInLower); i++ {
			if _, ok := (*set)[wordInLower[i]]; !ok {
				inRow = false
				break
			}
		}
		if inRow {
			ans = append(ans, word)
		}
	}
	return ans
}
