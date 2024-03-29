// Code generated by "stringer -type=Type"; DO NOT EDIT.

package day7

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[HighCard-1]
	_ = x[OnePair-2]
	_ = x[TwoPair-3]
	_ = x[ThreeOfAKind-4]
	_ = x[FullHouse-5]
	_ = x[FourOfAKind-6]
	_ = x[FiveOfAKind-7]
}

const _Type_name = "HighCardOnePairTwoPairThreeOfAKindFullHouseFourOfAKindFiveOfAKind"

var _Type_index = [...]uint8{0, 8, 15, 22, 34, 43, 54, 65}

func (i Type) String() string {
	i -= 1
	if i < 0 || i >= Type(len(_Type_index)-1) {
		return "Type(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
