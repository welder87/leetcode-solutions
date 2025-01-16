package problem744

import (
	"fmt"
	"testing"
)

func TestNextGreatestLetter(t *testing.T) {
	testNextGreatestLetter(t, nextGreatestLetter)
}

func TestNextGreatestLetterV1(t *testing.T) {
	testNextGreatestLetter(t, nextGreatestLetterV1)
}

func testNextGreatestLetter(t *testing.T, function fn) {
	testCases := []struct {
		letters []byte
		target  byte
		ans     byte
	}{
		{
			letters: []byte("cfj"),
			target:  byte('a'),
			ans:     byte('c'),
		},
		{
			letters: []byte("cfj"),
			target:  byte('c'),
			ans:     byte('f'),
		},
		{
			letters: []byte("xxyy"),
			target:  byte('z'),
			ans:     byte('x'),
		},
		{
			letters: []byte("xy"),
			target:  byte('z'),
			ans:     byte('x'),
		},
		{
			letters: []byte("xy"),
			target:  byte('y'),
			ans:     byte('x'),
		},
		{
			letters: []byte("xy"),
			target:  byte('x'),
			ans:     byte('y'),
		},
		{
			letters: []byte("abd"),
			target:  byte('c'),
			ans:     byte('d'),
		},
		{
			letters: []byte("abd"),
			target:  byte('j'),
			ans:     byte('a'),
		},
		{
			letters: []byte("abcdefjhkm"),
			target:  byte('b'),
			ans:     byte('c'),
		},
		{
			letters: []byte("abcdefjhkmn"),
			target:  byte('b'),
			ans:     byte('c'),
		},
		{
			letters: []byte("acdefjhkmn"),
			target:  byte('b'),
			ans:     byte('c'),
		},
		{
			letters: []byte("ajhkmnoprs"),
			target:  byte('b'),
			ans:     byte('j'),
		},
	}
	for _, testCase := range testCases {
		testName := fmt.Sprintf("%v %v", string(testCase.letters), string(testCase.target))
		t.Run(testName, func(t *testing.T) {
			ans := function(testCase.letters, testCase.target)
			if testCase.ans != ans {
				t.Errorf("got %v, want %v", ans, testCase.ans)
			}
		})
	}
}

type fn func(letters []byte, target byte) byte
