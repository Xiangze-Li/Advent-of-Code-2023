//nolint:gomnd, gochecknoglobals, funlen, cyclop // dont bother me
package puzzle07

import (
	"advent2023/util"
	"sort"
)

var cardStrength1 = map[byte]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}
var cardStrength2 = map[byte]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	'J': 1,
}

func convertCardStrength1(s string) [5]int {
	util.Assert(len(s) == 5, "")
	var res [5]int
	for i := 0; i < 5; i++ {
		res[i] = cardStrength1[s[i]]
	}
	return res
}
func convertCardStrength2(s string) [5]int {
	util.Assert(len(s) == 5, "")
	var res [5]int
	for i := 0; i < 5; i++ {
		res[i] = cardStrength2[s[i]]
	}
	return res
}

type handType int

const (
	HighCard handType = iota
	OnePair
	TwoPairs
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type hand struct {
	Cards [5]int
	Bid   uint64
	Type  handType
}

func (h *hand) parseType1() {
	m := map[int]int{}
	for _, c := range h.Cards {
		m[c]++
	}
	var five, four, three, pair1, pair2 bool
	for _, v := range m {
		switch v {
		case 5:
			five = true
		case 4:
			four = true
		case 3:
			three = true
		case 2:
			if pair1 {
				pair2 = true
			} else {
				pair1 = true
			}
		}
	}

	switch {
	case five:
		h.Type = FiveOfAKind
	case four:
		h.Type = FourOfAKind
	case three && pair1:
		h.Type = FullHouse
	case three:
		h.Type = ThreeOfAKind
	case pair1 && pair2:
		h.Type = TwoPairs
	case pair1:
		h.Type = OnePair
	default:
		h.Type = HighCard
	}
}

func (h *hand) parseType2() {
	m := map[int]int{}
	for _, c := range h.Cards {
		m[c]++
	}
	numJ := m[cardStrength2['J']]
	delete(m, cardStrength2['J'])
	var five, four, three, pair1, pair2 bool
	for _, v := range m {
		switch v {
		case 5:
			five = true
		case 4:
			four = true
		case 3:
			three = true
		case 2:
			if pair1 {
				pair2 = true
			} else {
				pair1 = true
			}
		}
	}

	switch numJ {
	case 5, 4:
		h.Type = FiveOfAKind
	case 3:
		if pair1 {
			h.Type = FiveOfAKind
		} else {
			h.Type = FourOfAKind
		}
	case 2:
		switch {
		case three:
			h.Type = FiveOfAKind
		case pair1:
			h.Type = FourOfAKind
		default:
			h.Type = ThreeOfAKind
		}
	case 1:
		switch {
		case four:
			h.Type = FiveOfAKind
		case three:
			h.Type = FourOfAKind
		case pair1 && pair2:
			h.Type = FullHouse
		case pair1:
			h.Type = ThreeOfAKind
		default:
			h.Type = OnePair
		}
	case 0:
		switch {
		case five:
			h.Type = FiveOfAKind
		case four:
			h.Type = FourOfAKind
		case three && pair1:
			h.Type = FullHouse
		case three:
			h.Type = ThreeOfAKind
		case pair1 && pair2:
			h.Type = TwoPairs
		case pair1:
			h.Type = OnePair
		default:
			h.Type = HighCard
		}
	}
}

type handSlice []hand

func (hs handSlice) Len() int      { return len(hs) }
func (hs handSlice) Swap(i, j int) { hs[i], hs[j] = hs[j], hs[i] }
func (hs handSlice) Less(i, j int) bool {
	if hs[i].Type != hs[j].Type {
		return hs[i].Type < hs[j].Type
	}
	for k := 0; k < 5; k++ {
		if hs[i].Cards[k] != hs[j].Cards[k] {
			return hs[i].Cards[k] < hs[j].Cards[k]
		}
	}
	return false
}

var _ sort.Interface = handSlice{}
