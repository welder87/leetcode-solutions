package backspacecompare

import "errors"

func BackspaceCompare(s string, t string) bool {
	si1 := StringIterator{
		index:  len(s) - 1,
		str:    s,
		offset: 0,
	}
	si2 := StringIterator{
		index:  len(t) - 1,
		str:    t,
		offset: 0,
	}
	for {
		val1, err1 := si1.GetNext()
		val2, err2 := si2.GetNext()
		if err1 != nil && err2 != nil {
			return true
		}
		if err1 == nil && err2 != nil || err2 == nil && err1 != nil {
			return false
		}
		if val1 != val2 {
			return false
		}
	}
}

type StringIterator struct {
	index  int
	str    string
	offset int
}

func (si *StringIterator) HasNext() bool {
	return si.index >= 0
}

func (si *StringIterator) GetNext() (symbol byte, err error) {
	sharp := byte('#')
	for {
		if si.HasNext() {
			symbol := si.str[si.index]
			si.index--
			if symbol == sharp {
				si.offset++
			} else if si.offset > 0 {
				si.offset--
			} else if si.offset == 0 {
				return symbol, nil
			}
		} else {
			return byte(0), errors.New("end of string")
		}
	}
}
