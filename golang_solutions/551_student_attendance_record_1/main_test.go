package problem551

import (
	"testing"
)

func TestCheckRecord(t *testing.T) {
	testCheckRecord(t, checkRecord)
}

func TestCheckRecordV1(t *testing.T) {
	testCheckRecord(t, checkRecordV1)
}

func testCheckRecord(t *testing.T, fn func(string) bool) {
	testCases := []struct {
		s    string
		ans  bool
		name string
	}{
		{"PPALLP", true, "Preset Case 1"},
		{"PPALLL", false, "Preset Case 2"},
		{"AA", false, "Preset Case 3"},
		{"P", true, "Single character P"},
		{"A", true, "Single character A"},
		{"L", true, "Single character L"},
		{"PAPLP", true, "Mixed with 1 Absence"},
		{"PAPALP", false, "Mixed with 2 Absences"},
		{"APALA", false, "Multiple Absences separated"},
		{"LL", true, "Exactly 2 consecutive L's"},
		{"LLL", false, "Exactly 3 consecutive L's"},
		{"LPLPL", true, "3 L's with break"},
		{"LLPAP", true, "Late at beginning"},
		{"PAPLL", true, "Late at end"},
		{"LLLL", false, "4 consecutive L's"},
		{"LLLLL", false, "More than 3 consecutive L's"},
		{"ALLP", true, "One A and 2 consecutive L's"},
		{"ALLLP", false, "One A and 3 consecutive L's"},
		{"ALLAP", false, "Two A's with late"},
		{"PPPPP", true, "All Present"},
		{"PPPPPPPPPP", true, "Long present string"},
		{"PLLPALLP", true, "Valid complex case 1"},
		{"LPAPLPLL", true, "Valid complex case 2"},
		{"PLAPLLP", true, "Valid with A in middle"},
		{"PAPAP", false, "Invalid: 2 A's"},
		{"PLLLP", false, "Invalid: 3 consecutive L's in middle"},
		{"ALLAPLL", false, "Invalid: both conditions"},
		{"ALLL", false, "Invalid: A and LLL"},
		{"APP", true, "Exactly 2 P's at ends"},
		{"PAAP", false, "Exactly 2 A's in middle"},
		{"LLLP", false, "3 L's at start"},
		{"PLLL", false, "3 L's at end"},
		{"LLALL", true, "3 L's separated by A"},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ans := fn(testCase.s)
			if ans != testCase.ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}
