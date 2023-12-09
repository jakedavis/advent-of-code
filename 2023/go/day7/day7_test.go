package day7

import "testing"

func TestHandType(t *testing.T) {
	tt := []struct {
		Hand     Hand
		Expected Type
	}{
		{
			Hand:     []Card{Three, Two, Ten, Nine, King},
			Expected: HighCard,
		},
		{
			Hand:     []Card{Three, Two, Ten, Three, King},
			Expected: OnePair,
		},
		{
			Hand:     []Card{Three, Two, Two, Three, King},
			Expected: TwoPair,
		},
		{
			Hand:     []Card{Ten, Five, Five, Jack, Five},
			Expected: ThreeOfAKind,
		},
		{
			Hand:     []Card{Ten, Five, Five, Ten, Five},
			Expected: FullHouse,
		},
		{
			Hand:     []Card{Five, Five, Five, Five, Ten},
			Expected: FourOfAKind,
		},
		{
			Hand:     []Card{Five, Five, Five, Five, Five},
			Expected: FiveOfAKind,
		},
	}

	for _, test := range tt {
		if handType := test.Hand.Type(); handType != Type(test.Expected) {
			t.Errorf("type %v did not match expected %v", handType, test.Expected)
		}
	}
}

func TestHandJokerType(t *testing.T) {
	useJokers = true

	tt := []struct {
		Hand     Hand
		Expected Type
	}{
		{
			Hand:     []Card{Three, Two, Ten, Joker, King},
			Expected: OnePair,
		},
		{
			Hand:     []Card{Three, Two, Joker, Three, King},
			Expected: ThreeOfAKind,
		},
		{
			Hand:     []Card{Three, Two, Two, Three, King},
			Expected: TwoPair,
		},
		{
			Hand:     []Card{Ten, Five, Five, Joker, Five},
			Expected: FourOfAKind,
		},
		{
			Hand:     []Card{Ten, Five, Five, Ten, Five},
			Expected: FullHouse,
		},
		{
			Hand:     []Card{Five, Five, Five, Joker, Ten},
			Expected: FourOfAKind,
		},
		{
			Hand:     []Card{Five, Five, Joker, Five, Five},
			Expected: FiveOfAKind,
		},
	}

	for i, test := range tt {
		if handType := test.Hand.Type(); handType != test.Expected {
			t.Errorf("[%d] type %v did not match expected %v", i, handType, test.Expected)
		}
	}
}
